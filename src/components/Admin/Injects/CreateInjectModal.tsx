import { useState } from "react";

import { Close } from "@mui/icons-material";
import {
  Box,
  Button,
  FormControl,
  IconButton,
  Modal,
  Paper,
  TextField,
  Typography,
} from "@mui/material";
import { DateTimePicker, LocalizationProvider } from "@mui/x-date-pickers";
import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";
import { Dayjs } from "dayjs";
import { enqueueSnackbar } from "notistack";

import { RubricTemplateInput, useCreateInjectMutation } from "../../../graph";
import FileChip from "../../Common/FileChip";
import FileDrop from "../../Common/FileDrop";

type props = {
  open: boolean;
  setOpen: (isOpen: boolean) => void;
  handleRefetch: () => void;
};

export default function CreateCheckModal({
  open,
  setOpen,
  handleRefetch,
}: props) {
  const [createInjectMutation] = useCreateInjectMutation({
    onCompleted: () => {
      enqueueSnackbar("Inject created successfully", { variant: "success" });
      setOpen(false);
      handleRefetch();
    },
    onError: (error) => {
      enqueueSnackbar(error.message, { variant: "error" });
    },
  });

  const [name, setName] = useState<string>("");
  const [startTime, setStartTime] = useState<Dayjs | null>(null);
  const [endTime, setEndTime] = useState<Dayjs | null>(null);
  const [files, setFiles] = useState<File[] | null>(null);
  const [rubric, setRubric] = useState<RubricTemplateInput>({
    max_score: 0,
    fields: [],
  });

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

  const removeFile = (index: number) => {
    setFiles((prev) => {
      if (prev) {
        return prev.filter((_, i) => i !== index);
      } else {
        return null;
      }
    });
  };

  const handleCreateInject = () => {
    if (name === "" || !startTime || !endTime || !files) {
      enqueueSnackbar("Please fill out all fields", { variant: "error" });
      return;
    }

    createInjectMutation({
      variables: {
        title: name,
        start_time: startTime.toISOString(),
        end_time: endTime.toISOString(),
        files,
        rubric,
      },
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
            Create New Inject
          </Typography>
          <FormControl></FormControl>
          <LocalizationProvider dateAdapter={AdapterDayjs}>
            <Box
              sx={{
                marginTop: "24px",
                display: "flex",
                flexDirection: "column",
              }}
            >
              <TextField
                label='Name'
                variant='outlined'
                value={name}
                onChange={(event) => {
                  setName(event.target.value as string);
                }}
              />
              <DateTimePicker
                sx={{ marginTop: "24px" }}
                label='Start Time'
                value={startTime}
                onChange={(date) => {
                  setStartTime(date);
                }}
              />
              <DateTimePicker
                sx={{ marginTop: "24px" }}
                label='End Time'
                value={endTime}
                onChange={(date) => {
                  setEndTime(date);
                }}
              />
              <Paper
                sx={{
                  marginTop: "24px",
                  padding: "16px",
                  display: "flex",
                  flexDirection: "column",
                  gap: "16px",
                }}
                elevation={1}
              >
                {rubric.fields.map((field, i) => (
                  <Paper key={i} elevation={2}>
                    <Box
                      sx={{
                        display: "flex",
                        justifyContent: "space-between",
                        padding: "12px",
                        gap: "16px",
                      }}
                    >
                      <TextField
                        label='Field Name'
                        variant='outlined'
                        size='small'
                        value={field.name}
                        onChange={(e) => {
                          setRubric((prev) => ({
                            ...prev,
                            fields: prev.fields.map((f, index) =>
                              index === i ? { ...f, name: e.target.value } : f
                            ),
                          }));
                        }}
                        fullWidth
                      />
                      <TextField
                        label='Max Score'
                        variant='outlined'
                        size='small'
                        type='number'
                        value={field.max_score === 0 ? "" : field.max_score}
                        onChange={(e) => {
                          const newValue = e.target.value.replace(/^0+/, "");
                          const newScore = parseInt(newValue, 10) || 0;
                          setRubric((prev) => ({
                            max_score:
                              prev.max_score + newScore - field.max_score,
                            fields: prev.fields.map((f, index) =>
                              index === i ? { ...f, max_score: newScore } : f
                            ),
                          }));
                        }}
                        inputProps={{ inputMode: "numeric" }}
                      />
                      <IconButton
                        onClick={() => {
                          setRubric((prev) => ({
                            max_score: prev.max_score - field.max_score,
                            fields: prev.fields.filter(
                              (_, index) => index !== i
                            ),
                          }));
                        }}
                      >
                        <Close />
                      </IconButton>
                    </Box>
                  </Paper>
                ))}
                <Box sx={{ display: "flex", gap: "16px" }}>
                  <Button
                    variant='contained'
                    onClick={() => {
                      setRubric((prev) => ({
                        ...prev,
                        fields: [...prev.fields, { name: "", max_score: 0 }],
                      }));
                    }}
                    color='inherit'
                    fullWidth
                  >
                    Add New Field
                  </Button>
                  <TextField
                    label='Max Score'
                    variant='outlined'
                    size='small'
                    type='number'
                    value={rubric.max_score}
                    onChange={(e) => {
                      const newScore = parseInt(e.target.value, 10);
                      setRubric((prev) => ({
                        max_score: newScore,
                        fields: prev.fields,
                      }));
                    }}
                  />
                </Box>
              </Paper>
              <FileDrop onDrop={onDrop} onError={onError} />
              {files && files.length > 0 && (
                <Box
                  sx={{
                    display: "flex",
                    flexWrap: "wrap",
                    mt: "8px",
                    gap: "8px",
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
            </Box>
          </LocalizationProvider>
          <Button
            variant='contained'
            disabled={name === "" || !startTime || !endTime}
            onClick={handleCreateInject}
          >
            Create Inject
          </Button>
        </Box>
      </Box>
    </Modal>
  );
}
