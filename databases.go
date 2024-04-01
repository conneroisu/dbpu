package dbpu

import "fmt"

// CreateDatabase creates a database with the given name and group.
func (c *Client) CreateDatabase(orgToken, orgName, name, group string) (Database, error) {
	req, err := c.newCreateDatabaseReq(orgToken, orgName, name, group)
	if err != nil {
		return Database{}, fmt.Errorf("failed to create request for database: %v", err)
	}
	done, err := c.Do(req)
	if err != nil {
		return Database{}, fmt.Errorf("failed to create database: %v", err)
	}
	defer done.Body.Close()
	parsed, err := parseResponse[DbResp](done)
	if err != nil {
		return Database{}, fmt.Errorf("failed to parse response: %v", err)
	}
	return parsed.Database, nil
}

// CreateDatabaseToken creates a token for a database owned by an organization with an optional given expiration and authorization.
func (c *Client) CreateDatabaseToken(orgName, dbName, apiTok string, opts ...newDbTokenOpt) (*Jwt, error) {
	config := newDbTokenConfig(opts...)
	req, err := c.newCreateDatabaseTokenReq(orgName, dbName, apiTok, config)
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
	return &jwt, nil
}

// ListDatabases lists all databases for an organization.
func (c *Client) ListDatabases(orgName, orgToken string) (Dbs, error) {
	req, err := c.newListDatabasesReq(orgName, orgToken)
	if err != nil {
		return Dbs{}, fmt.Errorf("failed to create request for listing databases: %v", err)
	}
	done, err := c.Do(req)
	if err != nil {
		return Dbs{}, fmt.Errorf("failed to list databases: %v", err)
	}
	defer done.Body.Close()
	dbs, err := parseResponse[Dbs](done)
	if err != nil {
		return Dbs{}, fmt.Errorf("failed to parse response: %v", err)
	}
	return dbs, nil
}
