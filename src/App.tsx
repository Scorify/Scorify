import { useMemo, useState } from "react";

import {
  ApolloClient,
  ApolloProvider,
  InMemoryCache,
  split,
} from "@apollo/client";
import { GraphQLWsLink } from "@apollo/client/link/subscriptions";
import { getMainDefinition } from "@apollo/client/utilities";
import { ThemeProvider } from "@emotion/react";
import { CssBaseline } from "@mui/material";
import { createTheme } from "@mui/material/styles";
import { createUploadLink } from "apollo-upload-client";
import { createClient } from "graphql-ws";
import { SnackbarProvider } from "notistack";

import { Router } from "./Router";
import { gql_endpoint, ws_endpoint } from "./config";

const lightTheme = createTheme({
  palette: {
    mode: "light",
  },
});

const darkTheme = createTheme({
  palette: {
    mode: "dark",
  },
});

export default function App() {
  let savedTheme = localStorage.getItem("theme");
  if (!savedTheme) {
    savedTheme = window.matchMedia("(prefers-color-scheme: dark)").matches
      ? "dark"
      : "light";
  }

  const [theme, setTheme] = useState<"dark" | "light">(
    savedTheme as "dark" | "light"
  );
  const muiTheme = useMemo(() => {
    localStorage.setItem("theme", theme);

    return theme === "dark" ? darkTheme : lightTheme;
  }, [theme]);

  const wsLink = new GraphQLWsLink(
    createClient({
      url: ws_endpoint,
    })
  );

  const httpLink = createUploadLink({
    uri: gql_endpoint,
    credentials: "include",
  });

  const splitLink = split(
    ({ query }) => {
      const definition = getMainDefinition(query);
      return (
        definition.kind === "OperationDefinition" &&
        definition.operation === "subscription"
      );
    },
    wsLink,
    httpLink
  );

  const client = new ApolloClient({
    link: splitLink,
    cache: new InMemoryCache(),
  });

  return (
    <ThemeProvider theme={muiTheme}>
      <ApolloProvider client={client}>
        <CssBaseline />
        <SnackbarProvider maxSnack={3}>
          <Router theme={theme} setTheme={setTheme} apolloClient={client} />
        </SnackbarProvider>
      </ApolloProvider>
    </ThemeProvider>
  );
}
