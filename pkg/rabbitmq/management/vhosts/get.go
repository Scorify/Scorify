package vhosts

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/scorify/scorify/pkg/rabbitmq/management/types"
)

type vhostsResponse struct {
	ClusterState     map[string]string      `json:"cluster_state"`
	DefaultQueueType string                 `json:"default_queue_type"`
	Description      string                 `json:"description"`
	Metadata         map[string]interface{} `json:"metadata"`
	Name             string                 `json:"name"`
	Tags             []string               `json:"tags"`
	Tracing          bool                   `json:"tracing"`
}

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

	// io.Copy(os.Stdout, resp.Body)

	var vhostResp vhostsResponse
	err = json.NewDecoder(resp.Body).Decode(&vhostResp)
	if err != nil {
		return nil, nil, err
	}

	return &vhostResp, nil, nil
}
