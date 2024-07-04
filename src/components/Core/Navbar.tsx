import { Dispatch, SetStateAction } from "react";
import { useNavigate } from "react-router";

import { DarkMode, LightMode, Login, Logout, Menu } from "@mui/icons-material";
import {
  AppBar,
  Box,
  Button,
  Toolbar,
  Tooltip,
  Typography,
} from "@mui/material";

import { StatusIndicator } from "..";
import { EngineState, MeQuery } from "../../graph";
import { Cookies, JWT, RemoveCookie } from "../../models";

type props = {
  jwt: JWT;
  returnToAdmin: () => void;
  theme: string;
  setTheme: Dispatch<SetStateAction<"dark" | "light">>;
  setDrawerState: Dispatch<SetStateAction<boolean>>;
  cookies: Cookies;
  removeCookie: RemoveCookie;
  engineState: EngineState | undefined;
  me: MeQuery | undefined;
};

export default function Navbar({
  jwt,
  returnToAdmin,
  theme,
  setTheme,
  setDrawerState,
  removeCookie,
  engineState,
  me,
}: props) {
  const navigate = useNavigate();

  return (
    <Box sx={{ flexGrow: 1 }}>
      <AppBar position='static'>
        <Toolbar>
          <Box sx={{ width: "33%" }}>
            <Tooltip title='Open Drawer'>
              <Button
                onClick={() => {
                  setDrawerState(true);
                }}
              >
                <Menu sx={{ color: "white" }} />
              </Button>
            </Tooltip>
          </Box>
          <Box sx={{ width: "34%", display: "flex", justifyContent: "center" }}>
            {me && (
              <Button
                onClick={() => {
                  navigate("/");
                }}
                sx={{
                  color: "inherit",
                  textTransform: "none",
                }}
              >
                <Typography variant='h6'>{me?.me?.username}</Typography>
              </Button>
            )}
          </Box>
          <Box
            sx={{
              width: "33%",
              display: "flex",
              justifyContent: "flex-end",
              alignItems: "center",
            }}
          >
            <StatusIndicator
              status={engineState}
              positiveTitle='Engine is Scoring'
              negativeTitle='Engine is Paused'
              sx={{ margin: "10px" }}
            />
            {me?.me ? (
              <Tooltip title={jwt?.become ? "Return to Admin" : "Logout"}>
                <Button
                  onClick={() => {
                    if (jwt?.become) {
                      returnToAdmin();
                    } else {
                      removeCookie("auth");
                    }
                    navigate("/login");
                  }}
                  sx={{
                    color: "inherit",
                    minWidth: "0px",
                  }}
                >
                  <Logout />
                </Button>
              </Tooltip>
            ) : (
              <Tooltip title='Login'>
                <Button
                  onClick={() => {
                    navigate("/login");
                  }}
                  sx={{
                    color: "inherit",
                    minWidth: "0px",
                  }}
                >
                  <Login />
                </Button>
              </Tooltip>
            )}
            <Tooltip
              title={theme === "dark" ? "Set Light Mode" : "Set Dark Mode"}
            >
              <Button
                onClick={() => {
                  setTheme(theme === "dark" ? "light" : "dark");
                }}
                sx={{
                  color: "inherit",
                  minWidth: "0px",
                }}
              >
                {theme === "dark" ? (
                  <LightMode sx={{ color: "white" }} />
                ) : (
                  <DarkMode sx={{ color: "white" }} />
                )}
              </Button>
            </Tooltip>
          </Box>
        </Toolbar>
      </AppBar>
    </Box>
  );
}
