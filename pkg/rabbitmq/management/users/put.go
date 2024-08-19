package users

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/scorify/scorify/pkg/rabbitmq/management/types"
	"github.com/scorify/scorify/pkg/static"
)

type UserTag string

const (
	Admin        UserTag = "administrator"
	Monitoring   UserTag = "monitoring"
	Policymaker  UserTag = "policymaker"
	Management   UserTag = "management"
	Impersonator UserTag = "impersonator"
)

type createUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Tags     string `json:"tags"`
}

func (c *UsersClient) Put(user string, password string, tags []UserTag) (*types.ErrorResponse, error) {
	url := fmt.Sprintf("%s/api/users/%s", c.host, user)

	reqBody := createUserRequest{
		Username: user,
		Password: password,
		Tags:     strings.Join(static.MapSlice(tags, func(_ int, tag UserTag) string { return string(tag) }), ","),
	}

	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	reqBodyBuffer := bytes.NewBuffer(reqBodyBytes)

	req, err := http.NewRequest(http.MethodPut, url, reqBodyBuffer)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusNoContent {
		var errResponse types.ErrorResponse
		err := json.NewDecoder(resp.Body).Decode(&errResponse)
		if err != nil {
			return nil, err
		}

		return &errResponse, nil
	}

	return nil, nil
}
