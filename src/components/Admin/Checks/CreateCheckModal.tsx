import { useEffect, useMemo, useState } from "react";

import {
  Box,
  Button,
  Divider,
  FormControl,
  InputLabel,
  MenuItem,
  Modal,
  Select,
  SelectChangeEvent,
  TextField,
  Typography,
} from "@mui/material";
import { enqueueSnackbar } from "notistack";

import { ConfigField, Multiselect } from "../..";
import {
  ChecksQuery,
  SchemaFieldType,
  useCreateCheckMutation,
  useValidateCheckMutation,
} from "../../../graph";

type props = {
  data?: ChecksQuery;
  open: boolean;
  setOpen: (isOpen: boolean) => void;
  handleRefetch: () => void;
};

export default function CreateCheckModal({
  data,
  open,
  setOpen,
  handleRefetch,
}: props) {
  const [source, setSource] = useState<string>("");
  const [name, setName] = useState<string>("");
  const [weight, setWeight] = useState<number>(1);
  const [config, setConfig] = useState<{
    [key: string]: string | number | boolean;
  }>({});
  const [editableFields, setEditableFields] = useState<string[]>([]);

  const [createCheckMutation] = useCreateCheckMutation({
    onCompleted: () => {
      enqueueSnackbar("Check created successfully", { variant: "success" });
      setOpen(false);
      handleRefetch();

      setSource("");
      setName("");
      setWeight(1);
      setConfig({});
      setEditableFields([]);
    },
    onError: (error) => {
      enqueueSnackbar(error.message, { variant: "error" });
    },
  });

  const [validated, setValidated] = useState(false);
  const [validationError, setValidationError] = useState<string | undefined>();

  const [validateCheckMutation] = useValidateCheckMutation({
    onCompleted: () => {
      setValidated(true);
      setValidationError(undefined);
      enqueueSnackbar("Check successfully validated", { variant: "success" });
    },
    onError: (error) => {
      enqueueSnackbar("Check failed to be validate", { variant: "error" });
      setValidationError(error.message);
    },
  });

  const sourceSchema = useMemo<ChecksQuery["sources"][0]["schema"] | undefined>(
    () => data?.sources.find((s) => s.name === source)?.schema,
    [source, data]
  );

  useEffect(() => {
    let newConfig = {} as {
      [key: string]: string | number | boolean;
    };

    if (sourceSchema) {
      for (const [_, field] of Object.entries(sourceSchema)) {
        if (field.type === SchemaFieldType.Bool) {
          newConfig[field.name] = field.default
            ? field.default.toLowerCase() === "true"
            : false;
        } else if (field.type === SchemaFieldType.Int) {
          newConfig[field.name] = field.default ? parseInt(field.default) : 0;
        } else if (field.type === SchemaFieldType.String) {
          newConfig[field.name] = field.default || "";
        }
      }
    }

    setValidated(false);
    setConfig(newConfig);
  }, [sourceSchema]);

  const handleInputChange = (key: string, value: string | number | boolean) => {
    setValidated(false);
    setConfig({
      ...config,
      [key]: value,
    });
  };

  return (
    <Modal
      open={open}
      onClose={() => {
        setOpen(false);
      }}
    >
      <Box
        sx={{
          position: "absolute",
          top: "25%",
          left: "50%",
          transform: "translate(-50%, -25%)",
          width: "auto",
          maxWidth: "90vw",
          bgcolor: "background.paper",
          border: `1px solid #000`,
          borderRadius: "8px",
          boxShadow: 24,
          p: 4,
        }}
      >
        <Box
          sx={{
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
          }}
        >
          <Typography component='h1' variant='h3'>
            Create New Check
          </Typography>
          <FormControl sx={{ marginTop: "24px" }}>
            <InputLabel id='source'>Source</InputLabel>
            <Select
              labelId='source'
              value={source}
              label='Source'
              onChange={(event: SelectChangeEvent) => {
                setSource(event.target.value as string);
              }}
            >
              <MenuItem value=''>None</MenuItem>
              {data?.sources.map((source) => (
                <MenuItem key={source.name} value={source.name}>
                  {source.name}
                </MenuItem>
              ))}
            </Select>

            <TextField
              label='Name'
              variant='outlined'
              sx={{
                marginTop: "24px",
              }}
              value={name}
              onChange={(event) => {
                setName(event.target.value as string);
              }}
            />

            <TextField
              label='Weight'
              variant='outlined'
              sx={{
                marginTop: "24px",
              }}
              type='number'
              value={weight}
              onChange={(event) => {
                setWeight(parseInt(event.target.value));
              }}
            />

            <Box sx={{ justifyContent: "center" }}>
              <Typography
                component='h1'
                variant='h4'
                marginTop='24px'
                align='center'
              >
                Check Configuration
              </Typography>
            </Box>
            <Box
              sx={{
                display: "flex",
                gap: "8px",
                flexWrap: "wrap",
                justifyContent: "center",
              }}
            >
              {source !== "" && data && sourceSchema ? (
                <Box
                  sx={{
                    display: "flex",
                    gap: "0px 16px",
                    flexWrap: "wrap",
                    justifyContent: "center",
                  }}
                >
                  {Object.entries(sourceSchema).map(([_, field]) => (
                    <ConfigField
                      key={field.name}
                      handleInputChange={handleInputChange}
                      fieldName={field.name}
                      fieldType={field.type}
                      defaultValue={field.default ?? undefined}
                      enumValues={field.enum ?? undefined}
                      checkConfig={config}
                    />
                  ))}
                </Box>
              ) : (
                <Typography component='h1' variant='body1' marginTop='12px'>
                  Select a source to see the configuration options
                </Typography>
              )}
            </Box>
            {validationError !== undefined && (
              <>
                <Divider sx={{ margin: "16px 20% 20px 20%" }} />
                <TextField
                  label='Validation Error'
                  multiline
                  error
                  value={validationError}
                  InputProps={{
                    readOnly: true,
                  }}
                />
              </>
            )}
            {source !== "" && data && sourceSchema && (
              <>
                <Divider sx={{ margin: "16px 20% 20px 20%" }} />
                <Multiselect
                  label='Set User Editable Fields'
                  placeholder='Select fields'
                  options={Object.keys(config)}
                  selected={editableFields}
                  setSelected={setEditableFields}
                />
              </>
            )}
          </FormControl>
          <Box sx={{ marginTop: "24px", display: "flex", gap: "24px" }}>
            <Button
              variant='contained'
              disabled={source === ""}
              onClick={() => {
                if (source === "") {
                  enqueueSnackbar("source must be set before validation", {
                    variant: "error",
                  });
                }

                validateCheckMutation({
                  variables: {
                    source: source,
                    config: JSON.stringify(config),
                  },
                });
              }}
            >
              Validate Check
            </Button>
            <Button
              variant='contained'
              color='success'
              disabled={source === "" || name === "" || !validated}
              onClick={() => {
                if (source === "") {
                  enqueueSnackbar("Source must be set", {
                    variant: "error",
                  });
                  return;
                }

                if (name === "") {
                  enqueueSnackbar("Name must be set", {
                    variant: "error",
                  });
                  return;
                }

                if (!validated) {
                  enqueueSnackbar("Configuration has not been validation", {
                    variant: "error",
                  });
                  return;
                }

                createCheckMutation({
                  variables: {
                    source: source,
                    name: name,
                    weight: weight,
                    config: JSON.stringify(config),
                    editable_fields: editableFields,
                    display: "",
                  },
                });
              }}
            >
              Create Check
            </Button>
          </Box>
        </Box>
      </Box>
    </Modal>
  );
}
