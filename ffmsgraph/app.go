package ffmsgraph

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// AadApp -
type AadApp struct {
	ID          string
	AppID       string
	DisplayName string
}

// QueryValueAadApp -
type QueryValueAadApp struct {
	Value []struct {
		ID          string `json:"id"`
		AppID       string `json:"appId"`
		DisplayName string `json:"displayName"`
	} `json:"value"`
}

// Get AAD app -
func (c *Client) getAadApp(appID string) (*AadApp, error) {

	filter := "?$filter=appId%20eq%20"
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/%s%s'%s'", c.HostURL, c.Version, "servicePrincipals", filter, appID), nil)
	req.Header.Add("Authorization", c.Token)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var queryValue QueryValueAadApp
	err = json.Unmarshal(body, &queryValue)
	if err != nil {
		return nil, err
	}

	if len(queryValue.Value) == 1 {
		aadApp := AadApp{
			ID:          queryValue.Value[0].ID,
			AppID:       queryValue.Value[0].AppID,
			DisplayName: queryValue.Value[0].DisplayName,
		}
		return &aadApp, nil
	} else if len(queryValue.Value) == 0 {
		return nil, fmt.Errorf("Can't query this App: %s", appID)
	} else {
		return nil, fmt.Errorf("Mutiple Duplicated AadApp: %s, length: %s", appID, len(queryValue.Value))
	}
}
