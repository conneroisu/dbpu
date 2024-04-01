package dbpu

import "fmt"

// CreateToken creates a new API token with the given name.
func (c *Client) CreateToken(apiToken, tokenName string) (ApiToken, error) {
	req, err := c.newCreateTokenRequest(tokenName)
	if err != nil {
		return ApiToken{}, fmt.Errorf("failed to create request for token: %v", err)
	}
	done, err := c.Do(req)
	if err != nil {
		return ApiToken{}, fmt.Errorf("failed to create token: %v", err)
	}
	par, err := parseResponse[ApiToken](done)
	if err != nil {
		return ApiToken{}, fmt.Errorf("failed to parse response: %v", err)
	}
	defer done.Body.Close()
	return par, nil
}

// ValidateToken validates the given API token beloning to a user.
func (c *Client) ValidateToken(apiToken string) (ValidTokResp, error) {
	req, err := c.newValidateTokenRequest(apiToken)
	if err != nil {
		return ValidTokResp{}, fmt.Errorf("failed to create request for token validation: %v", err)
	}
	done, err := c.Do(req)
	if err != nil {
		return ValidTokResp{}, fmt.Errorf("failed to validate token: %v", err)
	}
	par, err := parseResponse[ValidTokResp](done)
	if err != nil {
		return ValidTokResp{}, fmt.Errorf("failed to parse response: %v", err)
	}
	defer done.Body.Close()
	return par, nil
}

// ListTokens lists the API tokens for the user.
func (c *Client) ListTokens(apiToken string) (ListToksResp, error) {
	req, err := c.newListTokensRequest(apiToken)
	if err != nil {
		return ListToksResp{}, fmt.Errorf("failed to create request for listing tokens: %v", err)
	}
	done, err := c.Do(req)
	if err != nil {
		return ListToksResp{}, fmt.Errorf("failed to list tokens: %v", err)
	}
	par, err := parseResponse[ListToksResp](done)
	if err != nil {
		return ListToksResp{}, fmt.Errorf("failed to parse response: %v", err)
	}
	defer done.Body.Close()
	return par, nil
}

// RevokeToken revokes the given API token.
func (c *Client) RevokeToken(apiToken, tokenName string) (RevokeTokResp, error) {
	req, err := c.newRevokeTokenRequest(apiToken, tokenName)
	if err != nil {
		return RevokeTokResp{}, fmt.Errorf("failed to create request for revoking token: %v", err)
	}
	done, err := c.Do(req)
	if err != nil {
		return RevokeTokResp{}, fmt.Errorf("failed to revoke token: %v", err)
	}
	par, err := parseResponse[RevokeTokResp](done)
	if err != nil {
		return RevokeTokResp{}, fmt.Errorf("failed to parse response: %v", err)
	}
	defer done.Body.Close()
	return par, nil
}
