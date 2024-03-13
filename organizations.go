package dbpu

import (
	"bytes"
	"fmt"
	"net/http"
)

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
// It is used to update an organization with the NewUpdateOrganiationConfig function.
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

// WithBlockedReads is a functional configuration for updating an organization.
// It sets the BlockedReads field of an organization when used with
// NewUpdateOrganiationConfig.
func WithBlockedReads(blockedReads bool) UpdateOrganiationOptions {
	return func(c *Org) { c.BlockedReads = blockedReads }
}

// WithBlockedWrites is a functional configuration for updating an organization.
// It sets the BlockedWrites field of an organization when used with
// NewUpdateOrganiationConfig.
func WithBlockedWrites(blockedWrites bool) UpdateOrganiationOptions {
	return func(c *Org) { c.BlockedWrites = blockedWrites }
}

// WithName is a functional configuration for updating an organization.
// It sets the Name field of an organization when used with
// NewUpdateOrganiationConfig.
func WithName(name string) UpdateOrganiationOptions {
	return func(c *Org) { c.Name = name }
}

// WithOverages is a functional configuration for updating an organization.
// It sets the Overages field of an organization when used with
// NewUpdateOrganiationConfig.
func WithOverages(overages bool) UpdateOrganiationOptions {
	return func(c *Org) { c.Overages = overages }
}

// WithSlug is a functional configuration for updating an organization.
// It sets the Slug field of an organization.
func WithSlug(slug string) UpdateOrganiationOptions {
	return func(c *Org) { c.Slug = slug }
}

// WithType is a functional configuration for updating an organization.
// It sets the Type field of an organization when used with
// NewUpdateOrganiationConfig.
func WithType(orgType string) UpdateOrganiationOptions {
	return func(c *Org) { c.Type = orgType }
}

// ListOrganizations lists the organizations that the user has access to.
func ListOrganizations(apiToken string) ([]Org, error) {
	req, err := newListOrganizationsRequest(apiToken)
	if err != nil {
		return []Org{}, fmt.Errorf("error reading request. %v", err)
	}
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return []Org{}, fmt.Errorf("error sending request. %v", err)
	}
	defer resp.Body.Close()
	response, err := parseResponse[[]Org](resp)
	if err != nil {
		return []Org{}, fmt.Errorf("error decoding body. %v", err)
	}
	return response, nil
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
	req, err := newUpdateOrganizationRequest(organization.Name, config)
	if err != nil {
		return Org{}, fmt.Errorf("error reading request. %v", err)
	}
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return Org{}, fmt.Errorf("error sending request. %v", err)
	}
	response, err := parseResponse[Org](resp)
	if err != nil {
		return Org{}, fmt.Errorf("error decoding body. %v", err)
	}
	defer resp.Body.Close()
	return response, nil
}

// NewUpdateOrganizationRequest returns a new http.Request for updating an organization.
// It is used to update an organization with the UpdateOrganization function.
// It is not exported.
func newUpdateOrganizationRequest(orgName string, config Org) (*http.Request, error) {
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

// newListOrganizationsRequest returns a new http.Request for listing organizations.
func newListOrganizationsRequest(apiToken string) (*http.Request, error) {
	url := fmt.Sprintf("%s/organizations", tursoEndpoint)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	return req, nil
}
