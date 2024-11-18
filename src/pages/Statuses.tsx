import { Box, Button, Container, TextField, Typography } from "@mui/material";
import { DateTimePicker, LocalizationProvider } from "@mui/x-date-pickers";
import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";
import dayjs from "dayjs";

import { Multiselect } from "../components";
import {
  StatusEnum,
  StatusesOptionQuery,
  useStatusesOptionQuery,
  useStatusesQuery,
} from "../graph";
import { useURLParam } from "../hooks";
import { enqueueSnackbar } from "notistack";
import { useEffect } from "react";

export default function Statuses() {
  const { data, refetch } = useStatusesOptionQuery({
    onError: (error) => {
      enqueueSnackbar("Failed to fetch statuses options", {
        variant: "error",
      });
      console.error(error);
    },
  });

  useEffect(() => {
    refetch();
  });

  const { parameter: fromTime, setParameter: setFromTime } =
    useURLParam<dayjs.Dayjs>(
      "fromTime",
      (date) => date.toISOString(),
      (s) => dayjs(s)
    );
  const { parameter: toTime, setParameter: setToTime } =
    useURLParam<dayjs.Dayjs>(
      "toTime",
      (date) => date.toISOString(),
      (s) => dayjs(s)
    );
  const { parameter: fromRound, setParameter: setFromRound } =
    useURLParam<number>(
      "fromRound",
      (n) => (Number.isNaN(n) ? "" : n.toString()),
      parseInt
    );
  const { parameter: toRound, setParameter: setToRound } = useURLParam<number>(
    "toRound",
    (n) => (Number.isNaN(n) ? "" : n.toString()),
    parseInt
  );
  const { parameter: minions, setParameter: setMinions } = useURLParam<
    StatusesOptionQuery["minions"]
  >("minions", JSON.stringify, JSON.parse);
  const { parameter: checks, setParameter: setChecks } = useURLParam<
    StatusesOptionQuery["checks"]
  >("checks", JSON.stringify, JSON.parse);
  const { parameter: teams, setParameter: setTeams } = useURLParam<
    StatusesOptionQuery["teams"]
  >("teams", JSON.stringify, JSON.parse);
  const { parameter: statuses, setParameter: setStatuses } = useURLParam<
    StatusEnum[]
  >("statuses", JSON.stringify, (s) => {
    return JSON.parse(s).filter((status: string) => {
      return Object.values(StatusEnum).includes(status as StatusEnum);
    });
  });
  const { parameter: limit, setParameter: setLimit } = useURLParam<number>(
    "limit",
    (n) => (Number.isNaN(n) ? "" : n.toString()),
    parseInt
  );
  const { parameter: offset, setParameter: setOffset } = useURLParam<number>(
    "offset",
    (n) => (Number.isNaN(n) ? "" : n.toString()),
    parseInt
  );

  const {} = useStatusesQuery({
    variables: {
      statusesInputQuery: {
        from_time: fromTime,
        to_time: toTime,
        from_round: fromRound,
        to_round: toRound,
        minions: minions?.map((minion) => minion.id),
        checks: checks?.map((check) => check.id),
        users: teams?.map((team) => team.id),
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
      <Box sx={{ mt: 3 }}>
        <Box sx={{ display: "flex", flexDirection: "column", gap: 2 }}>
          <LocalizationProvider dateAdapter={AdapterDayjs}>
            <Box sx={{ display: "flex", flexDirection: "row", gap: 2 }}>
              <DateTimePicker
                sx={{ flex: 1 }}
                label='Start Time'
                value={fromTime}
                onChange={(date) => {
                  setFromTime(date || undefined);
                }}
              />
              <DateTimePicker
                sx={{ flex: 1 }}
                label='End Time'
                value={toTime}
                onChange={(date) => {
                  setToTime(date || undefined);
                }}
              />
            </Box>
          </LocalizationProvider>
          <Box sx={{ display: "flex", flexDirection: "row", gap: 2 }}>
            <TextField
              label='From Round'
              type='number'
              value={fromRound || ""}
              onChange={(e) => setFromRound(parseInt(e.target.value))}
              sx={{ flex: 1 }}
            />
            <TextField
              label='To Round'
              type='number'
              value={toRound || ""}
              onChange={(e) => setToRound(parseInt(e.target.value))}
              sx={{ flex: 1 }}
            />
            <TextField
              label='Limit'
              type='number'
              value={limit || 10}
              onChange={(e) => setLimit(parseInt(e.target.value))}
              sx={{ flex: 1 }}
            />
            <TextField
              label='Offset'
              type='number'
              value={offset || 0}
              onChange={(e) => setOffset(parseInt(e.target.value))}
              sx={{ flex: 1 }}
            />
          </Box>
          <Box sx={{ display: "flex", flexDirection: "row", gap: 2 }}>
            <Multiselect
              label='Teams'
              placeholder='Select Teams'
              options={data?.teams.map((team) => team.username) || []}
              selected={teams?.map((team) => team.username) || []}
              setSelected={(teams) =>
                setTeams(
                  data?.teams.filter((team) => teams.includes(team.username)) ||
                    []
                )
              }
              sx={{ flex: 1 }}
            />
            <Multiselect
              label='Checks'
              placeholder='Select Checks'
              options={data?.checks.map((check) => check.name) || []}
              selected={checks?.map((check) => check.name) || []}
              setSelected={(checks) =>
                setChecks(
                  data?.checks.filter((check) => checks.includes(check.name)) ||
                    []
                )
              }
              sx={{ flex: 1 }}
            />
          </Box>
          <Box sx={{ display: "flex", flexDirection: "row", gap: 2 }}>
            <Multiselect
              label='Minions'
              placeholder='Select Minions'
              options={data?.minions.map((minion) => minion.id) || []}
              selected={minions?.map((minion) => minion.id) || []}
              setSelected={(minions) =>
                setMinions(
                  data?.minions.filter((minion) =>
                    minions.includes(minion.id)
                  ) || []
                )
              }
              sx={{ flex: 1 }}
            />
            <Multiselect
              label='Statuses'
              placeholder='Select Statuses'
              options={[StatusEnum.Up, StatusEnum.Down, StatusEnum.Unknown]}
              selected={statuses || []}
              setSelected={(statuses) =>
                setStatuses(
                  statuses.filter((status: string) => {
                    return Object.values(StatusEnum).includes(
                      status as StatusEnum
                    );
                  }) as StatusEnum[]
                )
              }
              sx={{ flex: 1 }}
            />
          </Box>
          <Button variant='contained'>
            <Typography>Search</Typography>
          </Button>
        </Box>
      </Box>
    </Container>
  );
}
