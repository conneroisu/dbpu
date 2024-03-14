package dbpu

import (
	"fmt"
	"net/http"
)

// CreateToken creates a new API token with the given name.
func CreateToken(apiToken, tokenName string) (ApiToken, error) {
	req, reqErr := newCreateTokenRequest(tokenName)
	done, doErr := (&http.Client{}).Do(req)
	apiTokenResp, parErr := parseResponse[ApiToken](done)
	defer done.Body.Close()
	return resolveApiCall(apiTokenResp, wReqError(reqErr), wDoError(doErr), wParError(parErr))
}

// ValidateToken validates the given API token beloning to a user.
func ValidateToken(apiToken string) (ValidTokResp, error) {
	req, reqErr := newValidateTokenRequest(apiToken)
	done, doErr := (&http.Client{}).Do(req)
	parseDatabaseResponse, parErr := parseResponse[ValidTokResp](done)
	defer done.Body.Close()
	return resolveApiCall(parseDatabaseResponse, wReqError(reqErr), wDoError(doErr), wParError(parErr))
}

// ListTokens lists the API tokens for the user.
func ListTokens(apiToken string) (ListToksResp, error) {
	req, reqErr := newListTokensRequest(apiToken)
	done, doErr := (&http.Client{}).Do(req)
	parsed, respErr := parseResponse[ListToksResp](done)
	defer done.Body.Close()
	return resolveApiCall(parsed, wReqError(reqErr), wDoError(doErr), wParError(respErr))
}

// RevokeToken revokes the given API token.
func RevokeToken(apiToken, tokenName string) (RevokeTokResp, error) {
	req, reqErr := newRevokeTokenRequest(apiToken, tokenName)
	done, doErr := (&http.Client{}).Do(req)
	revokeTokResponse, parErr := parseResponse[RevokeTokResp](done)
	defer done.Body.Close()
	return resolveApiCall(revokeTokResponse, wReqError(reqErr), wDoError(doErr), wParError(parErr))
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
func newRevokeTokenRequest(apiToken, tokenName string) (*http.Request, error) {
	url := fmt.Sprintf("%s/auth/api-tokens/%s", tursoEndpoint, tokenName)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	return req, nil
}
