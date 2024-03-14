package dbpu

// Db is a database.
type Db struct {
	ID            string   `json:"DbId"`
	Hostname      string   `json:"Hostname"`
	Name          string   `json:"Name"`
	Group         string   `json:"group"`
	PrimaryRegion string   `json:"primaryRegion"`
	Regions       []string `json:"regions"`
	Type          string   `json:"type"`
	Version       string   `json:"version"`
}

// Dbs is a list of dbs.
type Dbs struct {
	Databases []Db `json:"databases"`
}

// DbResp is a response to creating a database.
type DbResp struct {
	Database Db `json:"database"`
}

// DbTokenConfig is a configuration for creating a database token.
type DbTokenConfig struct {
	expiration    string // Expiration time for the token (e.g., 2w1d30m).
	authorization string // Authorization level for the token (full-access or read-only).
}

// newDbTokenOpt is a functional option for configuring a CreateDatabaseTokenConfig.
type newDbTokenOpt func(*DbTokenConfig)

// Group is a group in the turso API.
type Group struct {
	Archived      bool     `json:"archived"`
	Locations     []string `json:"locations"`
	Name          string   `json:"name"`
	PrimaryRegion string   `json:"primary"`
	Uuid          string   `json:"uuid"`
}

// GroupResp is a response to adding a location to a group.
type GroupResp struct {
	Group Group `json:"group"`
}

// ListGroupsResp is a response to listing groups.
type ListGroupsResp struct {
	Groups []Group `json:"groups"`
}

// ServerClient is a struct that contains the server and client locations.
type ServerClient struct {
	Server string `json:"server"`
	Client string `json:"client"`
}

// Locations is a list of locations.
type Locations struct {
	Locations map[string]string `json:"locations"`
}

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

// ApiToken is a response to creating a new API token.
type ApiToken struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Token string `json:"token"`
}

// Token is a response to listing API tokens.
type Token struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// ListToksResp is a response to listing API tokens.
type ListToksResp struct {
	Tokens []Token `json:"tokens"`
}

// Jwt is a JSON Web Token.
type Jwt struct {
	Jwt string `json:"jwt"` // jwt is the JSON Web Token.
}

// ValidTokResp is a response to creating a new API token.
type ValidTokResp struct {
	Exp int `json:"exp"`
}

// RevokeTokResp is a response to revoking an API token.
type RevokeTokResp struct {
	Token string `json:"token"`
}
