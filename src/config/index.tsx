const host = import.meta.env.DEV ? "localhost" : import.meta.env.VITE_DOMAIN;
const port = import.meta.env.DEV ? import.meta.env.VITE_PORT : "";
const secure = import.meta.env.DEV
  ? false
  : import.meta.env.VITE_DOMAIN?.startsWith("http://")
  ? false
  : true;

const http_endpoint = `${secure ? "https" : "http"}://${host}${
  port !== "" ? `:${port}` : ""
}`;
const ws_endpoint = `${secure ? "wss" : "ws"}://${host}${
  port !== "" ? `:${port}` : ""
}/api/query`;
const gql_endpoint = `${secure ? "https" : "http"}://${host}${
  port !== "" ? `:${port}` : ""
}/api/query`;

export { http_endpoint, gql_endpoint, host, port, secure, ws_endpoint };
