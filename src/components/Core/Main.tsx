import { Dispatch, SetStateAction, useContext, useState } from "react";
import { Outlet, useNavigate } from "react-router-dom";

import { Box, Container, Fade } from "@mui/material";
import { enqueueSnackbar } from "notistack";

import { Drawer, Navbar } from "..";
import { EngineState, useAdminLoginMutation } from "../../graph";
import { AuthContext } from "../Context";

type props = {
  theme: string;
  setTheme: Dispatch<SetStateAction<"dark" | "light">>;
  engineState: EngineState | undefined;
};

export default function Main({ theme, setTheme, engineState }: props) {
  const navigate = useNavigate();
  const [drawerState, setDrawerState] = useState(false);
  const { cookies, setCookie, removeCookie, jwt, me } = useContext(AuthContext);

  const [useAdminLogin] = useAdminLoginMutation({
    onCompleted: (data) => {
      setCookie("auth", data.adminLogin.token, {
        path: data.adminLogin.path,
        expires: new Date(data.adminLogin.expires * 1000),
        httpOnly: data.adminLogin.httpOnly,
        secure: data.adminLogin.secure,
      });

      navigate("/");

      enqueueSnackbar("Reauthenticated successfully", { variant: "success" });
    },
    onError: (error) => {
      enqueueSnackbar(error.message, { variant: "error" });
      console.error(error);
    },
  });

  const returnToAdmin = () => {
    if (jwt) {
      useAdminLogin({
        variables: {
          id: jwt.id,
        },
      });
    }
  };

  return (
    <Box sx={{ minHeight: "100vh", backgroundColor: "default" }}>
      <Drawer
        drawerState={drawerState}
        setDrawerState={setDrawerState}
        me={me}
        jwt={jwt}
        removeCookie={removeCookie}
        returnToAdmin={returnToAdmin}
      />
      <Navbar
        theme={theme}
        setTheme={setTheme}
        setDrawerState={setDrawerState}
        cookies={cookies}
        removeCookie={removeCookie}
        engineState={engineState}
        me={me}
        jwt={jwt}
        returnToAdmin={returnToAdmin}
      />
      <Fade in={true} timeout={1000}>
        <Container component='main'>
          <Outlet />
        </Container>
      </Fade>
    </Box>
  );
}
