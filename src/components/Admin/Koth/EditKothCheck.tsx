import { useMemo, useState } from "react";

import { Box, Button, Chip, TextField, Typography } from "@mui/material";

import { enqueueSnackbar } from "notistack";
import { DeleteKothCheckModal, Dropdown } from "../..";
import {
  KothChecksQuery,
  useDeleteKothCheckMutation,
  useUpdateKothCheckMutation,
} from "../../../graph";

type props = {
  check: KothChecksQuery["kothChecks"][0];
  handleRefetch: () => void;
  visible: boolean;
};

export default function EditKothCheck({
  check,
  visible,
  handleRefetch,
}: props) {
  const [expanded, setExpanded] = useState(false);

  const [name, setName] = useState(check.name);
  const nameChanged = useMemo(() => name != check.name, [name, check.name]);

  const [file, setFile] = useState(check.file);
  const fileChanged = useMemo(() => file != check.file, [file, check.file]);

  const [host, setHost] = useState(check.host);
  const hostChanged = useMemo(() => host != check.host, [host, check.host]);

  const [weight, setWeight] = useState<number>(check.weight);
  const weightChanged = useMemo(
    () => weight != check.weight,
    [weight, check.weight]
  );

  const [topic, setTopic] = useState(check.topic);
  const topicChanged = useMemo(
    () => topic != check.topic,
    [topic, check.topic]
  );

  const [open, setOpen] = useState(false);

  const [updateKothCheckMutation] = useUpdateKothCheckMutation({
    onCompleted: () => {
      enqueueSnackbar("KoTH Check updated successfully", {
        variant: "success",
      });
      handleRefetch();
    },
    onError: (error) => {
      enqueueSnackbar(error.message, { variant: "error" });
      console.error(error);
    },
  });

  const [DeleteKothCheckMutation] = useDeleteKothCheckMutation({
    onCompleted: () => {
      enqueueSnackbar("Check deleted successfully", { variant: "success" });
      handleRefetch();
    },
    onError: (error) => {
      enqueueSnackbar(error.message, { variant: "error" });
      console.error(error);
    },
  });

  const handleSave = () => {
    updateKothCheckMutation({
      variables: {
        id: check.id,
        name: nameChanged ? name : undefined,
        file: fileChanged ? file : undefined,
        host: hostChanged ? host : undefined,
        weight: weightChanged ? weight : undefined,
        topic: topicChanged ? topic : undefined,
      },
    });
  };

  const handleDelete = () => {
    DeleteKothCheckMutation({
      variables: {
        id: check.id,
      },
    });
  };

  return (
    <Dropdown
      modal={
        <DeleteKothCheckModal
          check={check.name}
          open={open}
          setOpen={setOpen}
          handleDelete={handleDelete}
        />
      }
      title={
        <>
          {expanded ? (
            <TextField
              label='Name'
              value={name}
              onClick={(e) => {
                e.stopPropagation();
              }}
              onChange={(e) => {
                setName(e.target.value);
              }}
              sx={{ marginRight: "24px" }}
              size='small'
            />
          ) : (
            <Typography variant='h6' component='div' marginRight='24px'>
              {check.name}
            </Typography>
          )}
          {expanded ? (
            <>
              <TextField
                label='Weight'
                type='number'
                value={weight}
                onClick={(e) => {
                  e.stopPropagation();
                }}
                onChange={(e) => {
                  setWeight(parseInt(e.target.value));
                }}
                sx={{ marginRight: "24px", width: "100px" }}
                size='small'
              />
              <TextField
                label='KoTH Topic Key'
                type='text'
                value={topic}
                onClick={(e) => {
                  e.stopPropagation();
                }}
                onChange={(e) => {
                  setTopic(e.target.value);
                }}
                sx={{ marginRight: "24px", width: "100px" }}
                size='small'
              />
            </>
          ) : (
            <>
              <Chip size='small' label={`weight:${weight}`} />
              <Chip
                size='small'
                label={`topic:${topic}`}
                onClick={(e) => {
                  navigator.clipboard.writeText(topic).then(
                    () => {
                      enqueueSnackbar("Topic copied to clipboard", {
                        variant: "success",
                      });
                    },
                    (err) => {
                      enqueueSnackbar("Failed to copy topic to clipboard", {
                        variant: "error",
                      });
                      console.error(err);
                    }
                  );
                  e.stopPropagation();
                }}
              />
            </>
          )}
        </>
      }
      expandableButtons={[
        <Button
          variant='contained'
          onClick={() => {
            setOpen(true);
          }}
          color='error'
        >
          Delete
        </Button>,
      ]}
      toggleButton={
        <Box sx={{ display: "contents" }}>
          <Button
            variant='contained'
            color='success'
            onClick={(e) => {
              if (!expanded) {
                e.stopPropagation();
              }

              handleSave();
            }}
          >
            Save
          </Button>
        </Box>
      }
      expanded={expanded}
      setExpanded={setExpanded}
      visible={visible}
      toggleButtonVisible={
        nameChanged ||
        fileChanged ||
        weightChanged ||
        hostChanged ||
        topicChanged
      }
    >
      <Box
        sx={{
          display: "flex",
          gap: "16px",
          flexWrap: "wrap",
          justifyContent: "center",
        }}
      >
        <TextField
          label='File'
          value={file}
          onClick={(e) => {
            e.stopPropagation();
          }}
          onChange={(e) => {
            setFile(e.target.value);
          }}
        />
        <TextField
          label='Hostname'
          value={host}
          onClick={(e) => {
            e.stopPropagation();
          }}
          onChange={(e) => {
            setHost(e.target.value);
          }}
        />
      </Box>
    </Dropdown>
  );
}
