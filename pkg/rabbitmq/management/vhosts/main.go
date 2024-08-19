package vhosts

import "net/http"

type VhostsClient struct {
	host       string
	httpClient *http.Client
}

func Client(host string, httpClient *http.Client) *VhostsClient {
	return &VhostsClient{
		host:       host,
		httpClient: httpClient,
	}
}
