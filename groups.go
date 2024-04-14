package dbpu

import (
	"fmt"
	"net/http"
)

// AddLocation adds a location to a group.
func (c *Client) AddLocation(location string) (*GroupResp, error) {
	req, err := c.newAddLocationReq(location)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	done, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error doing request: %v", err)
	}
	response, err := parseResponse[GroupResp](done)
	if err != nil {
		return nil, fmt.Errorf("error parsing response: %v", err)
	}
	defer done.Body.Close()
	return &response, nil
}

// ListGroups lists the groups in an organization.
func (c *Client) ListGroups() (*ListGroupsResp, error) {
	req, err := c.newListGroupsReq()
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	done, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error doing request: %v", err)
	}
	response, err := parseResponse[ListGroupsResp](done)
	if err != nil {
		return nil, fmt.Errorf("error parsing response: %v", err)
	}
	defer done.Body.Close()
	return &response, nil
}

// CreateGroup creates a group in an organization.
func (c *Client) CreateGroup(location string) (*GroupResp, error) {
	req, err := c.newCreateGroupReq(location)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	done, err := (&http.Client{}).Do(req)
	if err != nil {
		return nil, fmt.Errorf("error doing request: %v", err)
	}
	response, err := parseResponse[GroupResp](done)
	if err != nil {
		return nil, fmt.Errorf("error parsing response: %v", err)
	}
	defer done.Body.Close()
	return &response, nil
}

// GetGroup gets a group in an organization.
func (c *Client) GetGroup() (*GroupResp, error) {
	req, err := c.newGetGroupReq(c.GroupName)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	done, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error doing request: %v", err)
	}
	parsed, err := parseResponse[GroupResp](done)
	if err != nil {
		return nil, fmt.Errorf("error parsing response: %v", err)
	}
	defer done.Body.Close()
	return &parsed, nil
}

// AddLocationToGroup adds a location to a group.
func (c *Client) AddLocationToGroup(location string) (*GroupResp, error) {
	req, err := c.newAddLocationToGroupReq(location)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	done, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error doing request: %v", err)
	}
	parsed, err := parseResponse[GroupResp](done)
	if err != nil {
		return nil, fmt.Errorf("error parsing response: %v", err)
	}
	defer done.Body.Close()
	return &parsed, nil
}

// CreateGroupToken creates a token for a group.
func (c *Client) CreateGroupToken(expiration, authorization string) (*Jwt, error) {
	req, err := c.newGroupTokenReq(expiration, authorization)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	done, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error doing request: %v", err)
	}
	jwt, err := parseResponse[Jwt](done)
	if err != nil {
		return nil, fmt.Errorf("error parsing response: %v", err)
	}
	defer done.Body.Close()
	return &jwt, nil
}

// RemoveLocationFromGroup removes a location from a group.
func (c *Client) RemoveLocationFromGroup(location string) (*GroupResp, error) {
	req, err := c.newRemoveLocationFromGroupReq(location)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	done, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error doing request: %v", err)
	}
	parsed, err := parseResponse[GroupResp](done)
	if err != nil {
		return nil, fmt.Errorf("error parsing response: %v", err)
	}
	defer done.Body.Close()
	return &parsed, nil
}

// InvalidateGroupTokens invalidates all tokens for a group.
func (c *Client) InvalidateGroupTokens() error {
	req, err := c.newInvalidateGroupTokensReq()
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}
	resp, err := c.Do(req)
	if err != nil {
		return fmt.Errorf("error doing request: %v", err)
	}
	defer resp.Body.Close()
	return nil
}

// UpdateVersionGroup updates the version group.
func (c *Client) UpdateVersionGroup() (*http.Response, error) {
	req, err := c.newUpdateGroupReq()
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	done, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error doing request: %v", err)
	}
	done.Body.Close()
	return done, nil
}
