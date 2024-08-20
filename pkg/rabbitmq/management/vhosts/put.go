package vhosts

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/scorify/scorify/pkg/rabbitmq/management/types"
)

func (c *VhostsClient) Put(name string, description string, tags []string, defaultQueueType QueueType) (*types.ErrorResponse, error) {
	escapedVhost := url.PathEscape(name)

	url := fmt.Sprintf("%s/api/vhosts/%s", c.host, escapedVhost)

	reqBody := vhostsRequest{
		DefaultQueueType: defaultQueueType,
		Description:      description,
		Name:             name,
		Tags:             strings.Join(tags, ","),
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
