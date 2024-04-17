package dbpu

import "fmt"

// ListOrganizations lists the organizations that the user has access to.
func (c *Client) ListOrganizations(apiToken string) (*[]Org, error) {
	req, err := c.newListOrgReq(apiToken)
	if err != nil {
		return nil, err
	}
	done, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error doing request: %v", err)
	}
	parsed, err := parseResponse[[]Org](done)
	if err != nil {
		return nil, err
	}
	done.Body.Close()
	return parsed, nil
}

// UpdateOrganiation updates the organization with the given name.
// It is used to update an organization to match the UpdateOrganiationOptions passed as opts.
func (c *Client) UpdateOrganiation(apiToken string, organization Org, opts ...UpdateOrganiationOptions) (*Org, error) {
	config := NewUpdateOrganiationConfig(organization, opts...)
	req, err := c.newUpdateOrgReq(organization.Name, config)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	done, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error doing request: %v", err)
	}
	parsed, err := parseResponse[Org](done)
	if err != nil {
		return nil, fmt.Errorf("error parsing response: %v", err)
	}
	done.Body.Close()
	return parsed, nil
}

// ListOrgMembers lists the members of the organization with the given name.
func (c *Client) ListOrgMembers(apiToken, orgName string) (*[]Member, error) {
	req, err := c.newListOrgMembersReq(apiToken, orgName)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	done, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error doing request: %v", err)
	}
	parsed, err := parseResponse[[]Member](done)
	if err != nil {
		return nil, fmt.Errorf("error parsing response: %v", err)
	}
	done.Body.Close()
	return parsed, nil
}

// AddOrgMember adds a member to the organization with the given name.
func (c *Client) AddOrgMember(apiToken, orgName, username, role string) (*Member, error) {
	req, err := c.newAddOrgMemberReq(apiToken, orgName, username, role)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	done, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error doing request: %v", err)
	}
	parsed, err := parseResponse[Member](done)
	if err != nil {
		return nil, fmt.Errorf("error parsing response: %v", err)
	}
	done.Body.Close()
	return parsed, nil
}

// DeleteOrgMember deletes the member with the given username from the organization with the given name.
func (c *Client) DeleteOrgMember(apiToken, orgName, username string) (*DeleteOrgMemberResp, error) {
	req, err := c.newDeleteOrgMemberReq(apiToken, orgName, username)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	done, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error doing request: %v", err)
	}
	parsed, err := parseResponse[DeleteOrgMemberResp](done)
	if err != nil {
		return nil, fmt.Errorf("error parsing response: %v", err)
	}
	done.Body.Close()
	return parsed, nil
}

// ListInvites lists the invites of the organization with the given name.
func (c *Client) ListInvites(apiToken, orgName string) (*[]Invite, error) {
	req, err := c.newListInvitesReq(apiToken, orgName)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	done, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error doing request: %v", err)
	}
	parsed, err := parseResponse[[]Invite](done)
	if err != nil {
		return nil, fmt.Errorf("error parsing response: %v", err)
	}
	done.Body.Close()
	return parsed, nil
}

// CreateInvite creates an invite for the organization with the given name.
func (c *Client) CreateInvite(apiToken, orgName, email, role string) (*Invite, error) {
	req, err := c.newCreateInviteReq(apiToken, orgName, email, role)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	done, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error doing request: %v", err)
	}
	parsed, err := parseResponse[Invite](done)
	if err != nil {
		return nil, fmt.Errorf("error parsing response: %v", err)
	}
	done.Body.Close()
	return parsed, nil
}
