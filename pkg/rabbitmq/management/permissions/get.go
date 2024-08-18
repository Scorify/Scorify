package permissions

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type permissionsResponse []struct {
	User      string `json:"user"`
	Vhost     string `json:"vhost"`
	Configure string `json:"configure"`
	Write     string `json:"write"`
	Read      string `json:"read"`
}

func (c *PermissionsClient) Get(user string) (*permissionsResponse, error) {
	url := fmt.Sprintf("%s/api/users/%s/permissions", c.host, user)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get permissions for user %s", user)
	}

	var permissions permissionsResponse
	err = json.NewDecoder(resp.Body).Decode(&permissions)
	if err != nil {
		return nil, err
	}

	return &permissions, nil
}
