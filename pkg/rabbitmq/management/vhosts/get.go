package vhosts

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/scorify/scorify/pkg/rabbitmq/types"
)

func (c *VhostsClient) Get(vhost string) (*vhostsResponse, *types.ErrorResponse, error) {
	escapedVhost := url.PathEscape(vhost)

	url := fmt.Sprintf("%s/api/vhosts/%s", c.host, escapedVhost)

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

	var vhostResp vhostsResponse
	err = json.NewDecoder(resp.Body).Decode(&vhostResp)
	if err != nil {
		return nil, nil, err
	}

	return &vhostResp, nil, nil
}
