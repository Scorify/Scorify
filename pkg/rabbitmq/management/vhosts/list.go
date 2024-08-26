package vhosts

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/scorify/scorify/pkg/rabbitmq/types"
)

func (c *VhostsClient) List() ([]*vhostsResponse, *types.ErrorResponse, error) {
	url := fmt.Sprintf("%s/api/vhosts", c.host)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, nil, err
	}

	// fmt.Println(resp.StatusCode)
	// io.Copy(os.Stdout, resp.Body)

	if resp.StatusCode != http.StatusOK {
		var errorResponse types.ErrorResponse
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
