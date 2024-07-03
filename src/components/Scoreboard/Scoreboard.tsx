import React from "react";

import {
  Paper,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Tooltip,
  Typography,
} from "@mui/material";

import { StatusEnum } from "../../graph";
import { ScoreboardData, ScoreboardTheme } from "../../models";

type props = {
  theme: "dark" | "light";
  scoreboardData: ScoreboardData;
  scoreboardTheme: ScoreboardTheme;
  cornerLabel?: string;
  highlightedRow: number | null;
  highlightedColumn: number | null;
  setHighlightedRow: React.Dispatch<React.SetStateAction<number | null>>;
  setHighlightedColumn: React.Dispatch<React.SetStateAction<number | null>>;
};

export default function Scoreboard({
  theme,
  scoreboardData,
  scoreboardTheme,
  cornerLabel,
  highlightedRow,
  highlightedColumn,
  setHighlightedRow,
  setHighlightedColumn,
}: props) {
  return (
    <TableContainer
      component={Paper}
      onMouseLeave={() => {
        setHighlightedRow(null);
        setHighlightedColumn(null);
      }}
      sx={{
        position: "relative",
      }}
    >
      <Table sx={{ width: "100%" }}>
        <TableHead>
          <TableRow>
            <TableCell
              size='small'
              onMouseEnter={() => {
                setHighlightedRow(0);
                setHighlightedColumn(0);
              }}
              sx={{
                position: "sticky",
                left: 0,
                backgroundColor:
                  scoreboardTheme.heading[theme][
                    highlightedRow === 0 || highlightedColumn === 0
                      ? "highlighted"
                      : "plain"
                  ],
              }}
            >
              {cornerLabel && (
                <Typography variant='body2' align='center'>
                  {cornerLabel}
                </Typography>
              )}
            </TableCell>
            {scoreboardData.top.map((heading, column) => (
              <TableCell
                size='small'
                key={`top-${column}`}
                onMouseEnter={() => {
                  setHighlightedColumn(column + 1);
                  setHighlightedRow(null);
                }}
                sx={{
                  backgroundColor:
                    scoreboardTheme.heading[theme][
                      highlightedColumn === column + 1 || highlightedRow === 0
                        ? "highlighted"
                        : "plain"
                    ],
                }}
              >
                <Typography variant='body2' align='center'>
                  {heading}
                </Typography>
              </TableCell>
            ))}
          </TableRow>
        </TableHead>
        <TableBody>
          {scoreboardData.left.map((heading, row) => (
            <TableRow key={`row-${row}`}>
              <TableCell
                size='small'
                onMouseEnter={() => {
                  setHighlightedColumn(null);
                  setHighlightedRow(row + 1);
                }}
                sx={{
                  whiteSpace: "nowrap",
                  position: "sticky",
                  left: 0,
                  backgroundColor:
                    scoreboardTheme.heading[theme][
                      highlightedColumn === 0 || highlightedRow === row + 1
                        ? "highlighted"
                        : "plain"
                    ],
                }}
              >
                <Typography variant='body2' align='center'>
                  {heading}
                </Typography>
              </TableCell>
              {scoreboardData.values[row].map((value, column) => (
                <Tooltip
                  key={`cell-${row}-${column}-tooltip`}
                  arrow={true}
                  title={
                    value?.__typename ? (
                      <>
                        <Typography variant='caption'>
                          Updated: {value?.update_time}
                        </Typography>{" "}
                        {value?.error && (
                          <Typography variant='caption'>
                            Error: {value?.error}
                          </Typography>
                        )}
                      </>
                    ) : (
                      <Typography variant='caption'>
                        Status either did not report yet or does not exist
                      </Typography>
                    )
                  }
                >
                  <TableCell
                    size='small'
                    key={`cell-${row}-${column}`}
                    sx={{
                      aspectRatio: 1,
                      backgroundColor:
                        scoreboardTheme.cell[theme][
                          highlightedRow === row + 1 ||
                          highlightedColumn === column + 1
                            ? "highlighted"
                            : "plain"
                        ][value?.status ?? StatusEnum.Unknown],
                    }}
                    onMouseEnter={() => {
                      setHighlightedRow(row + 1);
                      setHighlightedColumn(column + 1);
                    }}
                  />
                </Tooltip>
              ))}
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
}
