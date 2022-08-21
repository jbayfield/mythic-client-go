package mythic

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
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

// RebootVPS - Reboots specific VPS
func (c *Client) RebootVPS(identifier string, authToken *string) error {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/vps/servers/%s/reboot", c.HostURL, identifier), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req, authToken)
	if err != nil {
		return err
	}

	return nil
}

// powerMgmt - Common power management endpoint
func powerMgmt(identifier string, action string, c *Client) error {
	bodyReader := strings.NewReader(fmt.Sprintf("{ \"power\": \"%s\" }", action))

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/vps/servers/%s/power", c.HostURL, identifier), bodyReader)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return err
	}

	_, err = c.doRequest(req, &c.Token)
	if err != nil {
		return err
	}

	return nil
}

// PowerOnVPS - Power on VPS
func (c *Client) PowerOnVPS(identifier string, authToken *string) error {
	return powerMgmt(identifier, "power-on", c)
}

// PowerOffVPS - Power off VPS
func (c *Client) PowerOffVPS(identifier string, authToken *string) error {
	return powerMgmt(identifier, "power-off", c)
}

// ShutdownVPS - Power on VPS
func (c *Client) ShutdownVPS(identifier string, authToken *string) error {
	return powerMgmt(identifier, "shutdown", c)
}
