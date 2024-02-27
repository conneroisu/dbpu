package pkg

import (
	"fmt"
	"github.com/bytedance/sonic/decoder"
	"github.com/charmbracelet/log"
	"io"
	"net/http"
)

// curl -L -X POST 'https://api.turso.tech/v1/organizations/{organizationName}/databases/{databaseName}/auth/tokens?expiration=2w&authorization=full-access' \ -H 'Authorization: Bearer TOKEN'
func CreateDatabaseToken(organizationName string, databaseName string, token string, expiration string, authorization string) (Jwt, error) {
	url := fmt.Sprintf("https://api.turso.tech/v1/organizations/%s/databases/%s/auth/tokens?expiration=%s&authorization=%s", organizationName, databaseName, expiration, authorization)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Errorf("DBPU: CreateDatabaseToken: failed to create database token: %v", err)
		return Jwt{}, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Errorf("DBPU: CreateDatabaseToken: failed to create database token: %v", err)
		return Jwt{}, err
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
		return Jwt{}, err
	}
	err = decoder.NewDecoder(string(body)).Decode(&jwt)
	if err != nil {
		log.Errorf("DBPU: CreateDatabaseToken: failed to decode response body: %v", err)
		return Jwt{}, err
	}
	return jwt, nil
}
