package api

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/bytedance/sonic/decoder"
	"github.com/charmbracelet/log"
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

// CreateDatabaseResponse is a response to creating a database.
type CreateDatabaseResponse struct {
	Database Database `json:"database"`
}

// CreateDatabase creates a database with the given name and group.
func CreateDatabase(orgName string, orgToken string, dbName string, dbGroup string) (CreateDatabaseResponse, error) {
	url := fmt.Sprintf(TursoEndpoint+"/organizations/%s/databases", orgName)
	reqJsonBody := fmt.Sprintf(`{"name": "%s", "group": "%s"}`, dbName, dbGroup)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(reqJsonBody)))
	if err != nil {
		log.Errorf("Error creating request. %v", err)
		return CreateDatabaseResponse{}, fmt.Errorf("Error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", orgToken))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("Error sending request. %v", err)
		return CreateDatabaseResponse{}, fmt.Errorf("Error sending request. %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("Error reading body. %v", err)
		return CreateDatabaseResponse{}, fmt.Errorf("Error reading body. %v", err)
	}
	fmt.Println(string(body))
	response, err := parseDatabaseResponse(body, err)
	if err != nil {
		log.Errorf("Error decoding body into CreateDatabaseResponse. %v", err)
		return CreateDatabaseResponse{}, fmt.Errorf("Error decoding body. %v", err)
	}
	fmt.Println(response)
	return response, nil
}

// parseDatabaseResponse parses the response from creating a database.
func parseDatabaseResponse(body []byte, err error) (CreateDatabaseResponse, error) {
	var db CreateDatabaseResponse
	err = decoder.NewDecoder(string(body)).Decode(&db)
	return db, nil
}

// createDatabaseTokenConfig is a configuration for creating a database token.
type createDatabaseTokenConfig struct {
	expiration    string // Expiration time for the token (e.g., 2w1d30m).
	authorization string // Authorization level for the token (full-access or read-only).
}

// NewCreateDatabaseTokenConfig returns a new CreateDatabaseTokenConfig.
func newCreateDatabaseTokenConfig(opts ...CreateDatabaseTokenOption) *createDatabaseTokenConfig {
	c := &createDatabaseTokenConfig{expiration: "", authorization: ""} // Default values.
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// CreateDatabaseTokenOption is a functional option for configuring a CreateDatabaseTokenConfig.
type CreateDatabaseTokenOption func(*createDatabaseTokenConfig)

// WithExpiration sets the expiration time for the token (e.g., 2w1d30m).
func WithExpiration(expiration string) CreateDatabaseTokenOption {
	return func(c *createDatabaseTokenConfig) { c.expiration = expiration }
}

// WithAuthorization sets the authorization level for the token (full-access or read-only).
func WithAuthorization(authorization string) CreateDatabaseTokenOption {
	return func(c *createDatabaseTokenConfig) { c.authorization = authorization }
}

// Jwt is a JSON Web Token.
type Jwt struct {
	Jwt string `json:"jwt"` // jwt is the JSON Web Token.
}

// CreateDatabaseToken creates a token for a database owned by an organization with an optional given expiration and authorization.
func CreateDatabaseToken(orgName string, dbName string, apiTok string, opts ...CreateDatabaseTokenOption) (Jwt, error) {
	config := newCreateDatabaseTokenConfig(opts...)
	url := fmt.Sprintf(TursoEndpoint+"/organizations/%s/databases/%s/auth/tokens?expiration=%s&authorization=%s", orgName, dbName, config.expiration, config.authorization)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Errorf("DBPU: CreateDatabaseToken: failed to create database token: %v", err)
		return Jwt{}, fmt.Errorf("failed to create database token: %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiTok))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Errorf("DBPU: CreateDatabaseToken: failed to create database token: %v", err)
		return Jwt{}, fmt.Errorf("failed to create database token: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Errorf("DBPU: CreateDatabaseToken: failed to create database token: %s", resp.Status)
		return Jwt{}, fmt.Errorf("failed to create database token: %s", resp.Status)
	}
	var jwt Jwt
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("DBPU: CreateDatabaseToken: failed to read response body: %v", err)
		return Jwt{}, fmt.Errorf("failed to read response body: %v", err)
	}
	err = decoder.NewDecoder(string(body)).Decode(&jwt)
	if err != nil {
		log.Errorf("DBPU: CreateDatabaseToken: failed to decode response body: %v", err)
		return Jwt{}, fmt.Errorf("failed to decode response body: %v", err)
	}
	return jwt, nil
}

// Databases is a list of dbs.
type Databases struct {
	Databases []Database `json:"databases"`
}

// ListDatabases lists all databases in an organization.
func ListDatabases(orgName string, orgToken string) (Databases, error) {
	url := fmt.Sprintf(TursoEndpoint+"/organizations/%s/databases", orgName)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Errorf("Error creating request. %v", err)
		return Databases{}, fmt.Errorf("Error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", orgToken))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("Error sending request. %v", err)
		return Databases{}, fmt.Errorf("Error sending request. %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("Error reading body. %v", err)
		return Databases{}, fmt.Errorf("Error reading body. %v", err)
	}
	response, err := parseDatabasesResponse(body, err)
	if err != nil {
		log.Errorf("Error decoding body into Databases. %v", err)
		return Databases{}, fmt.Errorf("Error decoding body. %v", err)
	}
	return response, nil
}

// parseDatabasesResponse parses the response from listing databases.
func parseDatabasesResponse(body []byte, err error) (Databases, error) {
	var dbs Databases
	err = decoder.NewDecoder(string(body)).Decode(&dbs)
	if err != nil {
		log.Errorf("Error decoding body into Databases. %v", err)
		return Databases{}, fmt.Errorf("Error decoding body. %v", err)
	}
	return dbs, nil
}
