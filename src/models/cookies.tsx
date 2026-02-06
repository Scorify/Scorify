import { CookieSetOptions } from "universal-cookie";

type Cookies = {
  auth?: string;
};

type SetCookie = (
  name: "auth",
  value: string,
  options?: CookieSetOptions | undefined
) => void;

type RemoveCookie = (
  name: "auth",
  options?: CookieSetOptions | undefined
) => void;

type UpdateCookie = () => void;

export type { Cookies, SetCookie, RemoveCookie, UpdateCookie };
