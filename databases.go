package dbpu

import (
	"time"
)

// Database is a database.
type Database struct {
	ID            string   `json:"DbId"`
	Hostname      string   `json:"Hostname"`
	Name          string   `json:"Name"`
	Group         string   `json:"group"`
	PrimaryRegion string   `json:"primaryRegion"`
	Regions       []string `json:"regions"`
	Type          string   `json:"type"`
	Version       string   `json:"version"`
}

type DBSeed struct {
	Type      string     `json:"type"`
	Name      string     `json:"value,omitempty"`
	URL       string     `json:"url,omitempty"`
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

// NewCreateDatabaseTokenConfig returns a new CreateDatabaseTokenConfig.
func newDbTokenConfig(opts ...newDbTokenOpt) *DbTokenConfig {
	c := &DbTokenConfig{
		expiration:    "",
		authorization: "",
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// WithExpiration sets the expiration time for the token (e.g., 2w1d30m).
func WithExpiration(expiration string) newDbTokenOpt {
	return func(c *DbTokenConfig) { c.expiration = expiration }
}

// WithAuthorization sets the authorization level for the token (full-access or read-only).
func WithAuthorization(authorization string) newDbTokenOpt {
	return func(c *DbTokenConfig) { c.authorization = authorization }
}

// CreateDatabase creates a database with the given name and group.
func (c *Client) CreateDatabase(orgToken, orgName, name, group string) (Database, error) {
	req, err := c.newCreateDatabaseReq(orgToken, orgName, name, group)
	if err != nil {
		return Database{}, err
	}
	done, err := c.Do(req)
	if err != nil {
		return Database{}, err
	}
	defer done.Body.Close()
	parsed, err := parseResponse[DbResp](done)
	if err != nil {
		return Database{}, err
	}
	return parsed.Database, nil
}

// CreateDatabaseToken creates a token for a database owned by an organization with an optional given expiration and authorization.
func (c *Client) CreateDatabaseToken(orgName, dbName, apiTok string, opts ...newDbTokenOpt) (Jwt, error) {
	config := newDbTokenConfig(opts...)
	req, err := c.newCreateDatabaseTokenReq(orgName, dbName, apiTok, config)
	if err != nil {
		return Jwt{}, err
	}
	done, err := c.Do(req)
	if err != nil {
		return Jwt{}, err
	}
	defer done.Body.Close()
	jwt, err := parseResponse[Jwt](done)
	if err != nil {
		return Jwt{}, err
	}
	return jwt, nil
}

// ListDatabases lists all databases for an organization.
func (c *Client) ListDatabases(orgName, orgToken string) (Dbs, error) {
	req, err := c.newListDatabasesReq(orgName, orgToken)
	if err != nil {
		return Dbs{}, err
	}
	done, err := c.Do(req)
	if err != nil {
		return Dbs{}, err
	}
	defer done.Body.Close()
	dbs, err := parseResponse[Dbs](done)
	if err != nil {
		return Dbs{}, err
	}
	return dbs, nil
}
