package permissions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/scorify/scorify/pkg/rabbitmq/management/types"
)

type permissionRequest struct {
	User      string `json:"user"`
	Vhost     string `json:"vhost"`
	Configure string `json:"configure"`
	Write     string `json:"write"`
	Read      string `json:"read"`
}

func (c *PermissionsClient) Put(user string, vhost string, configure string, read string, write string) (*types.ErrorResponse, error) {
	escapedUser := url.QueryEscape(user)
	escapedVhost := url.QueryEscape(vhost)

	url := fmt.Sprintf("%s/api/permissions/%s/%s", c.host, escapedVhost, escapedUser)

	reqBody := permissionRequest{
		User:      user,
		Vhost:     vhost,
		Configure: configure,
		Write:     write,
		Read:      read,
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

	if resp.StatusCode != http.StatusNoContent {
		return nil, fmt.Errorf("failed to get permissions for user %s", user)
	}

	var ErrorResponse types.ErrorResponse
	err = json.NewDecoder(resp.Body).Decode(&ErrorResponse)
	if err != nil {
		return nil, err
	}

	return &ErrorResponse, nil
}