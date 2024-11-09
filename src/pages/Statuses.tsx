import { useNavigate } from "react-router-dom";
import { useState } from "react";

import { Box, Container, Typography } from "@mui/material";

import { useURLParam } from "../hooks";
import { StatusEnum, useStatusesQuery } from "../graph";

export default function Statuses() {
  const navigate = useNavigate();

  const urlSearchParams = new URLSearchParams(location.search);
  const [fromTime, setFromTime] = useURLParam<Date>(
    urlSearchParams,
    "fromTime",
    (date) => date.toISOString(),
    (s) => new Date(s)
  );
  const [toTime, setToTime] = useURLParam<Date>(
    urlSearchParams,
    "toTime",
    (date) => date.toISOString(),
    (s) => new Date(s)
  );
  const [fromRound, setFromRound] = useState(
    parseInt(urlSearchParams.get("fromRound") || "") || undefined
  );
  const [toRound, setToRound] = useState(
    parseInt(urlSearchParams.get("toRound") || "") || undefined
  );
  const [minions, setMinions] = useURLParam<string[]>(
    urlSearchParams,
    "minions",
    JSON.stringify,
    JSON.parse
  );
  const [checks, setChecks] = useURLParam<string[]>(
    urlSearchParams,
    "checks",
    JSON.stringify,
    JSON.parse
  );
  const [teams, setTeams] = useURLParam<string[]>(
    urlSearchParams,
    "teams",
    JSON.stringify,
    JSON.parse
  );
  const [statuses, setStatuses] = useURLParam<StatusEnum[]>(
    urlSearchParams,
    "statuses",
    JSON.stringify,
    (s) => {
      return JSON.parse(s).filter((status: string) => {
        return Object.values(StatusEnum).includes(status as StatusEnum);
      });
    }
  );
  const [limit, setLimit] = useURLParam<number>(
    urlSearchParams,
    "limit",
    String,
    parseInt
  );
  const [offset, setOffset] = useURLParam<number>(
    urlSearchParams,
    "offset",
    String,
    parseInt
  );

  const {} = useStatusesQuery({
    variables: {
      statusesInputQuery: {
        from_time: fromTime,
        to_time: toTime,
        from_round: fromRound,
        to_round: toRound,
        minions,
        checks,
        users: teams,
        statuses,
        limit,
        offset,
      },
    },
  });

  return (
    <Container component='main' maxWidth='md'>
      <Box
        sx={{
          marginTop: 8,
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
        }}
      >
        <Typography component='h1' variant='h3' sx={{ marginBottom: "24px" }}>
          Status Query
        </Typography>
      </Box>
    </Container>
  );
}
