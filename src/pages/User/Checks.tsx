import { useState } from "react";
import { useNavigate } from "react-router-dom";

import {
  Container,
  Box,
  TextField,
  Typography,
  InputAdornment,
  IconButton,
} from "@mui/material";
import { Clear } from "@mui/icons-material";

import { useConfigsQuery } from "../../graph";
import { ConfigureCheck } from "../../components";

export default function Checks() {
  const { data, loading, error, refetch } = useConfigsQuery();

  const navigate = useNavigate();

  const urlSearchParams = new URLSearchParams(location.search);
  const [search, setSearch] = useState(urlSearchParams.get("q") || "");

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
          Checks
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
        {loading && (
          <Typography component='h1' variant='h5'>
            Loading...
          </Typography>
        )}
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
        {data &&
          (!data.configs.length ? (
            <Typography component='h1' variant='h5'>
              No checks found
            </Typography>
          ) : (
            data.configs.map((config) => (
              <ConfigureCheck
                key={config.id}
                config={config}
                handleRefetch={refetch}
                visible={
                  config.check.name
                    .toLowerCase()
                    .includes(search.toLowerCase()) ||
                  config.check.source.name
                    .toLowerCase()
                    .includes(search.toLowerCase())
                }
              />
            ))
          ))}
      </Box>
    </Container>
  );
}
