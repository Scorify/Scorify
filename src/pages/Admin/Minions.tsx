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

import { EditMinion, Loading } from "../../components";
import {
  MinionsQuery,
  useMinionMetricsSubscription,
  useMinionsQuery,
} from "../../graph";

export default function Minions() {
  const { loading, error, refetch } = useMinionsQuery({
    onCompleted: (data) => {
      setMinions(data.minions);
    },
  });

  const navigate = useNavigate();

  const urlSearchParams = new URLSearchParams(location.search);
  const [search, setSearch] = useState(urlSearchParams.get("q") || "");

  const [minions, setMinions] = useState<MinionsQuery["minions"]>([]);

  useMinionMetricsSubscription({
    onData: (data) => {
      let i = minions.findIndex(
        (minion) => minion.id === data.data.data?.minionUpdate?.minion_id
      );

      if (i !== -1 && data.data.data?.minionUpdate) {
        setMinions((prev) => {
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

        {minions ? (
          minions.map((minion) => (
            <EditMinion
              key={minion.id}
              minion={minion}
              handleRefetch={handleRefetch}
              visible={minion.name.toLowerCase().includes(search.toLowerCase())}
            />
          ))
        ) : (
          <Typography component='h1' variant='h4'>
            No Minions
          </Typography>
        )}
      </Box>
    </Container>
  );
}
