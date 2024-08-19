package users

import "net/http"

type UsersClient struct {
	host       string
	httpClient *http.Client
}

func Client(host string, httpClient *http.Client) *UsersClient {
	return &UsersClient{
		host:       host,
		httpClient: httpClient,
	}
}
