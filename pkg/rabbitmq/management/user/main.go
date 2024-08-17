package user

import "net/http"

type UserClient struct {
	host       string
	httpClient *http.Client
}

func Client(host string, httpClient *http.Client) *UserClient {
	return &UserClient{
		host:       host,
		httpClient: httpClient,
	}
}
