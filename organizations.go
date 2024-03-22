package dbpu

// WithBlockedReads is a functional configuration for updating an organization setting the blockedReads field.
func WithBlockedReads(blockedReads bool) UpdateOrganiationOptions {
	return func(c *Org) { c.BlockedReads = blockedReads }
}

// WithBlockedWrites is a functional configuration for updating an organization	setting the blockedWrites field.
func WithBlockedWrites(blockedWrites bool) UpdateOrganiationOptions {
	return func(c *Org) { c.BlockedWrites = blockedWrites }
}

// WithName is a functional configuration for updating an organization setting the name field.
func WithName(name string) UpdateOrganiationOptions {
	return func(c *Org) { c.Name = name }
}

// WithOverages is a functional configuration for updating an organization setting the overages field.
func WithOverages(overages bool) UpdateOrganiationOptions {
	return func(c *Org) { c.Overages = overages }
}

// WithSlug is a functional configuration for updating an organization setting the slug field.
func WithSlug(slug string) UpdateOrganiationOptions {
	return func(c *Org) { c.Slug = slug }
}

// WithType is a functional configuration for updating an organization setting the type field.
func WithType(orgType string) UpdateOrganiationOptions {
	return func(c *Org) { c.Type = orgType }
}

// ListOrganizations lists the organizations that the user has access to.
func (c *Client) ListOrganizations(apiToken string) ([]Org, error) {
	req, reqErr := c.newListOrgReq(apiToken)
	done, doErr := c.Do(req)
	parsed, parErr := parseResponse[[]Org](done)
	done.Body.Close()
	return resolveApi(parsed, wReqError(reqErr), wDoError(doErr), wParError(parErr))
}

// UpdateOrganiation updates the organization with the given name.
// It is used to update an organization to match the UpdateOrganiationOptions passed as opts.
func (c *Client) UpdateOrganiation(apiToken string, organization Org, opts ...UpdateOrganiationOptions) (Org, error) {
	config := NewUpdateOrganiationConfig(organization, opts...)
	req, reqErr := c.newUpdateOrgReq(organization.Name, config)
	done, doErr := c.Do(req)
	parsed, parErr := parseResponse[Org](done)
	done.Body.Close()
	return resolveApi(parsed, wReqError(reqErr), wDoError(doErr), wParError(parErr))
}

// ListOrgMembers lists the members of the organization with the given name.
func (c *Client) ListOrgMembers(apiToken, orgName string) ([]Member, error) {
	req, reqErr := c.newListOrgMembersReq(apiToken, orgName)
	done, doErr := c.Do(req)
	parsed, parErr := parseResponse[[]Member](done)
	done.Body.Close()
	return resolveApi(parsed, wReqError(reqErr), wDoError(doErr), wParError(parErr))
}

// AddOrgMember adds a member to the organization with the given name.
func (c *Client) AddOrgMember(apiToken, orgName, username, role string) (Member, error) {
	req, reqErr := c.newAddOrgMemberReq(apiToken, orgName, username, role)
	done, doErr := c.Do(req)
	parsed, parErr := parseResponse[Member](done)
	done.Body.Close()
	return resolveApi(parsed, wReqError(reqErr), wDoError(doErr), wParError(parErr))
}

// DeleteOrgMember deletes the member with the given username from the organization with the given name.
func (c *Client) DeleteOrgMember(apiToken, orgName, username string) (DeleteOrgMemberResp, error) {
	req, reqErr := c.newDeleteOrgMemberReq(apiToken, orgName, username)
	done, doErr := c.Do(req)
	parsed, parErr := parseResponse[DeleteOrgMemberResp](done)
	done.Body.Close()
	return resolveApi(parsed, wReqError(reqErr), wDoError(doErr), wParError(parErr))
}

// ListInvites lists the invites of the organization with the given name.
func (c *Client) ListInvites(apiToken, orgName string) ([]Invite, error) {
	req, reqErr := c.newListInvitesReq(apiToken, orgName)
	done, doErr := c.Do(req)
	parsed, parErr := parseResponse[[]Invite](done)
	done.Body.Close()
	return resolveApi(parsed, wReqError(reqErr), wDoError(doErr), wParError(parErr))
}

// CreateInvite creates an invite for the organization with the given name.
func (c *Client) CreateInvite(apiToken, orgName, email, role string) (Invite, error) {
	req, reqErr := c.newCreateInviteReq(apiToken, orgName, email, role)
	done, doErr := c.Do(req)
	parsed, parErr := parseResponse[Invite](done)
	done.Body.Close()
	return resolveApi(parsed, wReqError(reqErr), wDoError(doErr), wParError(parErr))
}
