package rabbitmq

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/scorify/scorify/pkg/config"
)

type managementClient struct {
	host       string
	httpClient *http.Client
}

type managementClientRoundTripper struct {
	transport  http.RoundTripper
	authHeader string
}

func (t *managementClientRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", t.authHeader)
	return t.transport.RoundTrip(req)
}

func ManagementClient() (*managementClient, error) {
	creds := fmt.Sprintf("%s:%s", config.RabbitMQ.User, config.RabbitMQ.Password)
	authHeader := fmt.Sprintf("Basic %s", string(base64.StdEncoding.EncodeToString([]byte(creds))))

	httpClient := &http.Client{
		Transport: &managementClientRoundTripper{
			transport:  http.DefaultTransport,
			authHeader: authHeader,
		},
	}

	return &managementClient{
		host:       fmt.Sprintf("http://%s:%d", config.RabbitMQ.Host, 15672),
		httpClient: httpClient,
	}, nil
}

type rabbitMQError struct {
	Error  string `json:"error"`
	Reason string `json:"reason"`
}

type getUserResponse struct {
	Name             string         `json:"name"`
	PasswordHash     string         `json:"password_hash"`
	HashingAlgorithm string         `json:"hashing_algorithm"`
	Tags             []string       `json:"tags"`
	Limits           map[string]int `json:"limits"`
}

func (c *managementClient) GetUser(name string) (*getUserResponse, *rabbitMQError, error) {
	url := fmt.Sprintf("%s/api/users/%s", c.host, name)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, nil, err
	}

	if resp.StatusCode != http.StatusOK {
		var errResp rabbitMQError
		err := json.NewDecoder(resp.Body).Decode(&errResp)
		if err != nil {
			return nil, nil, err
		}

		return nil, &errResp, nil
	}

	var user getUserResponse
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return nil, nil, err
	}

	return &user, nil, nil
}
