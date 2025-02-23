import { useState } from "react";

import {
  Box,
  Button,
  FormControl,
  Modal,
  TextField,
  Typography,
} from "@mui/material";
import { enqueueSnackbar } from "notistack";

import { useCreateKothCheckMutation } from "../../../graph";

type props = {
  open: boolean;
  setOpen: (isOpen: boolean) => void;
  handleRefetch: () => void;
};

export default function CreateKothCheckModal({
  open,
  setOpen,
  handleRefetch,
}: props) {
  const [name, setName] = useState<string>("");
  const [weight, setWeight] = useState<number>(1);
  const [host, setHost] = useState<string>("");
  const [file, setFile] = useState<string>("");
  const [topic, setTopic] = useState<string>("");

  const [createKothCheckMutation] = useCreateKothCheckMutation({
    onCompleted: () => {
      enqueueSnackbar("Check created successfully", { variant: "success" });
      setOpen(false);
      handleRefetch();

      setName("");
      setWeight(1);
      setHost("");
      setFile("");
      setTopic("");
    },
    onError: (error) => {
      enqueueSnackbar(error.message, { variant: "error" });
    },
  });

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
            Create New KoTH Check
          </Typography>
          <FormControl sx={{ marginTop: "24px" }}>
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
              label='File'
              variant='outlined'
              sx={{
                marginTop: "24px",
              }}
              value={file}
              onChange={(event) => {
                setFile(event.target.value as string);
              }}
            />

            <TextField
              label='Hostname'
              variant='outlined'
              sx={{
                marginTop: "24px",
              }}
              value={host}
              onChange={(event) => {
                setHost(event.target.value as string);
              }}
            />

            <TextField
              label='KoTH Topic Key'
              variant='outlined'
              sx={{
                marginTop: "24px",
              }}
              value={host}
              onChange={(event) => {
                setTopic(event.target.value as string);
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
          </FormControl>
          <Box sx={{ marginTop: "24px", display: "flex", gap: "24px" }}>
            <Button
              variant='contained'
              color='primary'
              disabled={
                name === "" || file === "" || host === "" || topic === ""
              }
              onClick={() => {
                if (name === "") {
                  enqueueSnackbar("Name must be set", {
                    variant: "error",
                  });
                  return;
                }

                if (file === "") {
                  enqueueSnackbar("File must be set", {
                    variant: "error",
                  });
                  return;
                }

                if (host === "") {
                  enqueueSnackbar("Host must be set", {
                    variant: "error",
                  });
                  return;
                }

                if (topic === "") {
                  enqueueSnackbar("Topic must be set", {
                    variant: "error",
                  });
                  return;
                }

                createKothCheckMutation({
                  variables: {
                    name: name,
                    file: file,
                    host: host,
                    weight: weight,
                    topic: topic,
                  },
                });
              }}
            >
              Create KoTH Check
            </Button>
          </Box>
        </Box>
      </Box>
    </Modal>
  );
}
