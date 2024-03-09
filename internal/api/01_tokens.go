package api

import (
	"fmt"
	"io"
	"net/http"

	"github.com/bytedance/sonic/decoder"
)

// ApiToken is a response to creating a new API token.
type ApiToken struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Token string `json:"token"`
}

// CreateToken creates a new API token with the given name.
func CreateToken(apiToken string, tokenName string) (ApiToken, error) {
	url := fmt.Sprintf(TursoEndpoint+"/auth/api-tokens/%s", tokenName)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return ApiToken{}, fmt.Errorf("Error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return ApiToken{}, fmt.Errorf("Error sending request. %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ApiToken{}, fmt.Errorf("Error reading body. %v", err)
	}
	var apiTokenResponse ApiToken
	err = decoder.NewDecoder(string(body)).Decode(&apiTokenResponse)
	if err != nil {
		return ApiToken{}, fmt.Errorf("Error decoding body. %v", err)
	}
	return apiTokenResponse, nil
}

// ValidateTokea is a response to creating a new API token.
type ValidateTokenResponse struct {
	Exp int `json:"exp"`
}

//	curl -L 'https://api.turso.tech/v1/auth/validate' \
//	  -H 'Authorization: Bearer TOKEN'
//
// ValidateToken validates the given API token beloning to a user.
func ValidateToken(apiToken string) (ValidateTokenResponse, error) {
	url := fmt.Sprintf(TursoEndpoint + "/auth/validate")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ValidateTokenResponse{}, fmt.Errorf("Error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return ValidateTokenResponse{}, fmt.Errorf("Error sending request. %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ValidateTokenResponse{}, fmt.Errorf("Error reading body. %v", err)
	}
	parseDatabaseResponse, err := parseValidateTokenResponse(body, err)
	if err != nil {
		return ValidateTokenResponse{}, fmt.Errorf("Error decoding body. %v", err)
	}
	return parseDatabaseResponse, nil
}

// parseDatabaseResponse parses the response from the server into a DatabaseResponse.
func parseValidateTokenResponse(body []byte, err error) (ValidateTokenResponse, error) {
	var response ValidateTokenResponse
	err = decoder.NewDecoder(string(body)).Decode(&response)
	return response, err
}

// Token is a response to listing API tokens.
type Token struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// ListTokens lists the API tokens for the user.
type ListTokensResponse struct {
	Tokens []Token `json:"tokens"`
}

// ListTokens lists the API tokens for the user.
func ListTokens(apiToken string) (ListTokensResponse, error) {
	url := fmt.Sprintf(TursoEndpoint + "/auth/api-tokens")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ListTokensResponse{}, fmt.Errorf("Error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return ListTokensResponse{}, fmt.Errorf("Error sending request. %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ListTokensResponse{}, fmt.Errorf("Error reading body. %v", err)
	}
	parseDatabaseResponse, err := parseListTokensResponse(body, err)
	if err != nil {
		return ListTokensResponse{}, fmt.Errorf("Error decoding body. %v", err)
	}
	return parseDatabaseResponse, nil
}

// parseDatabaseResponse parses the response from the server into a DatabaseResponse.
func parseListTokensResponse(body []byte, err error) (ListTokensResponse, error) {
	var response ListTokensResponse
	err = decoder.NewDecoder(string(body)).Decode(&response)
	return response, err
}

// RevokeTokenResponse is a response to revoking an API token.
type RevokeTokenResponse struct {
	Token string `json:"token"`
}

// curl -L -X DELETE 'https://api.turso.tech/v1/auth/api-tokens/{tokenName}' \
// -H 'Authorization: Bearer TOKEN'
func RevokeToken(apiToken string, tokenName string) (RevokeTokenResponse, error) {
	url := fmt.Sprintf(TursoEndpoint + "/auth/api-tokens/" + tokenName)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return RevokeTokenResponse{}, fmt.Errorf("Error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return RevokeTokenResponse{}, fmt.Errorf("Error sending request. %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return RevokeTokenResponse{}, fmt.Errorf("Error reading body. %v", err)
	}
	parseDatabaseResponse, err := parseRevokeTokenResponse(body, err)
	if err != nil {
		return RevokeTokenResponse{}, fmt.Errorf("Error decoding body. %v", err)
	}
	return parseDatabaseResponse, nil
}

// parseDatabaseResponse parses the response from the server into a DatabaseResponse.
func parseRevokeTokenResponse(body []byte, err error) (RevokeTokenResponse, error) {
	var response RevokeTokenResponse
	err = decoder.NewDecoder(string(body)).Decode(&response)
	return response, err
}
