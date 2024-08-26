package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/scorify/scorify/pkg/rabbitmq/types"
)

func (c *UsersClient) List() ([]*userResponse, error) {
	url := fmt.Sprintf("%s/api/users", c.host)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		var errResponse types.ErrorResponse
		err := json.NewDecoder(resp.Body).Decode(&errResponse)
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("failed to list users: %s", errResponse.Reason)
	}

	var users []*userResponse
	err = json.NewDecoder(resp.Body).Decode(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}
