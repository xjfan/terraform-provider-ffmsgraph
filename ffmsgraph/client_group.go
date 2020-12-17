package ffmsgraph

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// QueryValue -
type QueryValue struct {
	Odata string `json:"@odata.context"`
	Value []interface{}
}

// AadGroup -
type AadGroup struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	DisplayName string `json:"displayName"`
}

// Get AAD group -
func (c *Client) getAadGroup(display_name string) ([]QueryValue, error) {

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

	queryValue := []QueryValue{}
	err = json.Unmarshal(body, &queryValue)
	if err != nil {
		return nil, err
	}

	log.Printf(queryValue.AadGroup)
	return queryValue, nil
}
