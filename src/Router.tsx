import { ReactElement, ReactNode, Suspense } from "react";
import { RouterProvider, createBrowserRouter } from "react-router-dom";

import { ApolloClient, NormalizedCacheObject } from "@apollo/client";
import { CircularProgress } from "@mui/material";
import { enqueueSnackbar } from "notistack";

import { Admin, Error, Main, User } from "./components";
import { AuthContext } from "./components/Context/AuthProvider";
import {
  useEngineStateSubscription,
  useGlobalNotificationSubscription,
} from "./graph";
import { useAuth } from "./hooks";
import {
  AdminChecks,
  AdminPanel,
  ChangePassword,
  Home,
  Injects,
  Login,
  Me,
  Scoreboard,
  ScoreboardRound,
  UserChecks,
  UserInjects,
  Users,
} from "./pages";

const LazyComponent = ({ element }: { element: ReactNode }): ReactElement => {
  return <Suspense fallback={<CircularProgress />}>{element}</Suspense>;
};

type props = {
  theme: "dark" | "light";
  setTheme: React.Dispatch<React.SetStateAction<"dark" | "light">>;
  apolloClient: ApolloClient<NormalizedCacheObject>;
};

export function Router({ theme, setTheme, apolloClient }: props) {
  const {
    jwt,
    me,
    meLoading,
    meError,
    cookies,
    setCookie,
    removeCookie,
    updateCookie,
  } = useAuth(apolloClient);

  useGlobalNotificationSubscription({
    onData: (data) => {
      if (data.data.data?.globalNotification) {
        enqueueSnackbar(data.data.data.globalNotification.message, {
          variant: data.data.data.globalNotification.type,
        });
      }
    },
    onError: (error) => {
      console.error(error);
    },
  });

  const { data: engineState } = useEngineStateSubscription({
    onError: (error) => {
      console.error(error);
    },
  });

  const router = createBrowserRouter([
    {
      path: "/",
      element: (
        <LazyComponent
          element={
            <Main
              theme={theme}
              setTheme={setTheme}
              engineState={engineState?.engineState}
            />
          }
        />
      ),
      children: [
        {
          index: true,
          element: <LazyComponent element={<Home />} />,
        },
        {
          path: "login",
          element: <LazyComponent element={<Login />} />,
        },
        {
          path: "me",
          element: <LazyComponent element={<Me />} />,
        },
        {
          path: "password",
          element: <LazyComponent element={<ChangePassword />} />,
        },
        {
          path: "scoreboard",
          children: [
            {
              index: true,
              element: <LazyComponent element={<Scoreboard theme={theme} />} />,
            },
            {
              path: ":round",
              element: (
                <LazyComponent element={<ScoreboardRound theme={theme} />} />
              ),
            },
          ],
        },
        {
          path: "admin",
          element: <LazyComponent element={<Admin />} />,
          children: [
            {
              index: true,
              element: (
                <LazyComponent
                  element={
                    <AdminPanel engineState={engineState?.engineState} />
                  }
                />
              ),
            },
            {
              path: "checks",
              element: <LazyComponent element={<AdminChecks />} />,
            },
            {
              path: "users",
              element: <LazyComponent element={<Users />} />,
            },
            {
              path: "injects",
              element: <LazyComponent element={<Injects />} />,
            },
          ],
        },
        {
          element: <LazyComponent element={<User />} />,
          children: [
            {
              path: "checks",
              element: <LazyComponent element={<UserChecks />} />,
            },
            {
              path: "injects",
              element: <LazyComponent element={<UserInjects />} />,
            },
          ],
        },
        {
          path: "*",
          element: (
            <LazyComponent
              element={<Error code={404} message={"Page Not Found"} />}
            />
          ),
        },
      ],
    },
  ]);

  return (
    <AuthContext.Provider
      value={{
        jwt,
        cookies,
        setCookie,
        removeCookie,
        updateCookie,
        me,
        meLoading,
        meError,
      }}
    >
      <RouterProvider router={router} />
    </AuthContext.Provider>
  );
}
