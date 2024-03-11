package dbpu

import (
	"fmt"
	"io"
	"net/http"
)

// ApiToken is a response to creating a new API token.
type ApiToken struct {
	ID    string `json:"id"`    // The ID of the token.
	Name  string `json:"name"`  // The name of the token.
	Token string `json:"token"` // The token.
}

// Token is a response to listing API tokens.
type Token struct {
	Id   string `json:"id"`   // The ID of the token.
	Name string `json:"name"` // The name of the token.
}

// ListTokens lists the API tokens for the user.
type ListTokensResponse struct {
	Tokens []Token `json:"tokens"`
}

// ValidateTokea is a response to creating a new API token.
type ValidateTokenResponse struct {
	Exp int `json:"exp"` // The expiration time of the token.
}

// RevokeTokenResponse is a response to revoking an API token.
type RevokeTokenResponse struct {
	Token string `json:"token"` // The token that was revoked.
}

// CreateCreateTokenRequest creates a request for creating a new API token.j
func CreateCreateTokenRequest(tokenName string) (*http.Request, error) {
	url := fmt.Sprintf(tursoEndpoint+"/auth/api-tokens/%s", tokenName)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Error creating request. %v", err)
	}
	return req, nil
}

// CreateToken creates a new API token with the given name.
func CreateToken(apiToken string, tokenName string) (ApiToken, error) {
	req, err := CreateCreateTokenRequest(tokenName)
	if err != nil {
		return ApiToken{}, fmt.Errorf("Error reading request. %v", err)
	}
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return ApiToken{}, fmt.Errorf("Error sending request. %v", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ApiToken{}, fmt.Errorf("Error reading body. %v", err)
	}
	apiTokenResponse, err := parseStruct[ApiToken](body)
	if err != nil {
		return ApiToken{}, fmt.Errorf("Error decoding body. %v", err)
	}
	defer resp.Body.Close()
	return apiTokenResponse, nil
}

// ValidateToken validates the given API token beloning to a user.
func ValidateToken(apiToken string) (ValidateTokenResponse, error) {
	url := fmt.Sprintf(tursoEndpoint + "/auth/validate")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ValidateTokenResponse{}, fmt.Errorf("Error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return ValidateTokenResponse{}, fmt.Errorf("Error sending request. %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ValidateTokenResponse{}, fmt.Errorf("Error reading body. %v", err)
	}
	parseDatabaseResponse, err := parseStruct[ValidateTokenResponse](body)
	if err != nil {
		return ValidateTokenResponse{}, fmt.Errorf("Error decoding body. %v", err)
	}
	return parseDatabaseResponse, nil
}

// ListTokens lists the API tokens for the user.
func ListTokens(apiToken string) (ListTokensResponse, error) {
	url := fmt.Sprintf(tursoEndpoint + "/auth/api-tokens")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ListTokensResponse{}, fmt.Errorf("Error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return ListTokensResponse{}, fmt.Errorf("Error sending request. %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ListTokensResponse{}, fmt.Errorf("Error reading body. %v", err)
	}
	parseDatabaseResponse, err := parseStruct[ListTokensResponse](body)
	if err != nil {
		return ListTokensResponse{}, fmt.Errorf("Error decoding body. %v", err)
	}
	return parseDatabaseResponse, nil
}

// RevokeToken revokes the given API token.
func RevokeToken(apiToken string, tokenName string) (RevokeTokenResponse, error) {
	url := fmt.Sprintf(tursoEndpoint + "/auth/api-tokens/" + tokenName)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return RevokeTokenResponse{}, fmt.Errorf("Error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return RevokeTokenResponse{}, fmt.Errorf("Error sending request. %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return RevokeTokenResponse{}, fmt.Errorf("Error reading body. %v", err)
	}
	parseDatabaseResponse, err := parseStruct[RevokeTokenResponse](body)
	if err != nil {
		return RevokeTokenResponse{}, fmt.Errorf("Error decoding body. %v", err)
	}
	return parseDatabaseResponse, nil
}
