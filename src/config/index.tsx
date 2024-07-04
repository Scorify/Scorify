const host = "localhost";
const port = 8080;
const secure = false;

const ws_endpoint = `${secure ? "wss" : "ws"}://${host}:${port}/query`;
const gql_endpoint = `${secure ? "https" : "http"}://${host}:${port}/query`;

export { gql_endpoint, host, port, secure, ws_endpoint };
