package mythic

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const HostURL string = "https://api.mythic-beasts.com/beta"
const HostAuthURL string = "https://auth.mythic-beasts.com"

// Client -
type Client struct {
	HostURL    string
	HostAuthURL    string
	HTTPClient *http.Client
	Token      string
	Auth       AuthStruct
}

// AuthStruct -
type AuthStruct struct {
	KeyID string `json:"keyid"`
	Secret string `json:"secret"`
}

// AuthResponse -
type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
}

// NewClient -
func NewClient(host, keyid, secret *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		// Default Hashicups URL
		HostURL: HostURL,
		HostAuthURL: HostAuthURL,
	}

	if host != nil {
		c.HostURL = *host
	}

	// If username or password not provided, return empty client
	if keyid == nil || secret == nil {
		return &c, nil
	}

	c.Auth = AuthStruct{
		KeyID: *keyid,
		Secret: *secret,
	}

	ar, err := c.SignIn()
	if err != nil {
		return nil, err
	}

	c.Token = ar.AccessToken

	return &c, nil
}

func (c *Client) doRequest(req *http.Request, authToken *string) ([]byte, error) {
	token := c.Token

	if authToken != nil {
		token = *authToken
	}

	// If we don't have a token yet don't try to include one
	if token != "" {
		req.Header.Set("Authorization", "Bearer " + token)
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}

