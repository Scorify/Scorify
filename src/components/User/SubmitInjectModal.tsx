import { useState } from "react";

import { Box, Button, Modal, TextField, Typography } from "@mui/material";
import { enqueueSnackbar } from "notistack";

import { InjectsQuery, useSubmitInjectMutation } from "../../graph";
import FileChip from "../Common/FileChip";
import FileDrop from "../Common/FileDrop";

type props = {
  inject: InjectsQuery["injects"][0];
  open: boolean;
  setOpen: (isOpen: boolean) => void;
  handleRefetch: () => void;
};

export default function SubmitInjectModal({
  inject,
  open,
  setOpen,
  handleRefetch,
}: props) {
  const [submitInjectMutation] = useSubmitInjectMutation({
    onCompleted: () => {
      enqueueSnackbar("Inject submitted successfully", { variant: "success" });
      setOpen(false);
      handleRefetch();
    },
    onError: (error) => {
      enqueueSnackbar(error.message, { variant: "error" });
    },
  });

  const [notes, setNotes] = useState("");
  const [files, setFiles] = useState<File[] | null>([]);

  const removeFile = (index: number) => {
    setFiles((prev) => {
      if (prev) {
        return prev.filter((_, i) => i !== index);
      } else {
        return null;
      }
    });
  };

  const onDrop = (files: File[]) => {
    setFiles((prev) => {
      if (prev) {
        return prev.concat(files);
      } else {
        return files;
      }
    });
  };

  const onError = (error: Error) => {
    enqueueSnackbar(error.message, { variant: "error" });
    console.error(error);
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
          <Typography
            component='h1'
            variant='h3'
            sx={{ mb: "12px" }}
            align='center'
          >
            New Inject Submission
          </Typography>
          <Typography
            component='h2'
            variant='h5'
            sx={{ mb: "24px" }}
            align='center'
          >
            {inject.title}
          </Typography>
          <TextField
            label='Notes'
            value={notes}
            multiline
            fullWidth
            rows={3}
            onChange={(e) => {
              setNotes(e.target.value);
            }}
            required
          />
        </Box>
        <FileDrop onDrop={onDrop} onError={onError} />
        {files && files.length > 0 && (
          <Box
            sx={{
              display: "flex",
              flexWrap: "wrap",
              mt: "12px",
              mb: files.length > 0 ? "12px" : "0px",
              justifyContent: "center",
            }}
          >
            {files.map((file, i) => (
              <FileChip
                key={`${file.name}-${i}`}
                file={file}
                onDelete={() => removeFile(i)}
              />
            ))}
          </Box>
        )}

        <Box sx={{ display: "flex", justifyContent: "center" }}>
          <Button
            variant='contained'
            disabled={files !== null && files.length === 0}
            onClick={() => {
              submitInjectMutation({
                variables: {
                  id: inject.id,
                  notes,
                  files,
                },
              });
            }}
            sx={{ alignSelf: "center" }}
          >
            Submit Inject
          </Button>
        </Box>
      </Box>
    </Modal>
  );
}
