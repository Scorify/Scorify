import { Dispatch, KeyboardEvent, MouseEvent, SetStateAction } from "react";
import { useNavigate } from "react-router";

import {
  AdminPanelSettings,
  EditNote,
  Flag,
  Group,
  Home,
  KeyboardReturn,
  Login,
  Logout,
  Password,
  QueryBuilder,
  Scoreboard,
  SmartToy,
  Vaccines,
} from "@mui/icons-material";
import {
  Box,
  Divider,
  Drawer,
  List,
  ListItem,
  ListItemButton,
  ListItemIcon,
  ListItemText,
} from "@mui/material";
import { enqueueSnackbar } from "notistack";

import { MeQuery, useLogoutMutation } from "../../graph";
import { JWT } from "../../models";
import { RemoveCookie } from "../../models/cookies";

type drawerItemProps = {
  label: string;
  icon: JSX.Element;
  onClick?: () => void;
};

function DrawerItem({ label, icon, onClick }: drawerItemProps) {
  return (
    <ListItem disablePadding onClick={onClick}>
      <ListItemButton>
        <ListItemIcon>{icon}</ListItemIcon>
        <ListItemText primary={label} />
      </ListItemButton>
    </ListItem>
  );
}

type props = {
  drawerState: boolean;
  setDrawerState: Dispatch<SetStateAction<boolean>>;
  me: MeQuery | undefined;
  jwt: JWT;
  removeCookie: RemoveCookie;
  returnToAdmin: () => void;
};

export default function DrawerComponent({
  drawerState,
  setDrawerState,
  jwt,
  me,
  removeCookie,
  returnToAdmin,
}: props) {
  const navigate = useNavigate();
  const toggleDrawer =
    (open: boolean) => (event: KeyboardEvent | MouseEvent) => {
      if (
        event.type === "keydown" &&
        ((event as KeyboardEvent).key === "Tab" ||
          (event as KeyboardEvent).key === "Shift")
      ) {
        return;
      }

      setDrawerState(open);
    };

  const [logoutMutation] = useLogoutMutation({
    onCompleted: () => {
      enqueueSnackbar("Logged out", { variant: "success" });
      removeCookie("auth");
      navigate("/");
    },
    onError: (error) => {
      enqueueSnackbar(error.message, { variant: "error" });
    },
  });

  return (
    <Drawer anchor={"left"} open={drawerState} onClose={toggleDrawer(false)}>
      <Box
        sx={{ width: 250 }}
        role='presentation'
        onClick={toggleDrawer(false)}
        onKeyDown={toggleDrawer(false)}
      >
        <List>
          <DrawerItem
            label='Home'
            icon={<Home />}
            onClick={() => navigate("/")}
          />
          <DrawerItem
            label='Scoreboard'
            icon={<Scoreboard />}
            onClick={() => navigate("/scoreboard")}
          />
          <DrawerItem
            label='KoTH Scoreboard'
            icon={<Flag />}
            onClick={() => navigate("/koth-scoreboard")}
          />
        </List>
        <Divider />

        {me?.me ? (
          <>
            <List>
              {jwt?.become ? (
                <DrawerItem
                  label='Return to Admin'
                  icon={<KeyboardReturn />}
                  onClick={returnToAdmin}
                />
              ) : (
                <>
                  <DrawerItem
                    label='Logout'
                    icon={<Logout />}
                    onClick={() => {
                      logoutMutation();
                    }}
                  />
                  <DrawerItem
                    label='Change Password'
                    icon={<Password />}
                    onClick={() => navigate("/password")}
                  />
                  <DrawerItem
                    label='Status Query'
                    icon={<QueryBuilder />}
                    onClick={() => navigate("/status")}
                  />
                </>
              )}
            </List>
            <Divider />
            {me.me.role === "user" && (
              <List>
                <DrawerItem
                  label='Checks'
                  icon={<EditNote />}
                  onClick={() => navigate("/checks")}
                />
                <DrawerItem
                  label='Injects'
                  icon={<Vaccines />}
                  onClick={() => navigate("/injects")}
                />
              </List>
            )}
            {me.me.role === "admin" && (
              <List>
                <DrawerItem
                  label='Admin'
                  icon={<AdminPanelSettings />}
                  onClick={() => navigate("/admin")}
                />
                <DrawerItem
                  label='Users'
                  icon={<Group />}
                  onClick={() => navigate("/admin/users")}
                />
                <DrawerItem
                  label='Checks'
                  icon={<EditNote />}
                  onClick={() => navigate("/admin/checks")}
                />
                <DrawerItem
                  label='Injects'
                  icon={<Vaccines />}
                  onClick={() => navigate("/admin/injects")}
                />
                <DrawerItem
                  label='KoTH'
                  icon={<Flag />}
                  onClick={() => navigate("/admin/koth")}
                />
                <DrawerItem
                  label='Minions'
                  icon={<SmartToy />}
                  onClick={() => navigate("/admin/minions")}
                />
              </List>
            )}
          </>
        ) : (
          <List>
            <DrawerItem
              label='Login'
              icon={<Login />}
              onClick={() => navigate("/login")}
            />
          </List>
        )}
      </Box>
    </Drawer>
  );
}
