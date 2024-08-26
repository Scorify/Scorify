package users

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/scorify/scorify/pkg/rabbitmq/types"
)

func (c *UsersClient) Get(name string) (*userResponse, *types.ErrorResponse, error) {
	escapedUser := url.PathEscape(name)

	url := fmt.Sprintf("%s/api/users/%s", c.host, escapedUser)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, nil, err
	}

	if resp.StatusCode != http.StatusOK {
		var errResponse types.ErrorResponse
		err := json.NewDecoder(resp.Body).Decode(&errResponse)
		if err != nil {
			return nil, nil, err
		}

		return nil, &errResponse, nil
	}

	var user userResponse
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return nil, nil, err
	}

	return &user, nil, nil
}
