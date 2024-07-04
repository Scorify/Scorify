import { useNavigate } from "react-router-dom";

import { Box, Button, Container, TextField, Typography } from "@mui/material";
import { enqueueSnackbar } from "notistack";

import { useContext, useEffect } from "react";
import { PasswordInput } from "../components";
import { AuthContext } from "../components/Context";
import { useLoginMutation } from "../graph";

export default function Login() {
  const { setCookie, jwt } = useContext(AuthContext);
  const navigate = useNavigate();

  const urlParameters = new URLSearchParams(location.search);

  const [loginMutation] = useLoginMutation({
    onCompleted: (data) => {
      setCookie("auth", data.login.token, {
        path: data.login.path,
        expires: new Date(data.login.expires * 1000),
        httpOnly: data.login.httpOnly,
        secure: data.login.secure,
      });

      enqueueSnackbar("Logged in successfully", { variant: "success" });

      if (!urlParameters.has("next")) {
        navigate("/");
      }
    },
    onError: (error) => {
      enqueueSnackbar(error.message, { variant: "error" });
      console.error(error);
    },
  });

  useEffect(() => {
    if (jwt && urlParameters.has("next")) {
      const next = urlParameters.get("next");
      urlParameters.delete("next");

      navigate(next || "/");
    }
  }, [jwt]);

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    const form_data = new FormData(event.currentTarget);

    let username = form_data.get("username")?.toString();
    let password = form_data.get("password")?.toString();

    if (username && password) {
      loginMutation({
        variables: {
          username: username,
          password: password,
        },
      });
    }
  };

  return (
    <Container component='main' maxWidth='xs'>
      <Box
        sx={{
          marginTop: 8,
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
        }}
      >
        <Typography component='h1' variant='h3'>
          Sign in
        </Typography>
        <Box component='form' onSubmit={handleSubmit}>
          <TextField
            margin='normal'
            required
            fullWidth
            id='username'
            label='Username'
            name='username'
            autoFocus
          />
          <PasswordInput
            margin='normal'
            required
            name='password'
            label='Password'
            id='password'
            fullWidth
          />
          <Button
            type='submit'
            fullWidth
            variant='contained'
            sx={{ mt: 3, mb: 2 }}
          >
            Sign In
          </Button>
        </Box>
      </Box>
    </Container>
  );
}
