package dbpu

import (
	"fmt"
	"net/http"
)

// ApiToken is a response to creating a new API token.
type ApiToken struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Token string `json:"token"`
}

// Token is a response to listing API tokens.
type Token struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// ListToksResp is a response to listing API tokens.
type ListToksResp struct {
	Tokens []Token `json:"tokens"`
}

// ValidTokResp is a response to creating a new API token.
type ValidTokResp struct {
	Exp int `json:"exp"`
}

// RevokeTokResp is a response to revoking an API token.
type RevokeTokResp struct {
	Token string `json:"token"`
}

// CreateToken creates a new API token with the given name.
func CreateToken(apiToken string, tokenName string) (ApiToken, error) {
	req, reqErr := newCreateTokenRequest(tokenName)
	resp, doErr := (&http.Client{}).Do(req)
	apiTokenResp, parErr := parseResponse[ApiToken](resp)
	defer resp.Body.Close()
	return resolveApiCall[ApiToken](apiTokenResp, reqErr, doErr, parErr)
}

// ValidateToken validates the given API token beloning to a user.
func ValidateToken(apiToken string) (ValidTokResp, error) {
	req, reqErr := newValidateTokenRequest(apiToken)
	resp, doErr := (&http.Client{}).Do(req)
	parseDatabaseResponse, parErr := parseResponse[ValidTokResp](resp)
	defer resp.Body.Close()
	return resolveApiCall[ValidTokResp](parseDatabaseResponse, reqErr, doErr, parErr)
}

// ListTokens lists the API tokens for the user.
func ListTokens(apiToken string) (ListToksResp, error) {
	req, reqErr := newListTokensRequest(apiToken)
	resp, doErr := (&http.Client{}).Do(req)
	parsed, respErr := parseResponse[ListToksResp](resp)
	defer resp.Body.Close()
	return resolveApiCall[ListToksResp](parsed, reqErr, doErr, respErr)
}

// RevokeToken revokes the given API token.
func RevokeToken(apiToken string, tokenName string) (RevokeTokResp, error) {
	req, reqErr := newRevokeTokenRequest(apiToken, tokenName)
	resp, doErr := (&http.Client{}).Do(req)
	revokeTokResponse, parErr := parseResponse[RevokeTokResp](resp)
	defer resp.Body.Close()
	return resolveApiCall[RevokeTokResp](revokeTokResponse, reqErr, doErr, parErr)
}

// newCreateTokenRequest creates a request for creating a new API token.
func newCreateTokenRequest(tokenName string) (*http.Request, error) {
	url := fmt.Sprintf("%s/auth/api-tokens/%s", tursoEndpoint, tokenName)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request. %v", err)
	}
	return req, nil
}

// newValidateTokenRequest creates a request for validating an API token.
func newValidateTokenRequest(apiToken string) (*http.Request, error) {
	url := fmt.Sprintf("%s/auth/validate", tursoEndpoint)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	return req, nil
}

// newListTokensRequest creates a request for listing API tokens.
func newListTokensRequest(apiToken string) (*http.Request, error) {
	url := fmt.Sprintf("%s/auth/api-tokens", tursoEndpoint)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	return req, nil
}

// newRevokeTokenRequest creates a request for revoking an API token.
func newRevokeTokenRequest(apiToken string, tokenName string) (*http.Request, error) {
	url := fmt.Sprintf("%s/auth/api-tokens/%s", tursoEndpoint, tokenName)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	return req, nil
}
