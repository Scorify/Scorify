package vhosts

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/scorify/scorify/pkg/structs"
)

func (c *VhostsClient) List() ([]*vhostsResponse, *structs.RabbitMQErrorResponse, error) {
	url := fmt.Sprintf("%s/api/vhosts", c.host)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, nil, err
	}

	if resp.StatusCode != http.StatusOK {
		var errorResponse structs.RabbitMQErrorResponse
		err := json.NewDecoder(resp.Body).Decode(&errorResponse)
		if err != nil {
			return nil, nil, err
		}

		return nil, &errorResponse, nil
	}

	var vhosts []*vhostsResponse
	err = json.NewDecoder(resp.Body).Decode(&vhosts)
	if err != nil {
		return nil, nil, err
	}

	return vhosts, nil, nil
}
