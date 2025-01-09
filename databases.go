package dbpu

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/conneroisu/dbpu/internal/builders"
)

// Database is a libsql database.
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

// Config is a struct configures the creation of a database.
type Config struct {
	Name       string `json:"name" validate:"required"`
	Location   string `json:"location" validate:"required"`
	Image      string `json:"image,omitempty"`
	Extensions string `json:"extensions,omitempty"`
	Group      string `json:"group,omitempty"`
	Seed       *Seed  `json:"seed,omitempty"`
	Schema     string `json:"schema,omitempty"`
	IsSchema   bool   `json:"is_schema,omitempty"`
}

// Seed is a seed for a database.
type Seed struct {
	Type      string     `json:"type"`
	Name      string     `json:"value,omitempty"`
	URL       string     `json:"url,omitempty"`
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

// Create creates a database with the given name and group.
//
// Options can be provided to configure the database such as location, image,
// extensions, seed, schema, and isSchema.
func (c *Client) Create(ctx context.Context, config Config) (*Database, error) {
	req, err := builders.NewRequest(
		ctx,
		c.header,
		http.MethodPost,
		fmt.Sprintf("%s/organizations/%s/databases", c.BaseURL, c.OrgName),
		builders.WithBody(config),
	)
	if err != nil {
		return nil, err
	}
	var resp struct {
		Database Database `json:"database"`
	}
	err = c.sendRequest(req, resp)
	if err != nil {
		return nil, fmt.Errorf("failed to create database: %v", err)
	}
	return &resp.Database, nil
}

type (
	newDbTokenOpt func(*TokenConfig)

	// TokenConfig is a configuration for creating a database token.
	TokenConfig struct {
		// Expiration time for the token (e.g., 2w1d30m).
		expiration string `url:"expiration"`
		// Authorization level for the token (full-access or read-only).
		authorization string `url:"authorization"`
	}
)

// WithExpiration sets the expiration time for the token (e.g., 2w1d30m).
func WithExpiration(expiration string) func(*TokenConfig) {
	return func(c *TokenConfig) { c.expiration = expiration }
}

// WithAuthorization sets the authorization level for the token (full-access or read-only).
func WithAuthorization(authorization string) func(*TokenConfig) {
	return func(c *TokenConfig) { c.authorization = authorization }
}

// CreateDatabaseToken creates a token for a database owned by an organization
// with an optional given expiration and authorization.
func (c *Client) CreateDatabaseToken(
	ctx context.Context,
	dbName string,
	opts ...newDbTokenOpt,
) (string, error) {
	config := TokenConfig{}
	for _, opt := range opts {
		opt(&config)
	}
	uri, err := url.Parse(fmt.Sprintf(
		"%s/organizations/%s/databases/%s/auth/tokens",
		c.BaseURL, c.OrgName, dbName,
	))
	if err != nil {
		return "", err
	}
	vals, err := builders.Values(config)
	if err != nil {
		return "", err
	}
	uri.RawQuery = vals.Encode()
	var resp struct {
		Token string `json:"jwt"`
	}
	req, err := builders.NewRequest(
		ctx,
		c.header,
		http.MethodPost,
		uri.String(),
	)
	if err != nil {
		return "", err
	}
	err = c.sendRequest(req, &resp)
	return resp.Token, err
}
