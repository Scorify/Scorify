import React, { useMemo, useState } from "react";

import { Close } from "@mui/icons-material";
import {
  Box,
  Button,
  Chip,
  Divider,
  IconButton,
  Modal,
  Paper,
  TextField,
  Typography,
} from "@mui/material";
import { DateTimePicker, LocalizationProvider } from "@mui/x-date-pickers";
import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";

import dayjs, { Dayjs } from "dayjs";
import { enqueueSnackbar } from "notistack";
import {
  DeleteInjectModal,
  Dropdown,
  FileChip,
  FileDrop,
  Loading,
} from "../..";
import {
  InjectsQuery,
  RubricInput,
  RubricTemplateInput,
  SubmissionsQuery,
  useDeleteInjectMutation,
  useGradeSubmissionMutation,
  useSubmissionsQuery,
  useUpdateInjectMutation,
} from "../../../graph";

type props = {
  inject: InjectsQuery["injects"][0];
  handleRefetch: () => void;
  visible: boolean;
};

export default function EditInject({ inject, handleRefetch, visible }: props) {
  const [expanded, setExpanded] = useState(false);
  const [open, setOpen] = useState(false);

  const [startTime, setStartTime] = useState<Dayjs | null>(
    dayjs(inject.start_time)
  );
  const startTimeChanged = useMemo(
    () => startTime?.toISOString() != dayjs(inject.start_time).toISOString(),
    [startTime, inject.start_time]
  );

  const [endTime, setEndTime] = useState<Dayjs | null>(dayjs(inject.end_time));
  const endTimeChanged = useMemo(
    () => endTime?.toISOString() != dayjs(inject.end_time).toISOString(),
    [endTime, inject.end_time]
  );

  const [title, setTitle] = useState<string>(inject.title);
  const titleChanged = useMemo(
    () => title != inject.title,
    [title, inject.title]
  );

  const [deleteFiles, setDeleteFiles] = useState<string[]>([]);
  const [newFiles, setNewFiles] = useState<File[]>([]);
  const filesChanged = useMemo(
    () => deleteFiles.length > 0 || newFiles.length > 0,
    [deleteFiles, newFiles]
  );

  const [rubric, setRubric] = useState<RubricTemplateInput>({
    max_score: inject.rubric.max_score,
    fields: inject.rubric.fields.map((field) => ({
      name: field.name,
      max_score: field.max_score,
    })),
  });
  const rubricChanged = useMemo(
    () =>
      JSON.stringify(rubric) !=
      JSON.stringify({
        max_score: inject.rubric.max_score,
        fields: inject.rubric.fields.map((field) => ({
          name: field.name,
          max_score: field.max_score,
        })),
      }),
    [rubric, inject]
  );

  const [updateInjectMutation] = useUpdateInjectMutation({
    onCompleted: () => {
      enqueueSnackbar("Inject updated successfully", { variant: "success" });
      handleRefetch();
    },
    onError: (error) => {
      enqueueSnackbar(error.message, { variant: "error" });
    },
  });

  const [deleteInjectMutation] = useDeleteInjectMutation({
    onCompleted: () => {
      enqueueSnackbar("Inject deleted successfully", { variant: "success" });
      handleRefetch();
    },
    onError: (error) => {
      enqueueSnackbar(error.message, { variant: "error" });
    },
  });

  const [grading, setGrading] = useState(false);

  const handleSave = () => {
    if (
      !titleChanged &&
      !startTimeChanged &&
      !endTimeChanged &&
      !filesChanged &&
      !rubricChanged
    ) {
      return;
    }
    updateInjectMutation({
      variables: {
        id: inject.id,
        title: titleChanged ? title : undefined,
        start_time: startTimeChanged ? startTime?.toISOString() : undefined,
        end_time: endTimeChanged ? endTime?.toISOString() : undefined,
        add_files: newFiles.length > 0 ? newFiles : undefined,
        delete_files: deleteFiles.length > 0 ? deleteFiles : undefined,
        rubric: rubricChanged ? rubric : undefined,
      },
    });
    setDeleteFiles([]);
    setNewFiles([]);
  };

  const handleDelete = () => {
    deleteInjectMutation({
      variables: {
        id: inject.id,
      },
    });
  };

  return (
    <Dropdown
      elevation={1}
      modal={
        <DeleteInjectModal
          inject={inject.title}
          open={open}
          setOpen={setOpen}
          handleDelete={handleDelete}
        />
      }
      title={
        expanded && !grading ? (
          <TextField
            label='Name'
            value={title}
            onClick={(e) => {
              e.stopPropagation();
            }}
            onChange={(e) => {
              setTitle(e.target.value);
            }}
            sx={{ marginRight: "24px" }}
            size='small'
          />
        ) : (
          <Typography variant='h6' component='div' marginRight='24px'>
            {inject.title}
          </Typography>
        )
      }
      expandableButtons={[
        <Button
          variant='contained'
          onClick={(e) => {
            setGrading((prev) => !prev);
            e.stopPropagation();
          }}
          color='info'
        >
          {grading ? "Switch to Editting" : "Switch to Grading"}
        </Button>,
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
      toggleButtonVisible={
        titleChanged ||
        startTimeChanged ||
        endTimeChanged ||
        filesChanged ||
        rubricChanged
      }
      visible={visible}
      expanded={expanded}
      setExpanded={setExpanded}
    >
      {grading ? (
        <GradeInjectPanel inject={inject} handleRefetch={handleRefetch} />
      ) : (
        <EditInjectPanel
          rubric={rubric}
          setRubric={setRubric}
          startTime={startTime}
          setStartTime={setStartTime}
          endTime={endTime}
          setEndTime={setEndTime}
          newFiles={newFiles}
          setNewFiles={setNewFiles}
          deleteFiles={deleteFiles}
          setDeleteFiles={setDeleteFiles}
          inject={inject}
        />
      )}
    </Dropdown>
  );
}

type GradeSubmissonModalProps = {
  open: boolean;
  setOpen: React.Dispatch<React.SetStateAction<boolean>>;
  submission: SubmissionsQuery["injectSubmissionsByUser"][0]["submissions"][0];
  handleRefetch: () => void;
};

function GradeSubmissonModal({
  open,
  setOpen,
  submission,
  handleRefetch,
}: GradeSubmissonModalProps) {
  const [gradeSubmissionMutation] = useGradeSubmissionMutation({
    onCompleted: () => {
      enqueueSnackbar("Submission graded successfully", { variant: "success" });
      handleRefetch();
      setOpen(false);
    },
    onError: (error) => {
      enqueueSnackbar(error.message, { variant: "error" });
    },
  });

  const [rubricInput, setRubricInput] = useState<RubricInput>({
    fields: submission.inject.rubric.fields.map((field) => ({
      name: field.name,
      score:
        submission.rubric?.fields.find((f) => f.name === field.name)?.score ??
        0,
      notes:
        submission.rubric?.fields.find((f) => f.name === field.name)?.notes ??
        "",
    })),
    notes: submission.rubric?.notes ?? "",
  });

  const submitGrade = () => {
    gradeSubmissionMutation({
      variables: {
        submission_id: submission.id,
        rubric: rubricInput,
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
          <Typography variant='h4' align='center' onClick={submitGrade}>
            Grade Submission
          </Typography>
          <Paper
            sx={{
              display: "flex",
              flexDirection: "column",
              gap: "16px",
              width: "100%",
              marginTop: "16px",
              padding: "16px",
            }}
            elevation={1}
          >
            {rubricInput.fields.map((field, i) => (
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
                  />
                  <TextField
                    label='Notes'
                    variant='outlined'
                    size='small'
                    value={field.notes}
                    onChange={(e) => {
                      setRubricInput((prev) => ({
                        ...prev,
                        fields: prev.fields.map((f, index) =>
                          index === i ? { ...f, notes: e.target.value } : f
                        ),
                      }));
                    }}
                    fullWidth
                  />
                  <TextField
                    label='Score'
                    variant='outlined'
                    size='small'
                    type='number'
                    value={
                      field.score.toString().replace(/^0+/, "") === ""
                        ? "0"
                        : field.score.toString().replace(/^0+/, "")
                    }
                    onChange={(e) => {
                      const newScore = parseInt(e.target.value, 10);

                      if (isNaN(newScore)) {
                        setRubricInput((prev) => ({
                          ...prev,
                          fields: prev.fields.map((f, index) =>
                            index === i ? { ...f, score: 0 } : f
                          ),
                        }));
                        return;
                      }

                      const maxScore =
                        submission.inject.rubric.fields.find(
                          (f) => f.name === field.name
                        )?.max_score ?? 0;

                      setRubricInput((prev) => ({
                        ...prev,
                        fields: prev.fields.map((f, index) =>
                          index === i
                            ? {
                                ...f,
                                score:
                                  newScore >= maxScore ? maxScore : newScore,
                              }
                            : f
                        ),
                      }));
                    }}
                    inputProps={{ inputMode: "numeric" }}
                  />
                  <TextField
                    label='Max Score'
                    variant='outlined'
                    size='small'
                    value={
                      submission.inject.rubric.fields.find(
                        (f) => f.name === field.name
                      )?.max_score ?? 0
                    }
                  />
                </Box>
              </Paper>
            ))}
            <Divider />
            <Paper
              sx={{
                display: "flex",
                flexDirection: "column",
                padding: "12px",
                gap: "16px",
              }}
              elevation={2}
            >
              <TextField
                label='Notes'
                variant='outlined'
                size='small'
                value={rubricInput.notes}
                onChange={(e) => {
                  setRubricInput((prev) => ({
                    ...prev,
                    notes: e.target.value,
                  }));
                }}
                multiline
                rows={4}
                fullWidth
              />
              <Box
                sx={{
                  display: "flex",
                  flexDirection: "row",
                  justifyContent: "space-between",
                  gap: "16px",
                }}
              >
                <TextField
                  label='Total Score'
                  variant='outlined'
                  size='small'
                  value={rubricInput.fields.reduce((total, field) => {
                    return total + field.score;
                  }, 0)}
                  fullWidth
                />
                <TextField
                  label='Max Score'
                  variant='outlined'
                  size='small'
                  value={submission.inject.rubric.max_score}
                  fullWidth
                />
                <Button
                  variant='contained'
                  color='success'
                  onClick={submitGrade}
                  fullWidth
                >
                  Submit Grade
                </Button>
              </Box>
            </Paper>
          </Paper>
        </Box>
      </Box>
    </Modal>
  );
}

type SubmissionPanelProps = {
  submission: SubmissionsQuery["injectSubmissionsByUser"][0]["submissions"][0];
  title: string;
  handleRefetch: () => void;
};

function SubmissionPanel({
  submission,
  title,
  handleRefetch,
}: SubmissionPanelProps) {
  const [expanded, setExpanded] = useState(false);
  const [open, setOpen] = useState(false);

  const date = new Date(submission.create_time);

  return (
    <Dropdown
      elevation={3}
      modal={
        <GradeSubmissonModal
          open={open}
          setOpen={setOpen}
          submission={submission}
          handleRefetch={handleRefetch}
        />
      }
      title={
        <>
          <Typography variant='h6' component='div'>
            {`${title} - ${date.toLocaleDateString()} ${date.toLocaleTimeString()}`}
          </Typography>
          <Chip
            size='small'
            label={`${submission.files.length} ${
              submission.files.length === 1 ? "File" : "Files"
            }`}
          />
          {submission.graded && (
            <Chip
              size='small'
              label={`Graded: ${submission.rubric?.fields.reduce(
                (total, field) => {
                  return total + field.score;
                },
                0
              )} / ${submission.inject.rubric.max_score}`}
              color='success'
            />
          )}
        </>
      }
      expandableButtons={[
        <Button
          variant='contained'
          color='success'
          onClick={() => {
            setOpen(true);
          }}
        >
          Grade
        </Button>,
      ]}
      expanded={expanded}
      setExpanded={setExpanded}
    >
      {submission.notes && (
        <TextField
          label='Notes'
          value={submission.notes}
          multiline
          fullWidth
          sx={{ marginBottom: "8px" }}
        />
      )}
      <Box
        sx={{
          display: "flex",
          flexDirection: "row",
          gap: "8px",
          flexWrap: "wrap",
        }}
      >
        {submission.files.map((file) => (
          <FileChip key={file.id} file={file} />
        ))}
      </Box>
    </Dropdown>
  );
}

type TeamSubmissionsPanelProps = {
  user: SubmissionsQuery["injectSubmissionsByUser"][0]["user"];
  submissions: SubmissionsQuery["injectSubmissionsByUser"][0]["submissions"];
  handleRefetch: () => void;
};

function TeamSubmissionsPanel({
  user,
  submissions,
  handleRefetch,
}: TeamSubmissionsPanelProps) {
  const [expanded, setExpanded] = useState(false);

  const highestScore =
    submissions.filter((submission) => submission.graded).length === 0
      ? undefined
      : submissions
          .filter((submission) => submission.graded)
          .sort((a, b) => {
            return (
              (b.rubric?.fields.reduce(
                (total, field) => total + field.score,
                0
              ) ?? 0) -
              (a.rubric?.fields.reduce(
                (total, field) => total + field.score,
                0
              ) ?? 0)
            );
          })[0]
          .rubric?.fields.reduce((total, field) => total + field.score, 0) ??
        undefined;

  return (
    <Dropdown
      elevation={2}
      title={
        <>
          <Typography variant='h6' component='div'>
            {user.username}
          </Typography>
          <Chip
            label={`${submissions.length} ${
              submissions.length === 1 ? "Submission" : "Submissions"
            }`}
            size='small'
            color={submissions.length == 0 ? "error" : "success"}
          />
          {highestScore !== undefined &&
            submissions.filter((submission) => submission.graded).length >
              0 && (
              <Chip
                label={`Graded: ${highestScore}/${submissions[0].inject.rubric.max_score}`}
                color='success'
                size='small'
              />
            )}
        </>
      }
      expanded={expanded}
      setExpanded={setExpanded}
    >
      {submissions.length == 0 ? (
        <Typography variant='h6' align='center'>
          No Submissions
        </Typography>
      ) : (
        submissions.map((submission, i) => (
          <SubmissionPanel
            key={i}
            submission={submission}
            title={`Submission ${submissions.length - i}`}
            handleRefetch={handleRefetch}
          />
        ))
      )}
    </Dropdown>
  );
}

type GradeInjectPanelProps = {
  inject: InjectsQuery["injects"][0];
  handleRefetch: () => void;
};

function GradeInjectPanel({ inject, handleRefetch }: GradeInjectPanelProps) {
  const { data, loading, error } = useSubmissionsQuery({
    variables: {
      inject_id: inject.id,
    },
  });

  if (loading) {
    return <Loading />;
  }

  if (error) {
    return (
      <Typography variant='h6' color='error'>
        {error.message}
      </Typography>
    );
  }

  return (
    <>
      {data?.injectSubmissionsByUser.map(({ user, submissions }) => (
        <TeamSubmissionsPanel
          key={user.number}
          user={user}
          submissions={submissions}
          handleRefetch={handleRefetch}
        />
      ))}
    </>
  );
}

type EditInjectPanelProps = {
  rubric: RubricTemplateInput;
  setRubric: React.Dispatch<React.SetStateAction<RubricTemplateInput>>;
  startTime: Dayjs | null;
  setStartTime: React.Dispatch<React.SetStateAction<Dayjs | null>>;
  endTime: Dayjs | null;
  setEndTime: React.Dispatch<React.SetStateAction<Dayjs | null>>;
  newFiles: File[];
  setNewFiles: React.Dispatch<React.SetStateAction<File[]>>;
  deleteFiles: string[];
  setDeleteFiles: React.Dispatch<React.SetStateAction<string[]>>;
  inject: InjectsQuery["injects"][0];
};

function EditInjectPanel({
  rubric,
  setRubric,
  startTime,
  setStartTime,
  endTime,
  setEndTime,
  newFiles,
  setNewFiles,
  deleteFiles,
  setDeleteFiles,
  inject,
}: EditInjectPanelProps) {
  const onDrop = (files: File[]) => {
    setNewFiles((prev) => {
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
    <>
      <LocalizationProvider dateAdapter={AdapterDayjs}>
        <Box
          sx={{
            display: "flex",
            gap: "16px",
            flexWrap: "wrap",
            justifyContent: "center",
          }}
        >
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
        </Box>
      </LocalizationProvider>
      <Paper
        sx={{
          marginTop: "24px",
          padding: "16px",
          display: "flex",
          flexDirection: "column",
          gap: "16px",
        }}
        elevation={2}
      >
        {rubric.fields.map((field, i) => (
          <Paper key={i} elevation={3}>
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
                    max_score: prev.max_score + newScore - field.max_score,
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
                    fields: prev.fields.filter((_, index) => index !== i),
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
      <FileDrop onDrop={onDrop} onError={onError} elevation={4} />
      {(newFiles.length > 0 || (inject.files && inject.files.length > 0)) && (
        <Box
          sx={{
            display: "flex",
            flexWrap: "wrap",
            mt: "8px",
            gap: "8px",
          }}
        >
          {inject.files.map((file) => (
            <FileChip
              key={file.id}
              file={file}
              color={deleteFiles.includes(file.id) ? "error" : "default"}
              onDelete={() => {
                if (deleteFiles.includes(file.id)) {
                  setDeleteFiles((prev) => prev.filter((id) => id != file.id));
                  return;
                }
                setDeleteFiles((prev) => [...prev, file.id]);
              }}
            />
          ))}
          {newFiles.map((file, i) => (
            <FileChip
              key={`${file.name}-${i}`}
              file={file}
              color='success'
              onDelete={() => {
                setNewFiles((prev) => prev.filter((_, index) => i != index));
              }}
            />
          ))}
        </Box>
      )}
    </>
  );
}
