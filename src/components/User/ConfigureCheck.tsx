import { useMemo, useState } from "react";

import { Box, Button, Chip, Typography } from "@mui/material";

import { enqueueSnackbar } from "notistack";
import { ConfigField, Dropdown } from "..";
import { ConfigsQuery, useEditConfigMutation } from "../../graph";

type props = {
  config: ConfigsQuery["configs"][0];
  handleRefetch: () => void;
  visible: boolean;
};

export default function ConfigureCheck({
  config,
  visible,
  handleRefetch,
}: props) {
  const [expanded, setExpanded] = useState(false);

  const [checkConfig, setCheckConfig] = useState<{
    [key: string]: number | boolean | string;
  }>(JSON.parse(config.config));
  const configChanged = useMemo(
    () => JSON.stringify(checkConfig) != config.config,
    [checkConfig, config.config]
  );

  const [useEditConfig] = useEditConfigMutation({
    onCompleted: () => {
      enqueueSnackbar("Config saved", { variant: "success" });
      handleRefetch();
    },
    onError: (error) => {
      enqueueSnackbar("Error saving config", { variant: "error" });
      console.error(error);
    },
  });

  const handleConfigChange = (
    key: string,
    value: string | number | boolean
  ) => {
    setCheckConfig({ ...checkConfig, [key]: value });
  };

  const handleSave = () => {
    useEditConfig({
      variables: {
        id: config.id,
        config: JSON.stringify(checkConfig),
      },
    });
  };

  return (
    <Dropdown
      title={
        <>
          <Typography variant='h6' component='div' marginRight='24px'>
            {config.check.name}
          </Typography>
          <Typography
            variant='subtitle1'
            color='textSecondary'
            component='div'
            marginRight='24px'
          >
            {config.check.source.name}
          </Typography>
          <Chip size='small' label={`weight:${config.check.weight}`} />
        </>
      }
      toggleButton={
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
      }
      expanded={expanded}
      setExpanded={setExpanded}
      toggleButtonVisible={configChanged}
      visible={visible}
    >
      <Box
        sx={{
          display: "flex",
          gap: "16px",
          flexWrap: "wrap",
          justifyContent: "center",
        }}
      >
        {Object.keys(checkConfig).length ? (
          <>
            {Object.entries(checkConfig).map(([key]) => (
              <ConfigField
                key={key}
                index={key}
                handleInputChange={handleConfigChange}
                value={
                  JSON.parse(config.check.source.schema)[key] as
                    | "string"
                    | "int"
                    | "bool"
                }
                config={checkConfig}
              />
            ))}
          </>
        ) : (
          <Typography variant='h5'>No configuration required</Typography>
        )}
      </Box>
    </Dropdown>
  );
}
