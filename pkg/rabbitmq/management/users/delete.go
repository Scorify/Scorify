package users

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/scorify/scorify/pkg/rabbitmq/management/types"
)

func (c *UsersClient) Delete(name string) (*types.ErrorResponse, error) {
	escapedUser := url.QueryEscape(name)

	url := fmt.Sprintf("%s/api/users/%s", c.host, escapedUser)

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusNoContent {
		var errResponse types.ErrorResponse
		err := json.NewDecoder(resp.Body).Decode(&errResponse)
		if err != nil {
			return nil, err
		}

		return &errResponse, nil
	}

	return nil, nil
}
