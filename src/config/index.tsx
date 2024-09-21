const url = new URL(window.location.href);
const host = import.meta.env.DEV ? "localhost" : url.hostname;
const port = import.meta.env.DEV ? import.meta.env.VITE_PORT : "";
const secure = window.location.protocol === "https:";

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
