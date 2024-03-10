package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

// Org is a response to listing organizations.
type Org struct {
	BlockedReads  bool   `json:"blocked_reads"`
	BlockedWrites bool   `json:"blocked_writes"`
	Name          string `json:"name"`
	Overages      bool   `json:"overages"`
	Slug          string `json:"slug"`
	Type          string `json:"type"`
	Token         string `json:"token"`
}

// UpdateOrganiationOptions is a functional collective option type for updating an organization.
// It is used to update an organization with the NewUpdateOrganiationConfig function.
type UpdateOrganiationOptions func(*Org)

// AuditLogs is a response to listing organizations.
type AuditLogs struct {
	AuditLogs  []AuditLog `json:"audit_logs"`
	Pagination Pagination `json:"pagination"`
}

// AuditLog is a response to listing organizations.
type AuditLog struct {
	Author    string `json:"author"`
	Code      string `json:"code"`
	CreatedAt string `json:"created_at"`
	Data      string `json:"data"`
	Message   string `json:"message"`
	Origin    string `json:"origin"`
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
// It sets the Slug field of an organization when used with
// NewUpdateOrganiationConfig.
func WithSlug(slug string) UpdateOrganiationOptions {
	return func(c *Org) { c.Slug = slug }
}

// WithType is a functional configuration for updating an organization.
// It sets the Type field of an organization when used with
// NewUpdateOrganiationConfig.
func WithType(orgType string) UpdateOrganiationOptions {
	return func(c *Org) { c.Type = orgType }
}

// createListOrganizationsRequest returns a new http.Request for listing organizations.
func createListOrganizationsRequest(apiToken string) (*http.Request, error) {
	url := fmt.Sprintf("%s/organizations", TursoEndpoint)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	return req, nil
}

// ListOrganizations lists the organizations that the user has access to.
func ListOrganizations(apiToken string) ([]Org, error) {
	req, err := createListOrganizationsRequest(apiToken)
	if err != nil {
		return []Org{}, fmt.Errorf("Error reading request. %v", err)
	}
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return []Org{}, fmt.Errorf("Error sending request. %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []Org{}, fmt.Errorf("Error reading body. %v", err)
	}
	response, err := parseStruct[[]Org](body)
	if err != nil {
		return []Org{}, fmt.Errorf("Error decoding body. %v", err)
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

// NewUpdateOrganizationRequest returns a new http.Request for updating an organization.
// It is used to update an organization with the UpdateOrganization function.
// It is not exported.
func createUpdateOrganizationRequest(orgName string, config Org) (*http.Request, error) {
	url := fmt.Sprintf(
		"%s/organizations/%s",
		TursoEndpoint, orgName,
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

// UpdateOrganiation updates the organization with the given name.
// It is used to update an organization to match the UpdateOrganiationOptions passed as opts.
func UpdateOrganiation(apiToken string, organization Org, opts ...UpdateOrganiationOptions) (Org, error) {
	config := NewUpdateOrganiationConfig(organization, opts...)
	req, err := createUpdateOrganizationRequest(organization.Name, config)
	if err != nil {
		return Org{}, fmt.Errorf("Error reading request. %v", err)
	}
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return Org{}, fmt.Errorf("Error sending request. %v", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Org{}, fmt.Errorf("Error reading body. %v", err)
	}
	response, err := parseStruct[Org](body)
	if err != nil {
		return Org{}, fmt.Errorf("Error decoding body. %v", err)
	}
	defer resp.Body.Close()
	return response, nil
}
