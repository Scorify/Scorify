import { useNavigate } from "react-router-dom";
import { useEffect, useState } from "react";

import { Button, Box, Container, Typography, TextField } from "@mui/material";
import { DateTimePicker, LocalizationProvider } from "@mui/x-date-pickers";
import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";
import { Dayjs } from "dayjs";

import { useURLParam } from "../hooks";
import { StatusEnum, useStatusesQuery } from "../graph";
import { Multiselect } from "../components";

export default function Statuses() {
  const navigate = useNavigate();

  const urlSearchParams = new URLSearchParams(location.search);
  const setSearchParams = (key: string, value: string) => {
    urlSearchParams.set(key, value);
    navigate(`?${urlSearchParams.toString()}`);
  };
  const getUrlParam = (key: string) => urlSearchParams.get(key);
  const deleteUrlParam = (key: string) => urlSearchParams.delete(key);

  const [fromTime, setFromTime] = useURLParam<Dayjs>(
    setSearchParams,
    getUrlParam,
    deleteUrlParam,
    "fromTime",
    (date) => date.toISOString(),
    (s) => new Dayjs(s)
  );
  const [toTime, setToTime] = useURLParam<Dayjs>(
    setSearchParams,
    getUrlParam,
    deleteUrlParam,
    "toTime",
    (date) => date.toISOString(),
    (s) => new Dayjs(s)
  );
  const [fromRound, setFromRound] = useURLParam<number>(
    setSearchParams,
    getUrlParam,
    deleteUrlParam,
    "fromRound",
    (n) => (Number.isNaN(n) ? "" : n.toString()),
    parseInt
  );
  const [toRound, setToRound] = useURLParam<number>(
    setSearchParams,
    getUrlParam,
    deleteUrlParam,
    "toRound",
    String,
    parseInt
  );
  const [minions, setMinions] = useURLParam<string[]>(
    setSearchParams,
    getUrlParam,
    deleteUrlParam,
    "minions",
    JSON.stringify,
    JSON.parse
  );
  const [checks, setChecks] = useURLParam<string[]>(
    setSearchParams,
    getUrlParam,
    deleteUrlParam,
    "checks",
    JSON.stringify,
    JSON.parse
  );
  const [teams, setTeams] = useURLParam<string[]>(
    setSearchParams,
    getUrlParam,
    deleteUrlParam,
    "teams",
    JSON.stringify,
    JSON.parse
  );
  const [statuses, setStatuses] = useURLParam<StatusEnum[]>(
    setSearchParams,
    getUrlParam,
    deleteUrlParam,
    "statuses",
    JSON.stringify,
    (s) => {
      return JSON.parse(s).filter((status: string) => {
        return Object.values(StatusEnum).includes(status as StatusEnum);
      });
    }
  );
  const [limit, setLimit] = useURLParam<number>(
    setSearchParams,
    getUrlParam,
    deleteUrlParam,
    "limit",
    (n) => (Number.isNaN(n) ? "" : n.toString()),
    parseInt
  );
  const [offset, setOffset] = useURLParam<number>(
    setSearchParams,
    getUrlParam,
    deleteUrlParam,
    "offset",
    (n) => (Number.isNaN(n) ? "" : n.toString()),
    parseInt
  );

  useEffect(() => {
    console.log({
      fromTime,
      toTime,
      fromRound,
      toRound,
      minions,
      checks,
      teams,
      statuses,
      limit,
      offset,
    });
  }, [
    fromTime,
    toTime,
    fromRound,
    toRound,
    minions,
    checks,
    teams,
    statuses,
    limit,
    offset,
  ]);

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
            </Box>
          </LocalizationProvider>
          <Box sx={{ display: "flex", flexDirection: "row", gap: 2 }}>
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
              label='Statuses'
              placeholder='Select fields'
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
