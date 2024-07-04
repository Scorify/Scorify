import { useMemo } from "react";

import { Paper, Typography } from "@mui/material";

import { ScoreboardQuery } from "../../graph";

type props = {
  scores: ScoreboardQuery["scoreboard"]["scores"];
  highlightedTeam: number | null;
  setHighlightedTeam: React.Dispatch<React.SetStateAction<number | null>>;
};

export default function Scores({
  scores,
  highlightedTeam,
  setHighlightedTeam,
}: props) {
  const scoresFiltered = useMemo(() => {
    return scores.filter((score) => score !== null);
  }, [scores]);

  return (
    <Paper
      sx={{
        gap: 2,
        padding: 2,
        borderRadius: 2,
        marginTop: 5,
        display: "flex",
        justifyContent: "center",
        flexWrap: "wrap",
      }}
    >
      {scoresFiltered.map((score) => (
        <Paper
          key={score.user.number}
          sx={{
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
            paddingTop: 1,
            paddingBottom: 1,
            paddingLeft: 3,
            paddingRight: 3,
          }}
          elevation={highlightedTeam === score.user.number ? 9 : 3}
          onMouseEnter={() => {
            setHighlightedTeam(score.user.number || null);
          }}
          onMouseLeave={() => {
            setHighlightedTeam(null);
          }}
        >
          <Typography variant='body2'>{score.user.username}</Typography>
          <Typography variant='h6'>{score.score}</Typography>
        </Paper>
      ))}
    </Paper>
  );
}
