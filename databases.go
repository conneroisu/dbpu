package dbpu

import "fmt"

// CreateDbConfig is a configuration for creating a database.
type CreateDbConfig struct {
	Name       string  `json:"name"`
	Location   string  `json:"location"`
	Image      string  `json:"image,omitempty"`
	Extensions string  `json:"extensions,omitempty"`
	Group      string  `json:"group,omitempty"`
	Seed       *DBSeed `json:"seed,omitempty"`
	Schema     string  `json:"schema,omitempty"`
	IsSchema   bool    `json:"is_schema,omitempty"`
}

// CreateDatabase creates a database with the given name and group.
// Options can be provided to configure the database such as location, image, extensions, seed, schema, and isSchema.
//
// Location is the region where the database will be created.
//
// Image is the image to use for the database.
//
// Extensions are the extensions to add to the database.
//
// Seed is the seed to use for the database.
//
// Schema is the schema to use for the database. Allows for the same schema to be used across multiple databases.
//
// IsSchema is a boolean that indicates if the database is a schema.
func (c *Client) CreateDatabase(config CreateDbConfig) (*Database, error) {
	req, err := c.newCreateDatabaseReq(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create request for database: %v", err)
	}
	done, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to create database: %v", err)
	}
	defer done.Body.Close()
	parsed, err := parseResponse[DbResp](done)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}
	return &parsed.Database, nil
}

// CreateDatabaseToken creates a token for a database owned by an organization with an optional given expiration and authorization.
func (c *Client) CreateDatabaseToken(dbName, apiTok string, opts ...newDbTokenOpt) (*Jwt, error) {
	config := newDbTokenConfig(opts...)
	req, err := c.newCreateDatabaseTokenReq(dbName, apiTok, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create request for database token: %v", err)
	}
	done, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to create database token: %v", err)
	}
	defer done.Body.Close()
	jwt, err := parseResponse[Jwt](done)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}
	return jwt, nil
}

// ListDatabases lists all databases for an organization.
func (c *Client) ListDatabases() (*Dbs, error) {
	req, err := c.newListDatabasesReq()
	if err != nil {
		return nil, fmt.Errorf("failed to create request for listing databases: %v", err)
	}
	done, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to list databases: %v", err)
	}
	defer done.Body.Close()
	dbs, err := parseResponse[Dbs](done)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}
	return dbs, nil
}
