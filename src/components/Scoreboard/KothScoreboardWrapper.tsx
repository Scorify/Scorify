import { useState } from "react";

import { Box } from "@mui/material";

import { KothScoreboardQuery } from "../../graph";
import { ScoreboardTheme } from "../../models";
import KothScoreboard from "./KothScoreboard";
import Scores from "./Scores";

type props = {
  theme: "dark" | "light";
  data: KothScoreboardQuery["kothScoreboard"];
  scoreboardTheme: ScoreboardTheme;
};

export default function KothScoreboardWrapper({
  data,
  theme,
  scoreboardTheme,
}: props) {
  const [highlightedCheck, setHighlightedCheck] = useState<number | null>(null);
  const [highlightedTeam, setHighlightedTeam] = useState<number | null>(null);
  const [highlightedUser, setHighlightedUser] = useState<number | null>(null);

  return (
    <Box>
      <KothScoreboard
        theme={theme}
        scoreboardData={data}
        scoreboardTheme={scoreboardTheme}
        highlightedRow={highlightedCheck}
        highlightedColumn={highlightedTeam}
        highlightedUser={highlightedUser}
        setHighlightedRow={setHighlightedCheck}
        setHighlightedColumn={setHighlightedTeam}
        setHighlightedUser={setHighlightedUser}
      />
      <Scores
        scores={data.scores}
        highlightedTeam={highlightedUser}
        setHighlightedTeam={setHighlightedUser}
      />
    </Box>
  );
}
