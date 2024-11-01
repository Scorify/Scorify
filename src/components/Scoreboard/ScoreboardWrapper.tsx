import { useMemo, useState } from "react";

import { Box } from "@mui/material";

import { ScoreboardQuery } from "../../graph";
import { ScoreboardTheme } from "../../models";
import Scoreboard from "./Scoreboard";
import Scores from "./Scores";

type props = {
  theme: "dark" | "light";
  data: ScoreboardQuery["scoreboard"];
  cornerLabel?: string;
  scoreboardTheme: ScoreboardTheme;
  displays: Map<number, { [key: string]: string }>;
};

export default function ScoreboardWrapper({
  data,
  theme,
  displays,
  scoreboardTheme,
  cornerLabel,
}: props) {
  const scoreboardData = useMemo(() => {
    return {
      top: data.teams.map((team) => team.number),
      left: data.checks.map((check) => check.name),
      values: data.statuses,
    };
  }, [data]);

  const [highlightedCheck, setHighlightedCheck] = useState<number | null>(null);
  const [highlightedTeam, setHighlightedTeam] = useState<number | null>(null);

  return (
    <Box>
      <Scoreboard
        theme={theme}
        scoreboardData={scoreboardData}
        scoreboardTheme={scoreboardTheme}
        displays={displays}
        cornerLabel={cornerLabel}
        highlightedRow={highlightedCheck}
        highlightedColumn={highlightedTeam}
        setHighlightedRow={setHighlightedCheck}
        setHighlightedColumn={setHighlightedTeam}
      />
      <Scores
        scores={data.scores}
        highlightedTeam={highlightedTeam}
        setHighlightedTeam={setHighlightedTeam}
      />
    </Box>
  );
}
