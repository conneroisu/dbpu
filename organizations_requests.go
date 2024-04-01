package dbpu

import (
	"bytes"
	"fmt"
	"net/http"
)

// NewUpdateOrganizationRequest returns a new http.Request for updating an organization.
// It is used to update an organization with the UpdateOrganization function.
func (c *Client) newUpdateOrgReq(orgName string, config Org) (*http.Request, error) {
	url := fmt.Sprintf(
		"%s/organizations/%s",
		c.BaseURL, orgName,
	)
	reqJsonBody := fmt.Sprintf(
		`{
			"blocked_reads": %t, 
			"blocked_writes": %t, 
			"name": "%s", 
			"overages": %t, 
			"slug": "%s", 
			"type": "%s"
		}`,
		config.BlockedReads,
		config.BlockedWrites,
		config.Name,
		config.Overages,
		config.Slug,
		config.Type,
	)
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer([]byte(reqJsonBody)))
	req.Header.Set("Content-Type", "application/json")
	return req, err
}

// newListOrgReq returns a new http.Request for listing organizations.
func (c *Client) newListOrgReq(apiToken string) (*http.Request, error) {
	url := fmt.Sprintf("%s/organizations", c.BaseURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	return req, nil
}

// newListOrgMembersReq returns a new http.Request for listing organization members.
func (c *Client) newListOrgMembersReq(apiToken, orgName string) (*http.Request, error) {
	url := fmt.Sprintf("%s/organizations/%s/members", c.BaseURL, orgName)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	return req, nil
}

// AddOrgMember adds a member to the organization with the given name.
func (c *Client) newAddOrgMemberReq(apiToken, orgName, username, role string) (*http.Request, error) {
	url := fmt.Sprintf("%s/organizations/%s/members", c.BaseURL, orgName)
	reqJsonBody := fmt.Sprintf(
		`{
			"role": "%s", 
			"username": "%s"
		}`,
		role,
		username,
	)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(reqJsonBody)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	return req, err

}

// newDeleteOrgMemberReq returns a new http.Request for deleting an organization member.
func (c *Client) newDeleteOrgMemberReq(apiToken, orgName, username string) (*http.Request, error) {
	url := fmt.Sprintf("%s/organizations/%s/members/%s", c.BaseURL, orgName, username)
	req, err := http.NewRequest("DELETE", url, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	req.Header.Set("Content-Type", "application/json")
	return req, err
}

// newListInvitesReq returns a new http.Request for listing organization invites.
func (c *Client) newListInvitesReq(apiToken, orgName string) (*http.Request, error) {
	url := fmt.Sprintf("%s/organizations/%s/invites", c.BaseURL, orgName)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

// newCreateInviteReq returns a new http.Request for creating an organization invite.
func (c *Client) newCreateInviteReq(apiToken, orgName, email, role string) (*http.Request, error) {
	url := fmt.Sprintf("%s/organizations/%s/invites", c.BaseURL, orgName)
	reqJsonBody := fmt.Sprintf(
		`{
			"email": "%s", 
			"role": "%s"
		}`,
		email,
		role,
	)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(reqJsonBody)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	return req, err
}

// NewUpdateOrganiationConfig returns a new UpdateOrganiationConfig.
func NewUpdateOrganiationConfig(organization Org, opts ...UpdateOrganiationOptions) Org {
	config := Org{
		BlockedReads:  organization.BlockedReads,
		BlockedWrites: organization.BlockedWrites,
		Name:          organization.Name,
		Overages:      organization.Overages,
		Slug:          organization.Slug,
		Type:          organization.Type,
	}
	for _, opt := range opts {
		opt(&config)
	}
	return config
}
