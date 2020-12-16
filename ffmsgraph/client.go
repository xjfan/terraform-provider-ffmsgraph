package ffmsgraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

// graphapiEndpoint -
const graphapiEndpoint string = "https://graph.microsoft.com"

// Client -
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
}

// AuthStruct -
type AuthStruct struct {
	grantType    string `json:"grant_type"`
	clientID     string `json:"client_id"`
	clientSecret string `json:"client_secret"`
	resource     string `json:"resource"`
}

// AuthResponse -
type AuthResponse struct {
	Token string `json:"access_token"`
}

// APIClient -
func APIClient(tenantID string, clientID string, clientSecret string) (*Client, diag.Diagnostics) {
	var diags diag.Diagnostics

	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		// Default Hashicups URL
		HostURL: fmt.Sprintf("https://login.microsoftonline.com/%s/oauth2/token", tenantID),
	}

	rb, err := json.Marshal(AuthStruct{
		grantType:    "client_credentials",
		clientID:     clientID,
		clientSecret: clientSecret,
		resource:     graphapiEndpoint,
	})

	req, err := http.NewRequest("POST", c.HostURL, strings.NewReader(string(rb)))
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create oauth request",
			Detail:   fmt.Sprintf(string(rb)),
		})
		return nil, diags
	}

	body, err := c.doRequest(req)

	// parse response body
	ar := AuthResponse{}
	err = json.Unmarshal(body, &ar)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to convert to json",
			Detail:   fmt.Sprintf(string(body)),
		})
		return nil, diags
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
