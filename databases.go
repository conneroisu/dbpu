package dbpu

import (
	"bytes"
	"fmt"
	"net/http"
)

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

// Databases is a list of dbs.
type Databases struct {
	Databases []Db `json:"databases"`
}

// DatabaseResponse is a response to creating a database.
type DatabaseResponse struct {
	Database Db `json:"database"`
}

// createDatabaseTokenConfig is a configuration for creating a database token.
type createDatabaseTokenConfig struct {
	expiration    string // Expiration time for the token (e.g., 2w1d30m).
	authorization string // Authorization level for the token (full-access or read-only).
}

// CreateDatabaseTokenOption is a functional option for configuring a CreateDatabaseTokenConfig.
type CreateDatabaseTokenOption func(*createDatabaseTokenConfig)

// createCreateDatabaseRequest returns a new http.Request for creating a database.
func createCreateDatabaseRequest(orgToken string, orgName string, name string, group string) (*http.Request, error) {
	url := fmt.Sprintf(
		"%s/organizations/%s/databases",
		tursoEndpoint, orgName,
	)
	reqJsonBody := fmt.Sprintf(
		`{
			"name": "%s", 
			"group": "%s"
		}`,
		name,
		group,
	)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(reqJsonBody)))
	if err != nil {
		return nil, fmt.Errorf("error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", orgToken))
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

// NewCreateDatabaseTokenConfig returns a new CreateDatabaseTokenConfig.
func newCreateDatabaseTokenConfig(opts ...CreateDatabaseTokenOption) *createDatabaseTokenConfig {
	c := &createDatabaseTokenConfig{
		expiration:    "",
		authorization: "",
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// WithExpiration sets the expiration time for the token (e.g., 2w1d30m).
func WithExpiration(expiration string) CreateDatabaseTokenOption {
	return func(c *createDatabaseTokenConfig) { c.expiration = expiration }
}

// WithAuthorization sets the authorization level for the token (full-access or read-only).
func WithAuthorization(authorization string) CreateDatabaseTokenOption {
	return func(c *createDatabaseTokenConfig) { c.authorization = authorization }
}

// CreateDatabase creates a database with the given name and group.
func CreateDatabase(orgToken string, orgName string, name string, group string) (Db, error) {
	req, err := createCreateDatabaseRequest(orgToken, orgName, name, group)
	if err != nil {
		return Db{}, fmt.Errorf("error reading request. %v", err)
	}
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return Db{}, fmt.Errorf("error sending request. %v", err)
	}
	response, err := parseResponse[DatabaseResponse](resp)
	if err != nil {
		return Db{}, fmt.Errorf("error reading response. %v", err)
	}
	defer resp.Body.Close()
	return response.Database, nil
}

// createCreateDatabaseTokenRequest creates a request for creating a token for a database owned by an organization with an optional given expiration and authorization.
func createCreateDatabaseTokenRequest(orgName string, dbName string, apiTok string, opts ...CreateDatabaseTokenOption) (*http.Request, error) {
	config := newCreateDatabaseTokenConfig(opts...)
	url := fmt.Sprintf(
		"%s/organizations/%s/databases/%s/auth/tokens?expiration=%s&authorization=%s",
		tursoEndpoint, orgName, dbName, config.expiration, config.authorization,
	)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request for database token: %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiTok))
	return req, nil
}

// CreateDatabaseToken creates a token for a database owned by an organization with an optional given expiration and authorization.
func CreateDatabaseToken(orgName string, dbName string, apiTok string, opts ...CreateDatabaseTokenOption) (Jwt, error) {
	req, err := createCreateDatabaseTokenRequest(orgName, dbName, apiTok, opts...)
	if err != nil {
		return Jwt{}, fmt.Errorf("failed to create request for database token: %v", err)
	}
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return Jwt{}, fmt.Errorf("failed to send request for database token: %v", err)
	}
	jwt, err := parseResponse[Jwt](resp)
	if err != nil {
		return Jwt{}, fmt.Errorf("failed to parse response body: %v", err)
	}
	defer resp.Body.Close()
	return jwt, nil
}

// createListDatabasesRequest creates a request for listing all databases in an organization.
func createListDatabasesRequest(orgName string, orgToken string) (*http.Request, error) {
	url := fmt.Sprintf(
		"%s/organizations/%s/databases",
		tursoEndpoint, orgName,
	)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error reading request: %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", orgToken))
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

// ListDatabases lists all databases in an organization.
func ListDatabases(orgName string, orgToken string) (Databases, error) {
	req, err := createListDatabasesRequest(orgName, orgToken)
	if err != nil {
		return Databases{}, fmt.Errorf("error reading request: %v", err)
	}
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return Databases{}, fmt.Errorf("error sending request: %v", err)
	}
	dbs, err := parseResponse[Databases](resp)
	if err != nil {
		return Databases{}, fmt.Errorf("error decoding body: %v", err)
	}
	defer resp.Body.Close()
	return dbs, nil
}
