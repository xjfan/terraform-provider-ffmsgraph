package ffmsgraph

import (
	"bytes"
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

// AadInvitedUser -
type AadInvitedUser struct {
	ID                      string
	InvitedUserEmailAddress string
	InviteRedirectURL       string
}

// QueryValueAadInvitedUser -
type QueryValueAadInvitedUser struct {
	AadInvitedUserID struct {
		ID string `json:"id"`
	} `json:"invitedUser"`
	InvitedUserEmailAddress string `json:"invitedUserEmailAddress"`
	InviteRedirectURL       string `json:"inviteRedirectUrl"`
}

// Get AAD user -
func (c *Client) getAadUserByMail(mail string) (*AadUser, error) {

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

	if len(queryValue.Value) == 1 {
		aadUser := AadUser{
			ID:          queryValue.Value[0].ID,
			Mail:        queryValue.Value[0].Mail,
			DisplayName: queryValue.Value[0].DisplayName,
		}
		return &aadUser, nil
	} else if len(queryValue.Value) == 0 {
		return nil, fmt.Errorf("Can't query this User: %s", mail)
	} else {
		return nil, fmt.Errorf("Mutiple Duplicated AadUser: %s, length: %s", mail, len(queryValue.Value))
	}
}

// API just post the request without checking the user existed or not
func (c *Client) postAadInvitedUser(mail string, url string) (*AadInvitedUser, error) {
	AadUser, err := c.getAadGroupByMail(mail)
	if AadUser == nil && err != nil {
		requestBody, _ := json.Marshal(map[string]interface{}{
			"invitedUserEmailAddress": mail,
			"inviteRedirectUrl":       url,
			"sendInvitationMessage":   "true",
		})

		req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s/%s", c.HostURL, c.Version, "invitations"), bytes.NewBuffer(requestBody))
		req.Header.Add("Authorization", c.Token)
		if err != nil {
			return nil, err
		}

		body, err := c.doRequest(req)
		if err != nil {
			return nil, err
		}

		var queryValue QueryValueAadInvitedUser
		err = json.Unmarshal(body, &queryValue)
		if err != nil {
			return nil, err
		}

		AadInvitedUser := AadInvitedUser{
			ID:                      queryValue.AadInvitedUserID.ID,
			InvitedUserEmailAddress: queryValue.InvitedUserEmailAddress,
			InviteRedirectURL:       queryValue.InviteRedirectURL,
		}
		return &AadInvitedUser, nil
	} else if AadUser != nil && err == nil {
		AadInvitedUser := AadInvitedUser{
			ID:                      AadUser.ID,
			InvitedUserEmailAddress: mail,
			InviteRedirectURL:       url,
		}
		return &AadInvitedUser, nil
	} else {
		return nil, fmt.Errorf("Error on invitation of User: %s", mail)
	}
}
