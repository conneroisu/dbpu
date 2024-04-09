package dbpu

import (
	"bytes"
	"fmt"
	"net/http"
)

type CreateDatabaseBody struct {
	Name       string  `json:"name"`
	Location   string  `json:"location"`
	Image      string  `json:"image,omitempty"`
	Extensions string  `json:"extensions,omitempty"`
	Group      string  `json:"group,omitempty"`
	Seed       *DBSeed `json:"seed,omitempty"`
	Schema     string  `json:"schema,omitempty"`
	IsSchema   bool    `json:"is_schema,omitempty"`
}

// newListDatabasesReq creates a request for listing all databases in an organization.
func (c *Client) newListDatabasesReq() (*http.Request, error) {
	url := fmt.Sprintf("%s/organizations/%s/databases", c.BaseURL, c.OrgName)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error reading request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.OrgToken))
	return req, nil
}

// newCreateDatabaseTokenReq creates a request for creating a token for a database owned by an organization.
func (c *Client) newCreateDatabaseTokenReq(dbName, apiTok string, config *DbTokenConfig) (*http.Request, error) {
	url := fmt.Sprintf(
		"%s/organizations/%s/databases/%s/auth/tokens?expiration=%s&authorization=%s",
		c.BaseURL, c.OrgName, dbName, config.expiration, config.authorization,
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
func (c *Client) newCreateDatabaseReq(name, group string) (*http.Request, error) {
	url := fmt.Sprintf("%s/organizations/%s/databases", c.BaseURL, c.OrgName)
	reqJsonBody := fmt.Sprintf(`{ "name": "%s", "group": "%s" }`, name, group)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(reqJsonBody)))
	if err != nil {
		return nil, fmt.Errorf("error reading request. %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.OrgToken))
	return req, nil
}
