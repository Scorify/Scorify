import { useMemo, useState } from "react";

import { Box, Button, TextField, Typography } from "@mui/material";
import { enqueueSnackbar } from "notistack";
import { useNavigate } from "react-router-dom";

import { DeleteUserModal, Dropdown, PasswordInput } from "../..";
import {
  MeQuery,
  UsersQuery,
  useAdminBecomeMutation,
  useAdminLoginMutation,
  useDeleteUserMutation,
  useUpdateUserMutation,
} from "../../../graph";
import { SetCookie } from "../../../models/cookies";

type props = {
  user: UsersQuery["users"][0];
  handleRefetch: () => void;
  visible: boolean;
  setCookie: SetCookie;
  me: MeQuery | undefined;
};

export default function EditCheck({
  me,
  user,
  visible,
  handleRefetch,
  setCookie,
}: props) {
  const navigate = useNavigate();

  const [expanded, setExpanded] = useState(false);

  const [name, setName] = useState<string>(user.username);
  const nameChanged = useMemo(
    () => name !== user.username,
    [name, user.username]
  );

  const [password, setPassword] = useState<string>("");
  const passwordChanged = useMemo(() => password !== "", [password]);

  const [number, setNumber] = useState<number | null | undefined>(user.number);
  const numberChanged = useMemo(() => number !== user.number, [number]);

  const [openDeleteModal, setOpenDeleteModal] = useState(false);

  const [updateUserMutation] = useUpdateUserMutation({
    onCompleted: () => {
      enqueueSnackbar("User updated successfully", { variant: "success" });
      handleRefetch();
      setPassword("");
    },
    onError: (error) => {
      enqueueSnackbar(error.message, { variant: "error" });
      console.error(error);
    },
  });

  const [deleteUserMutation] = useDeleteUserMutation({
    onCompleted: () => {
      enqueueSnackbar("User deleted successfully", { variant: "success" });
      handleRefetch();
    },
    onError: (error) => {
      enqueueSnackbar(error.message, { variant: "error" });
      console.error(error);
    },
  });

  const [adminLoginMutation] = useAdminLoginMutation({
    onCompleted: (data) => {
      setCookie("auth", data.adminLogin.token, {
        path: data.adminLogin.path,
        expires: new Date(data.adminLogin.expires * 1000),
        httpOnly: data.adminLogin.httpOnly,
        secure: data.adminLogin.secure,
      });
      navigate("/");
    },
    onError: (error) => {
      enqueueSnackbar(error.message, { variant: "error" });
      console.error(error);
    },
  });

  const loginAs = () => {
    adminLoginMutation({
      variables: {
        id: user.id,
      },
    });
  };

  const [adminBecomeMutation] = useAdminBecomeMutation({
    onCompleted: (data) => {
      setCookie("auth", data.adminBecome.token, {
        path: data.adminBecome.path,
        expires: new Date(data.adminBecome.expires * 1000),
        httpOnly: data.adminBecome.httpOnly,
        secure: data.adminBecome.secure,
      });
      navigate("/");
    },
    onError: (error) => {
      enqueueSnackbar(error.message, { variant: "error" });
      console.error(error);
    },
  });

  const viewAs = () => {
    adminBecomeMutation({
      variables: {
        id: user.id,
      },
    });
  };

  const handleSave = () => {
    updateUserMutation({
      variables: {
        id: user.id,
        username: nameChanged ? name : undefined,
        password: passwordChanged ? password : undefined,
        number: numberChanged ? number : undefined,
      },
    });
  };

  const handleDelete = () => {
    deleteUserMutation({
      variables: {
        id: user.id,
      },
    });
  };

  return (
    <Dropdown
      modal={
        <DeleteUserModal
          user={user.username}
          open={openDeleteModal}
          setOpen={setOpenDeleteModal}
          handleDelete={handleDelete}
        />
      }
      title={
        <>
          {expanded ? (
            <TextField
              label='Name'
              value={name}
              onClick={(e) => {
                e.stopPropagation();
              }}
              onChange={(e) => {
                setName(e.target.value);
              }}
              sx={{ marginRight: "24px" }}
              size='small'
            />
          ) : (
            <Typography variant='h6' component='div' marginRight='24px'>
              {user.username}
            </Typography>
          )}
          <Typography variant='subtitle1' color='textSecondary' component='div'>
            {user.role}
          </Typography>
        </>
      }
      expandableButtons={
        me?.me?.username !== user.username
          ? [
              <Button variant='contained' onClick={loginAs} color='primary'>
                Become
              </Button>,
              <Button variant='contained' onClick={viewAs} color='secondary'>
                View As
              </Button>,
              <Button
                variant='contained'
                onClick={() => {
                  setOpenDeleteModal(true);
                }}
                color='error'
              >
                Delete
              </Button>,
            ]
          : undefined
      }
      visible={visible}
      expanded={expanded}
      setExpanded={setExpanded}
      toggleButton={
        <Button
          variant='contained'
          color='success'
          onClick={(e) => {
            if (!expanded) {
              e.stopPropagation();
            }

            handleSave();
          }}
        >
          Save
        </Button>
      }
      toggleButtonVisible={nameChanged || passwordChanged || numberChanged}
    >
      <Box
        sx={{
          display: "flex",
          gap: "16px",
          flexWrap: "wrap",
          justifyContent: "center",
        }}
      >
        <PasswordInput
          label='Password'
          value={password}
          onChange={(e) => {
            setPassword(e.target.value);
          }}
        />
        {user.role === "user" && (
          <TextField
            label='Number'
            value={number}
            onChange={(e) => {
              setNumber(parseInt(e.target.value));
            }}
            sx={{ marginRight: "24px" }}
            type='number'
          />
        )}
      </Box>
    </Dropdown>
  );
}
