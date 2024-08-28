package vhosts

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/scorify/scorify/pkg/rabbitmq/types"
)

func (c *VhostsClient) Delete(vhost string) (*types.ErrorResponse, error) {
	escapedVhost := url.PathEscape(vhost)

	url := fmt.Sprintf("%s/api/vhosts/%s", c.host, escapedVhost)

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
