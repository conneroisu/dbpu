package dbpu

import (
	"fmt"
	"net/http"
)

// newCreateTokenRequest creates a request for creating a new API token.
func (c *Client) newCreateTokenRequest(tokenName string) (*http.Request, error) {
	url := fmt.Sprintf("%s/auth/api-tokens/%s", c.BaseURL, tokenName)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request. %v", err)
	}
	return req, nil
}

// newValidateTokenRequest creates a request for validating an API token.
func (c *Client) newValidateTokenRequest(apiToken string) (*http.Request, error) {
	url := fmt.Sprintf("%s/auth/validate", c.BaseURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	return req, nil
}

// newListTokensRequest creates a request for listing API tokens.
func (c *Client) newListTokensRequest(apiToken string) (*http.Request, error) {
	url := fmt.Sprintf("%s/auth/api-tokens", c.BaseURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	return req, nil
}

// newRevokeTokenRequest creates a request for revoking an API token.
func (c *Client) newRevokeTokenRequest(apiToken, tokenName string) (*http.Request, error) {
	url := fmt.Sprintf("%s/auth/api-tokens/%s", c.BaseURL, tokenName)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	return req, nil
}
