package ffmsgraph

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// AadUser -
type AadUser struct {
	ID          string
	Mail        string
	DisplayName string
}

// QueryValueAadUser -
type QueryValueAadUser struct {
	Value []struct {
		ID          string `json:"id"`
		Mail        string `json:"mail"`
		DisplayName string `json:"displayName"`
	} `json:"value"`
}

// Get AAD user -
func (c *Client) getAadGroupByMail(mail string) (*AadUser, error) {

	filter := "?$filter=mail%20eq%20"
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/%s%s'%s'", c.HostURL, c.Version, "users", filter, mail), nil)
	req.Header.Add("Authorization", c.Token)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var queryValue QueryValueAadUser
	err = json.Unmarshal(body, &queryValue)
	if err != nil {
		return nil, err
	}

	aadUser := AadUser{
		ID:          queryValue.Value[0].ID,
		Mail:        queryValue.Value[0].Mail,
		DisplayName: queryValue.Value[0].DisplayName,
	}
	return &aadUser, nil
}
