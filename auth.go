package mythic

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

// SignIn - Get a new token for user
func (c *Client) SignIn() (*AuthResponse, error) {
	if c.Auth.KeyID == "" || c.Auth.Secret == "" {
		return nil, fmt.Errorf("define key ID and secret")
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/login", c.HostAuthURL), strings.NewReader("grant_type=client_credentials"))
	req.Header.Add("Authorization", "Basic "+basicAuth(c.Auth.KeyID, c.Auth.Secret))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	ar := AuthResponse{}
	err = json.Unmarshal(body, &ar)
	if err != nil {
		return nil, err
	}

	return &ar, nil
}

// SignIn - Get a new token for user
func (c *Client) GetUserTokenSignIn(auth AuthStruct) (*AuthResponse, error) {
	if auth.KeyID == "" || auth.Secret == "" {
		return nil, fmt.Errorf("define key ID and secret")
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/login", c.HostAuthURL), strings.NewReader("grant_type=client_credentials"))
	req.Header.Add("Authorization", "Basic "+basicAuth(auth.KeyID, auth.Secret))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, errors.New("unable to login")
	}

	ar := AuthResponse{}
	err = json.Unmarshal(body, &ar)
	if err != nil {
		return nil, err
	}

	return &ar, nil
}
