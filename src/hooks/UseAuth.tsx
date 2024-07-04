import { useEffect, useMemo } from "react";

import { ApolloClient, NormalizedCacheObject } from "@apollo/client";
import { jwtDecode } from "jwt-decode";
import { useCookies } from "react-cookie";

import { MeQuery, useMeQuery } from "../graph";
import { Cookies, JWT, RemoveCookie, SetCookie, UpdateCookie } from "../models";

type returns = {
  jwt: JWT | undefined;
  me: MeQuery | undefined;
  meLoading: boolean;
  meError: Error | undefined;
  cookies: Cookies;
  setCookie: SetCookie;
  removeCookie: RemoveCookie;
  updateCookie: UpdateCookie;
};

export function useAuth(
  apolloClient: ApolloClient<NormalizedCacheObject>
): returns {
  const [cookies, setCookie, removeCookie, updateCookie] = useCookies(["auth"]);

  const {
    data: me,
    loading: meLoading,
    error: meError,
    refetch: meRefetch,
  } = useMeQuery({
    onError: (error) => console.error(error),
  });

  useEffect(() => {
    apolloClient.clearStore().then(() => {
      meRefetch();
    });
  }, [cookies?.auth]);

  const jwt = useMemo(
    () => (cookies?.auth ? (jwtDecode(cookies?.auth) as JWT) : undefined),
    [cookies?.auth]
  );

  return {
    jwt,
    me,
    meLoading,
    meError,
    cookies,
    setCookie,
    removeCookie,
    updateCookie,
  };
}
