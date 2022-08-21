package mythic

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetVPSMap - Returns map of VPS on account
func (c *Client) GetVPSMap(authToken *string) (map[string]VPS, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/vps/servers", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	vpslist := map[string]VPS{}
	err = json.Unmarshal(body, &vpslist)
	if err != nil {
		return nil, err
	}

	return vpslist, nil
}

