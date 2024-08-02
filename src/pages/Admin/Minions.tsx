import { useState } from "react";
import { useNavigate } from "react-router-dom";

import {
  Box,
  Container,
  Typography,
  TextField,
  InputAdornment,
  IconButton,
} from "@mui/material";
import { Clear } from "@mui/icons-material";

import { useMinionsQuery } from "../../graph";
import { Loading, EditMinion } from "../../components";

export default function Minions() {
  const { data, loading, error, refetch } = useMinionsQuery();

  const navigate = useNavigate();

  const urlSearchParams = new URLSearchParams(location.search);
  const [search, setSearch] = useState(urlSearchParams.get("q") || "");

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

        {data?.minions.map((minion) => (
          <EditMinion
            key={minion.id}
            minion={minion}
            handleRefetch={handleRefetch}
            visible={minion.name.toLowerCase().includes(search.toLowerCase())}
          />
        ))}
      </Box>
    </Container>
  );
}
