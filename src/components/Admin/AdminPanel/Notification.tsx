import { useState } from "react";

import {
  Box,
  Button,
  ButtonGroup,
  Container,
  MenuItem,
  Select,
  TextField,
  Typography,
} from "@mui/material";
import { enqueueSnackbar } from "notistack";

import {
  NotificationType,
  useSendGlobalNotificationMutation,
} from "../../../graph";

export default function Notification() {
  const [sendGlobalNotification] = useSendGlobalNotificationMutation({
    onCompleted: () => {
      enqueueSnackbar("Notification Sent", { variant: "success" });
    },
    onError: (error) => {
      enqueueSnackbar(error.message, { variant: "error" });
    },
  });

  const [message, setMessage] = useState("");
  const [type, setType] = useState<NotificationType>(NotificationType.Info);

  return (
    <Container maxWidth='xs'>
      <Typography variant='h4' align='center'>
        Send Global Notification
      </Typography>
      <Box
        sx={{ m: 2 }}
        display='flex'
        alignItems='center'
        flexDirection='column'
      >
        <TextField
          sx={{ m: 2 }}
          label='Message'
          value={message}
          onChange={(e) => setMessage(e.target.value)}
          fullWidth
        />
        <Select
          value={type}
          onChange={(e) => setType(e.target.value as NotificationType)}
          fullWidth
        >
          <MenuItem value={NotificationType.Info}>Info</MenuItem>
          <MenuItem value={NotificationType.Warning}>Warning</MenuItem>
          <MenuItem value={NotificationType.Error}>Error</MenuItem>
          <MenuItem value={NotificationType.Success}>Success</MenuItem>
          <MenuItem value={NotificationType.Default}>Default</MenuItem>
        </Select>
        <Box sx={{ m: 2 }} />
        <ButtonGroup variant='contained' fullWidth>
          <Button
            onClick={() => {
              sendGlobalNotification({
                variables: {
                  message: message,
                  type: type,
                },
              });
            }}
          >
            <Typography variant='h6'>Send</Typography>
          </Button>
          <Button
            onClick={() => {
              enqueueSnackbar(message, { variant: type });
            }}
          >
            <Typography variant='h6'>Test</Typography>
          </Button>
        </ButtonGroup>
      </Box>
    </Container>
  );
}
