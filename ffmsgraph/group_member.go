package ffmsgraph

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// AadGroupMember -
type AadGroupMember struct {
	ID string
}

// QueryValueAadGroupMember -
type QueryValueAadGroupMember struct {
	Value []struct {
		ID string `json:"id"`
	} `json:"value"`
}

// Get AAD group member -
func (c *Client) getAadGroupMember(groupID string, memberID string) (*AadGroupMember, error) {

	filter := "?$filter=id%20eq%20"
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/%s/%s/%s%s'%s'", c.HostURL, c.Beta, "groups", groupID, "members", filter, memberID), nil)
	req.Header.Add("Authorization", c.Token)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var queryValue QueryValueAadGroupMember
	err = json.Unmarshal(body, &queryValue)
	if err != nil {
		return nil, err
	}

	aadGroupMember := AadGroupMember{
		ID: queryValue.Value[0].ID,
	}
	return &aadGroupMember, nil
}

// Create AAD group member -
func (c *Client) createAadGroupMember(groupID string, memberID string) error {

	member := []string{fmt.Sprintf("%s/%s/%s/%s", c.HostURL, c.Version, "directoryObjects", memberID)}

	requestBody, _ := json.Marshal(map[string]interface{}{
		"members@odata.bind": member,
	})

	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/%s/%s/%s", c.HostURL, c.Version, "groups", groupID), bytes.NewBuffer(requestBody))
	req.Header.Add("Authorization", c.Token)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return err
	}

	body, err := c.doRequest(req)

	if err != nil && body == nil {
		return err
	}

	return nil
}

// Delete AAD group -
func (c *Client) deleteAadGroupMember(groupID string, memberID string) error {

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s", c.HostURL, c.Version, "groups", groupID, "members", memberID, "$ref"), nil)
	req.Header.Add("Authorization", c.Token)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return err
	}

	body, err := c.doRequest(req)
	if err != nil && body == nil {
		return err
	}

	return nil
}
