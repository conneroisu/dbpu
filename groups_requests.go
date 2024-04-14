package dbpu

import (
	"bytes"
	"fmt"
	"net/http"
)

// newRemoveLocationFromGroupReq creates a request for removing a location from a group.
func (c *Client) newRemoveLocationFromGroupReq(location string) (*http.Request, error) {
	url := fmt.Sprintf("%s/organizations/%s/groups/%s/locations/%s", c.BaseURL, c.OrgName, c.GroupName, location)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.ApiToken))
	return req, nil
}

// newInvalidateGroupTokensRequest creates a request for invalidating all tokens for a group.
func (c *Client) newInvalidateGroupTokensReq() (*http.Request, error) {
	url := fmt.Sprintf("%s/organizations/%s/groups/%s/auth/rotate", c.BaseURL, c.OrgName, c.GroupName)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.ApiToken))
	return req, nil
}

// newCreateGroupReq creates a request for creating a group in an organization.
func (c *Client) newCreateGroupReq(location string) (*http.Request, error) {
	url := fmt.Sprintf("%s/organizations/%s/groups", c.BaseURL, c.OrgName)
	payload := fmt.Sprintf(`{"name": "%s", "location": "%s"}`, c.GroupName, location)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		return nil, fmt.Errorf("error reading request. %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.ApiToken))
	return req, nil
}

// newGroupTokenReq creates a request for creating a token for a group.
func (c *Client) newGroupTokenReq(expiration, authorization string) (*http.Request, error) {
	url := fmt.Sprintf("%s/organizations/%s/groups/%s/auth/tokens?expiration=%s&authorization=%s",
		c.BaseURL, c.OrgName, c.GroupName, expiration, authorization)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request. %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.ApiToken))
	return req, nil
}

// newUpdateGroupReq creates a request for updating a group.
func (c *Client) newUpdateGroupReq() (*http.Request, error) {
	url := fmt.Sprintf(c.BaseURL+"/organizations/%s/groups/%s/update", c.OrgName, c.GroupName)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.ApiToken))
	return req, nil
}

// newGetGroupReq creates a request for getting a group.
func (c *Client) newGetGroupReq(groupName string) (*http.Request, error) {
	url := fmt.Sprintf("%s/organizations/%s/groups/%s", c.BaseURL, c.OrgName, groupName)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.ApiToken))
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

// newAddLocationReq creates a request for adding a location to a group.
func (c *Client) newAddLocationReq(location string) (*http.Request, error) {
	url := fmt.Sprintf("%s/organizations/%s/groups/%s/locations/%s", c.BaseURL, c.OrgName, c.GroupName, location)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.ApiToken))
	return req, nil
}

// newListGroupsReq creates a request for listing groups.
func (c *Client) newListGroupsReq() (*http.Request, error) {
	url := fmt.Sprintf("%s/organizations/%s/groups", c.BaseURL, c.OrgName)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.ApiToken))
	return req, nil
}

// newAddLocationToGroupReq creates a request for adding a location to a group.
func (c *Client) newAddLocationToGroupReq(location string) (*http.Request, error) {
	url := fmt.Sprintf("%s/organizations/%s/groups/%s/locations/%s", c.BaseURL, c.OrgName, c.GroupName, location)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.ApiToken))
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}
