import { useNavigate } from "react-router-dom";
import { useEffect, useState } from "react";

import { Box, Container, Typography } from "@mui/material";
import { StatusEnum, useStatusesQuery } from "../graph";

interface ConvertToString<U> {
  (value: U): string;
}

interface ConvertFromString<U> {
  (value: string): U;
}

function useURLParam<U>(
  urlSearchParams: URLSearchParams,
  key: string,
  convertToString: ConvertToString<U>,
  convertFromString: ConvertFromString<U>
): [U | undefined, React.Dispatch<React.SetStateAction<U | undefined>>] {
  const raw = urlSearchParams.get(key);
  const [param, setParam] = useState<U | undefined>(
    raw === null ? undefined : convertFromString(raw)
  );
  useEffect(() => {
    if (param === undefined) {
      urlSearchParams.delete(key);
    } else {
      urlSearchParams.set(key, convertToString(param));
    }
  }, [param]);

  return [param, setParam];
}

export default function Statuses() {
  const navigate = useNavigate();

  const urlSearchParams = new URLSearchParams(location.search);
  const [fromTime, setFromTime] = useState(
    urlSearchParams.get("fromTime") || undefined
  );
  const [toTime, setToTime] = useState(
    urlSearchParams.get("toTime") || undefined
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
