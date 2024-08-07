import { useMemo, useState } from "react";
import { useNavigate } from "react-router-dom";

import { Clear } from "@mui/icons-material";
import {
  Box,
  Container,
  IconButton,
  InputAdornment,
  TextField,
  Typography,
} from "@mui/material";
import { enqueueSnackbar } from "notistack";

import { Loading, MinionGroup } from "../../components";
import {
  MinionsQuery,
  useMinionMetricsSubscription,
  useMinionsQuery,
} from "../../graph";

export default function Minions() {
  const [minions, setMinions] = useState<MinionsQuery["minions"]>([]);
  const [update, setUpdate] = useState(Date.now());

  const { loading, error, refetch } = useMinionsQuery({
    onCompleted(data) {
      setMinions(data.minions);
    },
    onError: (error) => {
      console.error(error);
      enqueueSnackbar("Failed to fetch minions", { variant: "error" });
    },
  });

  const activeMinions = useMemo(
    () =>
      minions.filter(
        (minion) =>
          new Date(minion.metrics?.timestamp).getTime() >
            Date.now() - 1000 * 60 && minion.deactivated === false
      ),
    [minions, update]
  );

  const staleMinions = useMemo(
    () =>
      minions.filter(
        (minion) =>
          new Date(minion.metrics?.timestamp).getTime() <=
            Date.now() - 1000 * 60 && minion.deactivated === false
      ),
    [minions, update]
  );

  const deactivatedMinions = useMemo(
    () => minions.filter((minion) => minion.deactivated === true),
    [minions, update]
  );

  const sortMinions = () => {
    setUpdate(Date.now());
  };

  const navigate = useNavigate();

  const urlSearchParams = new URLSearchParams(location.search);
  const [search, setSearch] = useState(urlSearchParams.get("q") || "");

  useMinionMetricsSubscription({
    onData: (data) => {
      if (!data.data.data?.minionUpdate) {
        return;
      }

      let i = minions.findIndex(
        (minion) => minion.id === data.data.data?.minionUpdate?.minion_id
      );

      if (i === -1) {
        refetch();
      }

      setMinions((prev) => {
        if (!prev) {
          return prev;
        }

        let copy = [...prev];
        copy[i] = {
          ...copy[i],
          metrics: data.data.data?.minionUpdate,
        };
        return copy;
      });
    },
  });

  const handleRefetch = () => {
    refetch();
  };

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
          Minions
        </Typography>
        <Box marginBottom='24px' display='flex' gap='12px'>
          <TextField
            label='Search'
            size='small'
            onChange={(e) => {
              setSearch(e.target.value);

              urlSearchParams.set("q", e.target.value);

              navigate(`?${urlSearchParams.toString()}`);
            }}
            value={search}
            InputProps={{
              endAdornment: (
                <InputAdornment position='end'>
                  <IconButton
                    size='small'
                    onClick={() => {
                      setSearch("");
                      navigate("");
                    }}
                  >
                    <Clear />
                  </IconButton>
                </InputAdornment>
              ),
            }}
          />
        </Box>

        {loading && <Loading />}

        {error && (
          <>
            <Typography component='h1' variant='h4'>
              Encountered Error
            </Typography>
            <Typography component='h1' variant='body1'>
              {error.message}
            </Typography>
          </>
        )}
        {!error && !loading && (
          <>
            <MinionGroup
              title='Active Minions'
              minions={activeMinions}
              handleRefetch={handleRefetch}
              sortMinions={sortMinions}
              search={search}
              emptyString='No Active Minions'
              elevation={1}
            />

            <MinionGroup
              title='Stale Minions'
              minions={staleMinions}
              handleRefetch={handleRefetch}
              sortMinions={() => {}}
              search={search}
              emptyString='No Stale Minions'
              elevation={1}
            />

            <MinionGroup
              title='Deactivated Minions'
              minions={deactivatedMinions}
              handleRefetch={handleRefetch}
              sortMinions={() => {}}
              search={search}
              emptyString='No Deactivated Minions'
              elevation={1}
            />
          </>
        )}
      </Box>
    </Container>
  );
}
