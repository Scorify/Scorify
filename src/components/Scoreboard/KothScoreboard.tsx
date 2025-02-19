import React from "react";

import {
  Box,
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

import { KothScoreboardQuery, StatusEnum } from "../../graph";
import { ScoreboardTheme } from "../../models";

type props = {
  theme: "dark" | "light";
  scoreboardData: KothScoreboardQuery["kothScoreboard"];
  scoreboardTheme: ScoreboardTheme;
  highlightedRow: number | null;
  highlightedColumn: number | null;
  highlightedUser: number | null;
  setHighlightedRow: React.Dispatch<React.SetStateAction<number | null>>;
  setHighlightedColumn: React.Dispatch<React.SetStateAction<number | null>>;
  setHighlightedUser: React.Dispatch<React.SetStateAction<number | null>>;
};

export default function KothScoreboard({
  theme,
  scoreboardData,
  scoreboardTheme,
  highlightedRow,
  highlightedColumn,
  highlightedUser,
  setHighlightedRow,
  setHighlightedColumn,
  setHighlightedUser,
}: props) {
  return (
    <TableContainer
      component={Paper}
      onMouseLeave={() => {
        setHighlightedRow(null);
        setHighlightedColumn(null);
        setHighlightedUser(null);
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
              <Typography variant='body2' align='center'>
                Status
              </Typography>
            </TableCell>
            <TableCell
              size='small'
              onMouseEnter={() => {
                setHighlightedColumn(1);
                setHighlightedRow(null);
              }}
              sx={{
                backgroundColor:
                  scoreboardTheme.heading[theme][
                    highlightedColumn === 1 || highlightedRow === 0
                      ? "highlighted"
                      : "plain"
                  ],
              }}
            >
              <Typography variant='body2' align='center'>
                Check
              </Typography>
            </TableCell>
            <TableCell
              size='small'
              onMouseEnter={() => {
                setHighlightedColumn(2);
                setHighlightedRow(null);
              }}
              sx={{
                backgroundColor:
                  scoreboardTheme.heading[theme][
                    highlightedColumn === 2 || highlightedRow === 0
                      ? "highlighted"
                      : "plain"
                  ],
              }}
            >
              <Typography variant='body2' align='center'>
                Owner
              </Typography>
            </TableCell>
            <TableCell
              size='small'
              onMouseEnter={() => {
                setHighlightedColumn(3);
                setHighlightedRow(null);
              }}
              sx={{
                backgroundColor:
                  scoreboardTheme.heading[theme][
                    highlightedColumn === 3 || highlightedRow === 0
                      ? "highlighted"
                      : "plain"
                  ],
              }}
            >
              <Typography variant='body2' align='center'>
                Host
              </Typography>
            </TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {scoreboardData.checks.map((check, row) => (
            <TableRow key={`row-${row}`}>
              <Tooltip
                arrow={true}
                title={
                  <Box display='flex' flexDirection='column'>
                    <Typography variant='caption'>
                      Updated: {check.update_time}
                    </Typography>{" "}
                    {check.error && (
                      <Typography variant='caption'>
                        Error: {check.error}
                      </Typography>
                    )}
                  </Box>
                }
              >
                <TableCell
                  size='small'
                  onMouseEnter={() => {
                    setHighlightedColumn(null);
                    setHighlightedRow(row + 1);
                    setHighlightedUser(check.user?.number || null);
                  }}
                  sx={{
                    whiteSpace: "nowrap",
                    position: "sticky",
                    left: 0,
                    backgroundColor:
                      scoreboardTheme.cell[theme][
                        highlightedRow === row + 1 ||
                        highlightedColumn === 0 ||
                        highlightedUser === check.user?.number
                          ? "highlighted"
                          : "plain"
                      ][check.user ? StatusEnum.Up : StatusEnum.Down],
                  }}
                />
              </Tooltip>
              <TableCell
                size='small'
                onMouseEnter={() => {
                  setHighlightedColumn(1);
                  setHighlightedRow(row + 1);
                  setHighlightedUser(check.user?.number || null);
                }}
                sx={{
                  backgroundColor:
                    scoreboardTheme.heading[theme][
                      highlightedColumn === 1 ||
                      highlightedRow === row + 1 ||
                      highlightedUser === check.user?.number
                        ? "highlighted"
                        : "plain"
                    ],
                }}
              >
                <Typography variant='body2' align='center'>
                  {check.name}
                </Typography>
              </TableCell>
              <TableCell
                size='small'
                onMouseEnter={() => {
                  setHighlightedColumn(2);
                  setHighlightedRow(row + 1);
                  setHighlightedUser(check.user?.number || null);
                }}
                sx={{
                  backgroundColor:
                    scoreboardTheme.heading[theme][
                      highlightedColumn === 2 ||
                      highlightedRow === row + 1 ||
                      highlightedUser === check.user?.number
                        ? "highlighted"
                        : "plain"
                    ],
                }}
              >
                <Typography variant='body2' align='center'>
                  {check.user ? check.user.username : "Unclaimed"}
                </Typography>
              </TableCell>
              <TableCell
                size='small'
                onMouseEnter={() => {
                  setHighlightedColumn(3);
                  setHighlightedRow(row + 1);
                  setHighlightedUser(check.user?.number || null);
                }}
                sx={{
                  backgroundColor:
                    scoreboardTheme.heading[theme][
                      highlightedColumn === 3 ||
                      highlightedRow === row + 1 ||
                      highlightedUser === check.user?.number
                        ? "highlighted"
                        : "plain"
                    ],
                }}
              >
                <Typography variant='body2' align='center'>
                  {check.host || "Unknown"}
                </Typography>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
}
