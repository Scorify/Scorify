package management

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/scorify/scorify/pkg/config"
	"github.com/scorify/scorify/pkg/rabbitmq/management/permissions"
	"github.com/scorify/scorify/pkg/rabbitmq/management/user"
)

type client struct {
	User        *user.UserClient
	Permissions *permissions.PermissionsClient
}

type roundTripper struct {
	transport  http.RoundTripper
	authHeader string
}

func (t *roundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", t.authHeader)
	return t.transport.RoundTrip(req)
}

func Client() (*client, error) {
	creds := fmt.Sprintf("%s:%s", config.RabbitMQ.User, config.RabbitMQ.Password)
	authHeader := fmt.Sprintf("Basic %s", string(base64.StdEncoding.EncodeToString([]byte(creds))))

	httpClient := &http.Client{
		Transport: &roundTripper{
			transport:  http.DefaultTransport,
			authHeader: authHeader,
		},
	}

	host := fmt.Sprintf("http://%s:%d", config.RabbitMQ.Host, 15672)

	return &client{
		User:        user.Client(host, httpClient),
		Permissions: permissions.Client(host, httpClient),
	}, nil
}
