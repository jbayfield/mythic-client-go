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

// GetVPS - Returns specific VPS
func (c *Client) GetVPS(identifier string, authToken *string) (*VPS, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/vps/servers/%s", c.HostURL, identifier), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	vps := VPS{}
	err = json.Unmarshal(body, &vps)
	if err != nil {
		return nil, err
	}

	return &vps, nil
}

// DestroyVPS - Destroys specific VPS
func (c *Client) DestroyVPS(identifier string, authToken *string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/vps/servers/%s", c.HostURL, identifier), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req, authToken)
	if err != nil {
		return err
	}

	return nil
}
