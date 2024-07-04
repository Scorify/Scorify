import { useState } from "react";
import { useNavigate } from "react-router-dom";

import { Clear } from "@mui/icons-material";
import {
  Box,
  Button,
  CircularProgress,
  Container,
  IconButton,
  InputAdornment,
  TextField,
  Typography,
} from "@mui/material";

import { CreateInjectModal, EditInject } from "../../components";
import { useInjectsQuery } from "../../graph";

export default function Injects() {
  const { data, loading, error, refetch } = useInjectsQuery();
  const [open, setOpen] = useState(false);

  const navigate = useNavigate();

  const urlSearchParams = new URLSearchParams(location.search);
  const [search, setSearch] = useState(urlSearchParams.get("q") || "");

  const handleRefetch = () => {
    refetch();
  };

  return (
    <Box>
      <CreateInjectModal
        open={open}
        setOpen={setOpen}
        handleRefetch={handleRefetch}
      />
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
            <Button
              variant='contained'
              onClick={() => {
                setOpen(true);
              }}
            >
              Create Inject
            </Button>
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
          {loading && <CircularProgress />}
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
                No Injects Found
              </Typography>
            ) : (
              data.injects.map((inject) => (
                <EditInject
                  key={inject.id}
                  inject={inject}
                  handleRefetch={handleRefetch}
                  visible={inject.title
                    .toLowerCase()
                    .includes(search.toLowerCase())}
                />
              ))
            ))}
        </Box>
      </Container>
    </Box>
  );
}
