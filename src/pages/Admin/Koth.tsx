import { useState } from "react";
import { useNavigate } from "react-router-dom";

import { Clear } from "@mui/icons-material";
import {
  Box,
  Button,
  Container,
  IconButton,
  InputAdornment,
  TextField,
  Typography,
} from "@mui/material";

import { CreateKothCheckModal, Loading, EditKothCheck } from "../../components";
import { useKothChecksQuery } from "../../graph";

export default function AdminKoth() {
  const { data, loading, error, refetch } = useKothChecksQuery();

  const [open, setOpen] = useState(false);

  const navigate = useNavigate();

  const urlSearchParams = new URLSearchParams(location.search);
  const [search, setSearch] = useState(urlSearchParams.get("q") || "");

  const handleRefetch = () => {
    refetch();
  };

  return (
    <Box>
      <CreateKothCheckModal
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
            KoTH Checks
          </Typography>
          <Box marginBottom='24px' display='flex' gap='12px'>
            <Button
              variant='contained'
              onClick={() => {
                setOpen(true);
              }}
            >
              Create KoTH Check
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
          {data &&
            (!data.kothChecks.length ? (
              <Typography component='h1' variant='h4'>
                No KoTH Checks Configured
              </Typography>
            ) : (
              <>
                {data.kothChecks.map((check) => (
                  <EditKothCheck
                    key={check.id}
                    check={check}
                    handleRefetch={handleRefetch}
                    visible={
                      check.name.toLowerCase().includes(search.toLowerCase()) ||
                      check.file.toLowerCase().includes(search.toLowerCase())
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
