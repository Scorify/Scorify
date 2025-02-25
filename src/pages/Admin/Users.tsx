import { useContext, useState } from "react";
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

import { CreateUserModal, EditUser } from "../../components";
import { AuthContext } from "../../components/Context";
import { useUsersQuery } from "../../graph";

export default function Checks() {
  const { me, setCookie } = useContext(AuthContext);
  const { data, loading, error, refetch } = useUsersQuery();
  const [open, setOpen] = useState(false);

  const navigate = useNavigate();

  const urlSearchParams = new URLSearchParams(location.search);
  const [search, setSearch] = useState(urlSearchParams.get("q") || "");

  const handleRefetch = () => {
    refetch();
  };

  return (
    <Box>
      <CreateUserModal
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
            Users
          </Typography>
          <Box marginBottom='24px' display='flex' gap='12px'>
            <Button
              variant='contained'
              onClick={() => {
                setOpen(true);
              }}
            >
              Create User
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
            (!data.users.length ? (
              <Typography component='h1' variant='h4'>
                No Users Exist
              </Typography>
            ) : (
              <>
                {data.users.map((user) => (
                  <EditUser
                    key={user.id}
                    user={user}
                    handleRefetch={handleRefetch}
                    visible={
                      user.username
                        .toLowerCase()
                        .includes(search.toLowerCase()) ||
                      user.role.toLowerCase().includes(search.toLowerCase()) ||
                      user.id.toString().includes(search)
                    }
                    setCookie={setCookie}
                    me={me}
                  />
                ))}
              </>
            ))}
        </Box>
      </Container>
    </Box>
  );
}
