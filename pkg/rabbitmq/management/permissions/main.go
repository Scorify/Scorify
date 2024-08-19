package permissions

import "net/http"

type PermissionsClient struct {
	host       string
	httpClient *http.Client
}

func Client(host string, httpClient *http.Client) *PermissionsClient {
	return &PermissionsClient{
		host:       host,
		httpClient: httpClient,
	}
}
