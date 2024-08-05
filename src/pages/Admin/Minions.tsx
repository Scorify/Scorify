import { useState } from "react";
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

import { EditMinion, Loading } from "../../components";
import {
  MinionsQuery,
  useMinionMetricsSubscription,
  useMinionsQuery,
} from "../../graph";

export default function Minions() {
  const { loading, error, refetch } = useMinionsQuery({
    onCompleted: (data) => {
      setActiveMinions(
        data.minions.filter(
          (minion) =>
            new Date(minion.metrics?.timestamp).getTime() >
            Date.now() - 1000 * 60
        )
      );
      setStaleMinions(
        data.minions.filter(
          (minion) =>
            new Date(minion.metrics?.timestamp).getTime() <=
            Date.now() - 1000 * 60
        )
      );
    },
    onError: (error) => {
      console.error(error);
      enqueueSnackbar("Failed to fetch minions", { variant: "error" });
    },
  });

  const navigate = useNavigate();

  const urlSearchParams = new URLSearchParams(location.search);
  const [search, setSearch] = useState(urlSearchParams.get("q") || "");

  const [activeMinions, setActiveMinions] = useState<MinionsQuery["minions"]>(
    []
  );
  const [staleMinions, setStaleMinions] = useState<MinionsQuery["minions"]>([]);
  const [deactivatedMinions, setDeactivatedMinions] = useState<
    MinionsQuery["minions"]
  >([]);

  useMinionMetricsSubscription({
    onData: (data) => {
      if (!data.data.data?.minionUpdate) {
        return;
      }

      let i = activeMinions.findIndex(
        (minion) => minion.id === data.data.data?.minionUpdate?.minion_id
      );

      if (i === -1) {
        i = staleMinions.findIndex(
          (minion) => minion.id === data.data.data?.minionUpdate?.minion_id
        );

        if (i === -1) {
          setActiveMinions((prev) => [
            ...prev,
            {
              ...staleMinions[i],
              metrics: data.data.data?.minionUpdate,
            },
          ]);

          setStaleMinions((prev) =>
            prev.filter(
              (minion) => minion.id !== data.data.data?.minionUpdate?.minion_id
            )
          );
        }
      } else {
        setActiveMinions((prev) => {
          const newMinions = [...prev];
          newMinions[i] = {
            ...newMinions[i],
            metrics: data.data.data?.minionUpdate,
          };
          return newMinions;
        });
      }
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
        <Box width='100%'>
          <Typography component='h1' variant='h4' sx={{ mb: "8px" }}>
            Active Minions
          </Typography>
          {activeMinions && activeMinions.length ? (
            activeMinions.map((minion) => (
              <EditMinion
                key={minion.id}
                minion={minion}
                handleRefetch={handleRefetch}
                visible={minion.name
                  .toLowerCase()
                  .includes(search.toLowerCase())}
              />
            ))
          ) : (
            <Typography component='h1' variant='h6' sx={{ mb: "16px" }}>
              No Active Minions
            </Typography>
          )}
        </Box>

        <Box width='100%'>
          <Typography component='h1' variant='h4' sx={{ mb: "8px" }}>
            Stale Minions
          </Typography>
        </Box>
        {staleMinions && staleMinions.length ? (
          staleMinions.map((minion) => (
            <EditMinion
              key={minion.id}
              minion={minion}
              handleRefetch={handleRefetch}
              visible={minion.name.toLowerCase().includes(search.toLowerCase())}
            />
          ))
        ) : (
          <Typography component='h1' variant='h6' sx={{ mb: "16px" }}>
            No Stale Minions
          </Typography>
        )}

        <Box width='100%'>
          <Typography component='h1' variant='h4' sx={{ mb: "8px" }}>
            Deactivated Minions
          </Typography>
        </Box>
        {deactivatedMinions && deactivatedMinions.length ? (
          deactivatedMinions.map((minion) => (
            <EditMinion
              key={minion.id}
              minion={minion}
              handleRefetch={handleRefetch}
              visible={minion.name.toLowerCase().includes(search.toLowerCase())}
            />
          ))
        ) : (
          <Typography component='h1' variant='h6'>
            No Deactivated Minions
          </Typography>
        )}
      </Box>
    </Container>
  );
}
