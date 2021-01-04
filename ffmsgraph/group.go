package ffmsgraph

import (
	"bytes"
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

// QueryValueAadGroup -
type QueryValueAadGroup struct {
	Value []struct {
		ID          string `json:"id"`
		Description string `json:"description"`
		DisplayName string `json:"displayName"`
	} `json:"value"`
}

// Get AAD group -
func (c *Client) getAadGroupByName(displayName string) (*AadGroup, error) {

	filter := "?$filter=displayName%20eq%20"
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/%s%s'%s'", c.HostURL, c.Version, "groups", filter, displayName), nil)
	req.Header.Add("Authorization", c.Token)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var queryValue QueryValueAadGroup
	err = json.Unmarshal(body, &queryValue)
	if err != nil {
		return nil, err
	}

	if len(queryValue.Value) == 1 {
		aadGroup := AadGroup{
			ID:          queryValue.Value[0].ID,
			Description: queryValue.Value[0].Description,
			DisplayName: queryValue.Value[0].DisplayName,
		}
		return &aadGroup, nil
	} else if len(queryValue.Value) == 0 {
		return nil, nil
	} else {
		return nil, fmt.Errorf("Mutiple Duplicated AadGroup: %s, length: %s", displayName, len(queryValue.Value))
	}
}

// Get AAD group -
func (c *Client) getAadGroup(ID string) (*AadGroup, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/%s/%s", c.HostURL, c.Version, "groups", ID), nil)
	req.Header.Add("Authorization", c.Token)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var aadGroup AadGroup
	err = json.Unmarshal(body, &aadGroup)
	if err != nil {
		return nil, err
	}

	return &aadGroup, nil
}

// Create AAD group -
func (c *Client) createAadGroup(displayName string) (*AadGroup, error) {

	owners := []string{fmt.Sprintf("%s/%s/%s/%s", c.HostURL, c.Version, "directoryObjects", "1c41b7f8-cbd3-4a31-84fd-8c57028ea49e"), fmt.Sprintf("%s/%s/%s/%s", c.HostURL, c.Version, "directoryObjects", c.ObjectID)}

	requestBody, _ := json.Marshal(map[string]interface{}{
		"displayName":       displayName,
		"mailNickname":      displayName,
		"description":       displayName,
		"mailEnabled":       "false",
		"securityEnabled":   "true",
		"owners@odata.bind": owners,
	})

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s/%s", c.HostURL, c.Version, "groups"), bytes.NewBuffer(requestBody))
	req.Header.Add("Authorization", c.Token)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var aadGroup AadGroup
	err = json.Unmarshal(body, &aadGroup)
	if err != nil {
		return nil, err
	}

	return &aadGroup, nil
}

// Delete AAD group -
func (c *Client) deleteAadGroup(ID string) error {

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/%s/%s/%s", c.HostURL, c.Version, "groups", ID), nil)
	req.Header.Add("Authorization", c.Token)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return err
	}

	body, err := c.doRequest(req)
	if err != nil && body != nil {
		return err
	}

	return nil
}
