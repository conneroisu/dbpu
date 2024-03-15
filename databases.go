package dbpu

import (
	"bytes"
	"fmt"
	"net/http"
)

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
func (c *Client) CreateDatabase(orgToken, orgName, name, group string) (Db, error) {
	req, reqErr := newCreateDatabaseReq(orgToken, orgName, name, group)
	done, doErr := c.Do(req)
	parsed, parErr := parseResponse[DbResp](done)
	defer done.Body.Close()
	return resolveApiCall(parsed.Database, wReqError(reqErr), wDoError(doErr), wParError(parErr))
}

// CreateDatabaseToken creates a token for a database owned by an organization with an optional given expiration and authorization.
func (c *Client) CreateDatabaseToken(orgName, dbName, apiTok string, opts ...newDbTokenOpt) (Jwt, error) {
	config := newDbTokenConfig(opts...)
	req, reqErr := newCreateDatabaseTokenReq(orgName, dbName, apiTok, config)
	done, doErr := c.Do(req)
	jwt, parErr := parseResponse[Jwt](done)
	defer done.Body.Close()
	return resolveApiCall(jwt, wReqError(reqErr), wDoError(doErr), wParError(parErr))
}

// ListDatabases lists all databases for an organization.
func (c *Client) ListDatabases(orgName, orgToken string) (Dbs, error) {
	req, reqErr := newListDatabasesReq(orgName, orgToken)
	done, doErr := c.Do(req)
	dbs, parErr := parseResponse[Dbs](done)
	defer done.Body.Close()
	return resolveApiCall(dbs, wReqError(reqErr), wDoError(doErr), wParError(parErr))
}

// newListDatabasesReq creates a request for listing all databases in an organization.
func newListDatabasesReq(orgName, orgToken string) (*http.Request, error) {
	url := fmt.Sprintf("%s/organizations/%s/databases", tursoEndpoint, orgName)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error reading request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", orgToken))
	return req, nil
}

// newCreateDatabaseTokenReq creates a request for creating a token for a database owned by an organization.
func newCreateDatabaseTokenReq(orgName, dbName, apiTok string, config *DbTokenConfig) (*http.Request, error) {
	url := fmt.Sprintf(
		"%s/organizations/%s/databases/%s/auth/tokens?expiration=%s&authorization=%s",
		tursoEndpoint, orgName, dbName, config.expiration, config.authorization,
	)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request for database token: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiTok))
	return req, nil
}

// newCreateDatabaseReq returns a new http.Request for creating a database.
func newCreateDatabaseReq(orgToken, orgName, name, group string) (*http.Request, error) {
	url := fmt.Sprintf(
		"%s/organizations/%s/databases",
		tursoEndpoint, orgName,
	)
	reqJsonBody := fmt.Sprintf(`{ "name": "%s", "group": "%s" }`, name, group)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(reqJsonBody)))
	if err != nil {
		return nil, fmt.Errorf("error reading request. %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", orgToken))
	return req, nil
}
