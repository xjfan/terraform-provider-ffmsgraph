package ffmsgraph

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// AadGroup -
type AadGroup struct {
	ID          string
	Description string
	DisplayName string
}

// QueryValue -
type QueryValue struct {
	Value []struct {
		ID          string `json:"id"`
		Description string `json:"description"`
		DisplayName string `json:"displayName"`
	} `json:"value"`
}

// Get AAD group -
func (c *Client) getAadGroup(display_name string) (*AadGroup, error) {

	filter := "?$filter=displayName%20eq%20"
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/%s%s'%s'", c.HostURL, c.Version, "groups", filter, display_name), nil)
	req.Header.Add("Authorization", c.Token)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var queryValue QueryValue
	err = json.Unmarshal(body, &queryValue)
	if err != nil {
		return nil, err
	}

	aadGroup := AadGroup{
		ID:          queryValue.Value[0].ID,
		Description: queryValue.Value[0].Description,
		DisplayName: queryValue.Value[0].DisplayName,
	}
	return &aadGroup, nil
}
