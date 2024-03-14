package dbpu

import (
	"bytes"
	"fmt"
	"net/http"
)

type DeleteOrgMemberResp struct {
	Member string `json:"member"`
}

// Org is a response to listing organizations.
type Org struct {
	BlockedReads  bool   `json:"blocked_reads"`  // indicates if the organization is blocked from reading data
	BlockedWrites bool   `json:"blocked_writes"` // indicates if the organization is blocked from writing data
	Name          string `json:"name"`           // the name of the organization
	Overages      bool   `json:"overages"`       // indicates if the organization is over its limits
	Slug          string `json:"slug"`           // the slug of the organization
	Type          string `json:"type"`           // the type of the organization
	Token         string `json:"token"`          // the token of the organization
}

// UpdateOrganiationOptions is a functional collective option type for updating an organization.
type UpdateOrganiationOptions func(*Org)

// AuditLogs is a response to listing organizations.
type AuditLogs struct {
	AuditLogs  []AuditLog `json:"audit_logs"` // the audit logs
	Pagination Pagination `json:"pagination"` // the pagination
}

// AuditLog is a response to listing organizations.
type AuditLog struct {
	Author    string `json:"author"`     // the author of the audit log
	Code      string `json:"code"`       // the code of the audit log
	CreatedAt string `json:"created_at"` // the creation date of the audit log
	Data      string `json:"data"`       // the data of the audit log
	Message   string `json:"message"`    // the message of the audit log
	Origin    string `json:"origin"`     // the origin of the audit log
}

// Pagination is a response to listing organizations.
type Pagination struct {
	Page       int `json:"page"`
	PageSize   int `json:"page_size"`
	TotalPages int `json:"total_pages"`
	TotalRows  int `json:"total_rows"`
}

type Invite struct {
	Accepted       bool   `json:"Accepted"`       // indicates if the invite has been Accepted
	CreatedAt      string `json:"CreatedAt"`      // the creation date of the Invite
	DeletedAt      string `json:"DeletedAt"`      // the deletion date of the Invite
	Email          string `json:"Email"`          // the email of the Invite
	ID             int    `json:"ID"`             // the ID of the Invite
	Organization   Org    `json:"Organization"`   // the organization of the Invite
	OrganizationID int    `json:"OrganizationID"` // the ID of the organization of the Invite
	Role           string `json:"Role"`           // the role of the Invite
	Token          string `json:"Token"`          // the token of the Invite
	UpdatedAt      string `json:"UpdatedAt"`      // the update date of the Invite
}

type Invites struct {
	Invites []Invite `json:"invites"` // the invites
}

type Member struct {
	Role     string `json:"role"`     // the role of the member
	Username string `json:"username"` // the username of the member
}

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
func ListOrganizations(apiToken string) ([]Org, error) {
	req, reqErr := newListOrgReq(apiToken)
	done, doErr := (&http.Client{}).Do(req)
	response, parErr := parseResponse[[]Org](done)
	done.Body.Close()
	return resolveApiCall(response, wReqError(reqErr), wDoError(doErr), wParError(parErr))
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

// UpdateOrganiation updates the organization with the given name.
// It is used to update an organization to match the UpdateOrganiationOptions passed as opts.
func UpdateOrganiation(apiToken string, organization Org, opts ...UpdateOrganiationOptions) (Org, error) {
	config := NewUpdateOrganiationConfig(organization, opts...)
	req, reqErr := newUpdateOrgReq(organization.Name, config)
	done, doErr := (&http.Client{}).Do(req)
	parsed, parErr := parseResponse[Org](done)
	done.Body.Close()
	return resolveApiCall(parsed, wReqError(reqErr), wDoError(doErr), wParError(parErr))
}

// ListOrgMembers lists the members of the organization with the given name.
func ListOrgMembers(apiToken, orgName string) ([]Member, error) {
	req, reqErr := newListOrgMembersReq(apiToken, orgName)
	done, doErr := (&http.Client{}).Do(req)
	parsed, parErr := parseResponse[[]Member](done)
	done.Body.Close()
	return resolveApiCall(parsed, wReqError(reqErr), wDoError(doErr), wParError(parErr))
}

// AddOrgMember adds a member to the organization with the given name.
func AddOrgMember(apiToken, orgName, username, role string) (Member, error) {
	req, reqErr := newAddOrgMemberReq(apiToken, orgName, username, role)
	done, doErr := (&http.Client{}).Do(req)
	parsed, parErr := parseResponse[Member](done)
	done.Body.Close()
	return resolveApiCall(parsed, wReqError(reqErr), wDoError(doErr), wParError(parErr))
}

// DeleteOrgMember deletes the member with the given username from the organization with the given name.
func DeleteOrgMember(apiToken, orgName, username string) (DeleteOrgMemberResp, error) {
	req, reqErr := newDeleteOrgMemberReq(apiToken, orgName, username)
	done, doErr := (&http.Client{}).Do(req)
	parsed, parErr := parseResponse[DeleteOrgMemberResp](done)
	done.Body.Close()
	return resolveApiCall(parsed, wReqError(reqErr), wDoError(doErr), wParError(parErr))
}

// ListInvites lists the invites of the organization with the given name.
func ListInvites(apiToken, orgName string) ([]Invite, error) {
	req, reqErr := newListInvitesReq(apiToken, orgName)
	done, doErr := (&http.Client{}).Do(req)
	parsed, parErr := parseResponse[[]Invite](done)
	done.Body.Close()
	return resolveApiCall(parsed, wReqError(reqErr), wDoError(doErr), wParError(parErr))
}

// CreateInvite creates an invite for the organization with the given name.
func CreateInvite(apiToken, orgName, email, role string) (Invite, error) {
	req, reqErr := newCreateInviteReq(apiToken, orgName, email, role)
	done, doErr := (&http.Client{}).Do(req)
	response, parErr := parseResponse[Invite](done)
	done.Body.Close()
	return resolveApiCall(response, wReqError(reqErr), wDoError(doErr), wParError(parErr))
}

// NewUpdateOrganizationRequest returns a new http.Request for updating an organization.
// It is used to update an organization with the UpdateOrganization function.
func newUpdateOrgReq(orgName string, config Org) (*http.Request, error) {
	url := fmt.Sprintf(
		"%s/organizations/%s",
		tursoEndpoint, orgName,
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
func newListOrgReq(apiToken string) (*http.Request, error) {
	url := fmt.Sprintf("%s/organizations", tursoEndpoint)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	return req, nil
}

// newListOrgMembersReq returns a new http.Request for listing organization members.
func newListOrgMembersReq(apiToken, orgName string) (*http.Request, error) {
	url := fmt.Sprintf("%s/organizations/%s/members", tursoEndpoint, orgName)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	return req, nil
}

// AddOrgMember adds a member to the organization with the given name.
func newAddOrgMemberReq(apiToken, orgName, username, role string) (*http.Request, error) {
	url := fmt.Sprintf("%s/organizations/%s/members", tursoEndpoint, orgName)
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
func newDeleteOrgMemberReq(apiToken, orgName, username string) (*http.Request, error) {
	url := fmt.Sprintf("%s/organizations/%s/members/%s", tursoEndpoint, orgName, username)
	req, err := http.NewRequest("DELETE", url, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	req.Header.Set("Content-Type", "application/json")
	return req, err
}

// newListInvitesReq returns a new http.Request for listing organization invites.
func newListInvitesReq(apiToken, orgName string) (*http.Request, error) {
	url := fmt.Sprintf("%s/organizations/%s/invites", tursoEndpoint, orgName)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

// newCreateInviteReq returns a new http.Request for creating an organization invite.
func newCreateInviteReq(apiToken, orgName, email, role string) (*http.Request, error) {
	url := fmt.Sprintf("%s/organizations/%s/invites", tursoEndpoint, orgName)
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
