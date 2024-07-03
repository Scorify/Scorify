import { useState } from "react";

import {
  Box,
  Button,
  Modal,
  Typography,
  FormControl,
  TextField,
  InputLabel,
  Select,
  MenuItem,
  SelectChangeEvent,
} from "@mui/material";
import { enqueueSnackbar } from "notistack";

import { Role, useCreateUserMutation } from "../../../graph";
import { PasswordInput } from "../..";

type props = {
  open: boolean;
  setOpen: (isOpen: boolean) => void;
  handleRefetch: () => void;
};

export default function CreateUserModal({
  open,
  setOpen,
  handleRefetch,
}: props) {
  const [createUserMutation] = useCreateUserMutation({
    onCompleted: () => {
      enqueueSnackbar("User created successfully", { variant: "success" });
      setOpen(false);
      handleRefetch();
    },
    onError: (error) => {
      enqueueSnackbar(error.message, { variant: "error" });
    },
  });

  const [username, setUsername] = useState<string>("");
  const [password, setPassword] = useState<string>("");
  const [role, setRole] = useState<Role>(Role.User);
  const [number, setNumber] = useState<number | undefined>(undefined);

  return (
    <Modal
      open={open}
      onClose={() => {
        setOpen(false);
      }}
    >
      <Box
        sx={{
          position: "absolute",
          top: "25%",
          left: "50%",
          transform: "translate(-50%, -25%)",
          width: "auto",
          maxWidth: "90vw",
          bgcolor: "background.paper",
          border: `1px solid #000`,
          borderRadius: "8px",
          boxShadow: 24,
          p: 4,
        }}
      >
        <Box
          sx={{
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
          }}
        >
          <Typography component='h1' variant='h3'>
            Create New User
          </Typography>

          <FormControl sx={{ gap: "24px", paddingTop: "24px" }}>
            <TextField
              label='Username'
              value={username}
              onChange={(e) => {
                setUsername(e.target.value);
              }}
              required
            />
            <PasswordInput
              label='Password'
              value={password}
              onChange={(e) => {
                setPassword(e.target.value);
              }}
              required
            />

            <FormControl>
              <InputLabel id='role'>Role</InputLabel>
              <Select
                labelId='role'
                value={role}
                label='Role'
                onChange={(event: SelectChangeEvent) => {
                  setRole(event.target.value as Role);
                }}
                required
              >
                {Object.values(Role).map((role) => (
                  <MenuItem key={role} value={role}>
                    {role}
                  </MenuItem>
                ))}
              </Select>
            </FormControl>

            <TextField
              label='Number'
              value={number}
              type='number'
              disabled={role === Role.Admin}
              onChange={(e) => {
                setNumber(parseInt(e.target.value));
              }}
            />
          </FormControl>

          <Button
            variant='contained'
            sx={{ marginTop: "24px" }}
            disabled={
              role === Role.User
                ? !username || !password || !number
                : !username || !password
            }
            onClick={() => {
              if (role === Role.Admin) {
                createUserMutation({
                  variables: {
                    username: username,
                    password: password,
                    role: role,
                  },
                });
              } else {
                createUserMutation({
                  variables: {
                    username: username,
                    password: password,
                    role: role,
                    number: number,
                  },
                });
              }
            }}
          >
            Create User
          </Button>
        </Box>
      </Box>
    </Modal>
  );
}
