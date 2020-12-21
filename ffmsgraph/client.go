package ffmsgraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// graphapiEndpoint -
const graphapiEndpoint string = "https://graph.microsoft.com"
const oauthEndpoint string = "https://login.microsoftonline.com/%s/oauth2/token"

// Client -
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
	Version    string
	Beta       string
}

// AuthResponse -
type AuthResponse struct {
	Token string `json:"access_token"`
}

// APIClient -
func APIClient(tenantID string, clientID string, clientSecret string) (*Client, error) {

	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostURL:    graphapiEndpoint,
		Version:    "v1.0",
		Beta:       "beta",
	}

	rb := url.Values{}
	rb.Set("grant_type", "client_credentials")
	rb.Set("client_id", clientID)
	rb.Set("client_secret", clientSecret)
	rb.Set("resource", graphapiEndpoint)

	req, err := http.NewRequest("POST", fmt.Sprintf(oauthEndpoint, tenantID), strings.NewReader(rb.Encode()))

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	// parse response body
	ar := AuthResponse{}
	err = json.Unmarshal(body, &ar)
	if err != nil {
		return nil, err
	}

	c.Token = ar.Token

	return &c, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {

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
