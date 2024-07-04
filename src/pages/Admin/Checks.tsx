import { useState } from "react";
import { useNavigate } from "react-router-dom";

import {
  Container,
  Box,
  Button,
  Typography,
  TextField,
  InputAdornment,
  IconButton,
} from "@mui/material";
import { Clear } from "@mui/icons-material";

import { useChecksQuery } from "../../graph";
import CreateCheckModal from "../../components/Admin/Checks/CreateCheckModal";
import { EditCheck } from "../../components";

export default function Checks() {
  const { data, loading, error, refetch } = useChecksQuery();
  const [open, setOpen] = useState(false);

  const navigate = useNavigate();

  const urlSearchParams = new URLSearchParams(location.search);
  const [search, setSearch] = useState(urlSearchParams.get("q") || "");

  const handleRefetch = () => {
    refetch();
  };

  return (
    <Box>
      <CreateCheckModal
        open={open}
        setOpen={setOpen}
        data={data}
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
            Checks
          </Typography>
          <Box marginBottom='24px' display='flex' gap='12px'>
            <Button
              variant='contained'
              onClick={() => {
                setOpen(true);
              }}
            >
              Create Check
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
            (!data.checks.length ? (
              <Typography component='h1' variant='h4'>
                No Checks Configured
              </Typography>
            ) : (
              <>
                {data.checks.map((check) => (
                  <EditCheck
                    key={check.name}
                    check={check}
                    handleRefetch={handleRefetch}
                    visible={
                      check.name.toLowerCase().includes(search.toLowerCase()) ||
                      check.source.name
                        .toLowerCase()
                        .includes(search.toLowerCase())
                    }
                  />
                ))}
              </>
            ))}
        </Box>
      </Container>
    </Box>
  );
}
