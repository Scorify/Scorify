package permissions

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/scorify/scorify/pkg/structs"
)

func (c *PermissionsClient) Delete(user string, vhost string) (*structs.RabbitMQErrorResponse, error) {
	escapedUser := url.PathEscape(user)
	escapedVhost := url.PathEscape(vhost)

	url := fmt.Sprintf("%s/api/permissions/%s/%s", c.host, escapedVhost, escapedUser)

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusNoContent {
		var errorResponse structs.RabbitMQErrorResponse
		err = json.NewDecoder(resp.Body).Decode(&errorResponse)
		if err != nil {
			return nil, err
		}

		return &errorResponse, nil
	}

	return nil, nil
}
