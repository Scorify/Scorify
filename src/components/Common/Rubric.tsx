import { Box, Paper, TextField, Divider } from "@mui/material";

import { InjectsQuery } from "../../graph";

type props = {
  elevation?: number;
  rubric: InjectsQuery["injects"][0]["rubric"];
  submission?: InjectsQuery["injects"][0]["submissions"][0];
};

export default function Rubric({ elevation = 1, rubric, submission }: props) {
  return (
    <Paper
      elevation={elevation}
      sx={{
        display: "flex",
        flexDirection: "column",
        justifyContent: "center",
        alignItems: "center",
        padding: "16px",
        marginTop: "8px",
        marginBottom: "24px",
      }}
    >
      {rubric.fields.length && (
        <Box display='flex' flexDirection='column' gap='8px' width='100%'>
          {rubric.fields.map((field) => (
            <Paper
              key={field.name}
              elevation={elevation + 1}
              sx={{
                display: "flex",
                flexDirection: "row",
                justifyContent: "space-between",
                alignItems: "center",
                padding: "12px",
                width: "100%",
                marginBottom: "4px",
                gap: "16px",
              }}
            >
              <TextField
                size='small'
                label='Criteria'
                value={field.name}
                fullWidth={!submission}
              />
              {submission && (
                <>
                  <TextField
                    size='small'
                    label='Notes'
                    value={submission.notes}
                    fullWidth
                  />
                  <TextField
                    size='small'
                    label='Score'
                    value={
                      submission.rubric?.fields.find(
                        (f) => f.name === field.name
                      )?.score
                    }
                  />
                </>
              )}
              <TextField
                size='small'
                label='Max Points'
                value={field.max_score}
              />
            </Paper>
          ))}
          {submission && (
            <>
              <Divider sx={{ marginBottom: "12px" }} />
              <TextField
                size='small'
                label='Notes'
                value={submission.rubric?.notes}
                fullWidth
              />
            </>
          )}
          <Box display='flex' flexDirection='row' gap='8px' marginTop='8px'>
            {submission && (
              <TextField
                size='small'
                label='Total Score'
                value={submission.rubric?.fields.reduce(
                  (acc, field) => acc + field.score,
                  0
                )}
                fullWidth
              />
            )}
            <TextField
              size='small'
              label='Max Score'
              value={rubric.max_score}
              fullWidth
            />
          </Box>
        </Box>
      )}
    </Paper>
  );
}
