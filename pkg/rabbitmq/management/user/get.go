package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/scorify/scorify/pkg/rabbitmq/management/types"
)

type userResponse struct {
	Name             string         `json:"name"`
	PasswordHash     string         `json:"password_hash"`
	HashingAlgorithm string         `json:"hashing_algorithm"`
	Tags             []string       `json:"tags"`
	Limits           map[string]int `json:"limits"`
}

func (c *UserClient) Get(name string) (*userResponse, *types.ErrorResponse, error) {
	url := fmt.Sprintf("%s/api/users/%s", c.host, name)

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
