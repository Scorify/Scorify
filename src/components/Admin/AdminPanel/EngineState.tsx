import { Box, Button, ButtonGroup, Container, Typography } from "@mui/material";
import { enqueueSnackbar } from "notistack";

import {
  EngineState,
  useStartEngineMutation,
  useStopEngineMutation,
} from "../../../graph";

type props = {
  engineState: EngineState | undefined;
};

export default function EngineStateComponent({ engineState }: props) {
  const [StartEngine] = useStartEngineMutation({
    onError: (error) => {
      enqueueSnackbar(error.message, { variant: "error" });
      console.error(error);
    },
  });

  const [StopEngine] = useStopEngineMutation({
    onError: (error) => {
      enqueueSnackbar(error.message, { variant: "error" });
      console.error(error);
    },
  });

  const color = {
    [EngineState.Paused]: "error",
    [EngineState.Running]: "success",
    [EngineState.Waiting]: "warning",
    [EngineState.Stopping]: "secondary",
    default: "info",
  };

  return (
    <Container maxWidth='xs'>
      <Typography variant='h4' align='center'>
        Engine State
      </Typography>
      <Box
        sx={{ m: 2, mt: 4 }}
        display='flex'
        alignItems='center'
        flexDirection='column'
      >
        <Button
          variant='contained'
          color={
            color[engineState || "default"] as
              | "error"
              | "success"
              | "warning"
              | "info"
          }
        >
          <Typography variant='h5'>{engineState || "disconnected"}</Typography>
        </Button>

        <Box sx={{ m: 2 }} />
        <ButtonGroup variant='contained' fullWidth>
          <Button
            onClick={() => {
              StartEngine();
            }}
            disabled={
              engineState === EngineState.Running ||
              engineState === EngineState.Waiting
            }
          >
            <Typography variant='h6'>Start</Typography>
          </Button>
          <Button
            onClick={() => {
              StopEngine();
            }}
            disabled={engineState === EngineState.Paused}
          >
            <Typography variant='h6'>Stop</Typography>
          </Button>
        </ButtonGroup>
      </Box>
    </Container>
  );
}
