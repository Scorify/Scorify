package permissions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/scorify/scorify/pkg/rabbitmq/types"
)

func (c *PermissionsClient) Put(user string, vhost string, configure string, read string, write string) (*types.ErrorResponse, error) {
	escapedUser := url.PathEscape(user)
	escapedVhost := url.PathEscape(vhost)

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

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusNoContent {
		var errorResponse types.ErrorResponse
		err = json.NewDecoder(resp.Body).Decode(&errorResponse)
		if err != nil {
			return nil, err
		}

		return &errorResponse, nil
	}

	return nil, nil
}
