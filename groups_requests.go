package dbpu

import (
	"bytes"
	"fmt"
	"net/http"
)

// newRemoveLocationFromGroupReq creates a request for removing a location from a group.
func (c *Client) newRemoveLocationFromGroupReq(orgName, apiToken, groupName, location string) (*http.Request, error) {
	url := fmt.Sprintf("/organizations/%s/groups/%s/locations/%s", c.BaseURL, orgName, groupName, location)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	return req, nil
}

// newInvalidateGroupTokensRequest creates a request for invalidating all tokens for a group.
func (c *Client) newInvalidateGroupTokensReq(orgName, apiToken, groupName string) (*http.Request, error) {
	url := fmt.Sprintf("%s/organizations/%s/groups/%s/auth/rotate", c.BaseURL, orgName, groupName)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	return req, nil
}

// newCreateGroupReq creates a request for creating a group in an organization.
func (c *Client) newCreateGroupReq(orgName, apiToken, groupName, location string) (*http.Request, error) {
	url := fmt.Sprintf("%s/organizations/%s/groups", c.BaseURL, orgName)
	payload := fmt.Sprintf(`{"name": "%s", "location": "%s"}`, groupName, location)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		return nil, fmt.Errorf("error reading request. %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	return req, nil
}

// newGroupTokenReq creates a request for creating a token for a group.
func (c *Client) newGroupTokenReq(orgName, apiToken, groupName, expiration, authorization string) (*http.Request, error) {
	url := fmt.Sprintf("%s/organizations/%s/groups/%s/auth/tokens?expiration=%s&authorization=%s", c.BaseURL, orgName, groupName, expiration, authorization)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request. %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	return req, nil
}

// newUpdateGroupReq creates a request for updating a group.
func (c *Client) newUpdateGroupReq(orgName, apiToken, groupName string) (*http.Request, error) {
	url := fmt.Sprintf(c.BaseURL+"/organizations/%s/groups/%s/update", orgName, groupName)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	return req, nil
}

// newGetGroupReq creates a request for getting a group.
func (c *Client) newGetGroupReq(orgName, apiToken, groupName string) (*http.Request, error) {
	url := fmt.Sprintf("%s/organizations/%s/groups/%s", c.BaseURL, orgName, groupName)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

// newAddLocationReq creates a request for adding a location to a group.
func (c *Client) newAddLocationReq(orgName, apiToken, groupName, location string) (*http.Request, error) {
	url := fmt.Sprintf("/organizations/%s/groups/%s/locations/%s", c.BaseURL, orgName, groupName, location)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	return req, nil
}

// newListGroupsReq creates a request for listing groups.
func (c *Client) newListGroupsReq(orgName, apiToken string) (*http.Request, error) {
	url := fmt.Sprintf("/organizations/%s/groups", c.BaseURL, orgName)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	return req, nil
}

// newDeleteGroupReq creates a request for deleting a group.
func (c *Client) newDeleteGroupReq(orgName, apiToken, groupName string) (*http.Request, error) {
	url := fmt.Sprintf("%s/organizations/%s/groups/%s", c.BaseURL, orgName, groupName)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	return req, nil
}

// newTransferGroupReq creates a request for transferring a group to a new organization.
func (c *Client) newTransferGroupReq(orgName, apiToken, groupName, newOrgName string) (*http.Request, error) {
	url := fmt.Sprintf("%s/organizations/%s/groups/%s/transfer", c.BaseURL, orgName, groupName)
	payload := fmt.Sprintf(`{"organization": "%s"}`, newOrgName)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		return nil, fmt.Errorf("error reading request. %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	return req, nil
}

// newAddLocationToGroupReq creates a request for adding a location to a group.
func (c *Client) newAddLocationToGroupReq(orgName, apiToken, groupName, location string) (*http.Request, error) {
	url := fmt.Sprintf("%s/organizations/%s/groups/%s/locations/%s", c.BaseURL, orgName, groupName, location)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}
