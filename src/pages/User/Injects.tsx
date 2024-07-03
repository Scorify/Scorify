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

import { useInjectsQuery } from "../../graph";
import { Inject } from "../../components";

export default function Injects() {
  const { data, loading, error, refetch } = useInjectsQuery();

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
          Injects
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
          (!data.injects.length ? (
            <Typography component='h1' variant='h5'>
              No injects found
            </Typography>
          ) : (
            data.injects.map((inject) => (
              <Inject
                key={inject.id}
                inject={inject}
                visible={inject.title
                  .toLowerCase()
                  .includes(search.toLowerCase())}
                handleRefetch={refetch}
              />
            ))
          ))}
      </Box>
    </Container>
  );
}
