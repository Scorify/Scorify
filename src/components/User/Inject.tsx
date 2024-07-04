import { useEffect, useState } from "react";

import { Box, Button, Chip, Typography } from "@mui/material";

import { InjectsQuery } from "../../graph";
import {
  Dropdown,
  FileChip,
  Rubric,
  Submission,
  SubmitInjectModal,
} from "../index";

type countdownChipProps = {
  target: number;
};

function CountdownChip({ target }: countdownChipProps) {
  const [difference, setDifference] = useState(target - Date.now() - 5);

  useEffect(() => {
    const interval = setInterval(() => {
      setDifference(target - Date.now() - 5);
    }, 1000);

    return () => clearInterval(interval);
  }, [target]);

  const handleLabel = (durationMs: number) => {
    let duration = Math.floor(durationMs / 1000);
    if (difference < 0) {
      if (duration < 60) {
        return `Closed ${Math.floor(duration)} seconds ago`;
      } else if (duration < 90 * 60) {
        return `Closed ${Math.floor(duration / 60)} minutes ago`;
      } else {
        return `Closed ${(duration / 3600).toFixed(1)} hours ago`;
      }
    } else {
      if (duration < 60) {
        return `Closes in ${Math.floor(duration)} seconds`;
      } else if (duration < 90 * 60) {
        return `Closes in ${Math.floor(duration / 60)} minutes`;
      } else {
        return `Closes in ${(duration / 3600).toFixed(1)} hours`;
      }
    }
  };

  return (
    <Chip
      color={difference < 0 ? "error" : "success"}
      label={handleLabel(Math.abs(difference))}
      size='small'
    />
  );
}

type props = {
  inject: InjectsQuery["injects"][0];
  handleRefetch: () => void;
  visible: boolean;
};

export default function Inject({ handleRefetch, inject, visible }: props) {
  const [expanded, setExpanded] = useState(false);
  const [open, setOpen] = useState(false);

  const sortedSubmissions = inject.submissions
    .filter((submission) => submission.graded)
    .sort(
      (submissionA, submissionB) =>
        (submissionB.rubric?.fields.reduce(
          (acc, field) => acc + field.score,
          0
        ) ?? 0) -
        (submissionA.rubric?.fields.reduce(
          (acc, field) => acc + field.score,
          0
        ) ?? 0)
    );

  return (
    <Dropdown
      modal={
        <SubmitInjectModal
          inject={inject}
          handleRefetch={handleRefetch}
          open={open}
          setOpen={setOpen}
        />
      }
      title={
        <>
          <Typography variant='h6' component='div' marginRight='24px'>
            {inject.title}
          </Typography>
          <CountdownChip target={new Date(inject.end_time).getTime()} />
          {sortedSubmissions.length > 0 && (
            <Chip
              label={`Score: ${
                sortedSubmissions[0].rubric?.fields.reduce(
                  (acc, field) => acc + field.score,
                  0
                ) ?? 0
              }/${inject.rubric.max_score}`}
              color='success'
              size='small'
              sx={{ marginLeft: "12px" }}
            />
          )}
        </>
      }
      expandableButtons={[
        <Button
          variant='contained'
          color='success'
          onClick={(e) => {
            if (!expanded) {
              e.stopPropagation();
            }

            setOpen(true);
          }}
        >
          Submit
        </Button>,
      ]}
      expanded={expanded}
      setExpanded={setExpanded}
      visible={visible}
    >
      <Typography variant='h4' align='center'>
        Files
      </Typography>
      <Box
        sx={{
          display: "flex",
          flexWrap: "wrap",
          justifyContent: "center",
          mt: "8px",
          gap: "8px",
          mb: "24px",
        }}
      >
        {inject.files.length ? (
          inject.files.map((file) => <FileChip key={file.id} file={file} />)
        ) : (
          <Typography variant='body1' align='center'>
            No Files
          </Typography>
        )}
      </Box>
      <Typography variant='h4' align='center'>
        Rubric
      </Typography>
      <Rubric
        submission={sortedSubmissions[0]}
        rubric={inject.rubric}
        elevation={2}
      />
      <Typography variant='h4' align='center'>
        Submissions
      </Typography>
      <Box
        sx={{
          display: "flex",
          flexWrap: "wrap",
          justifyContent: "center",
          mt: "8px",
          gap: "8px",
        }}
      >
        {inject.submissions.length ? (
          <>
            {[...inject.submissions]
              .sort(
                (a, b) =>
                  new Date(b.create_time).getTime() -
                  new Date(a.create_time).getTime()
              )
              .map((submission) => (
                <Submission
                  key={submission.id}
                  inject={inject}
                  submission={submission}
                />
              ))}
          </>
        ) : (
          <Typography variant='body1' align='center'>
            No Submissions
          </Typography>
        )}
      </Box>
    </Dropdown>
  );
}
