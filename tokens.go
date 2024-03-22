package dbpu

// CreateToken creates a new API token with the given name.
func (c *Client) CreateToken(apiToken, tokenName string) (ApiToken, error) {
	req, reqErr := c.newCreateTokenRequest(tokenName)
	done, doErr := c.Do(req)
	apiTokenResp, parErr := parseResponse[ApiToken](done)
	defer done.Body.Close()
	return resolveApi(apiTokenResp, wReqError(reqErr), wDoError(doErr), wParError(parErr))
}

// ValidateToken validates the given API token beloning to a user.
func (c *Client) ValidateToken(apiToken string) (ValidTokResp, error) {
	req, reqErr := c.newValidateTokenRequest(apiToken)
	done, doErr := c.Do(req)
	par, parErr := parseResponse[ValidTokResp](done)
	defer done.Body.Close()
	return resolveApi(par, wReqError(reqErr), wDoError(doErr), wParError(parErr))
}

// ListTokens lists the API tokens for the user.
func (c *Client) ListTokens(apiToken string) (ListToksResp, error) {
	req, reqErr := c.newListTokensRequest(apiToken)
	done, doErr := c.Do(req)
	par, respErr := parseResponse[ListToksResp](done)
	defer done.Body.Close()
	return resolveApi(par, wReqError(reqErr), wDoError(doErr), wParError(respErr))
}

// RevokeToken revokes the given API token.
func (c *Client) RevokeToken(apiToken, tokenName string) (RevokeTokResp, error) {
	req, reqErr := c.newRevokeTokenRequest(apiToken, tokenName)
	done, doErr := c.Do(req)
	par, parErr := parseResponse[RevokeTokResp](done)
	defer done.Body.Close()
	return resolveApi(par, wReqError(reqErr), wDoError(doErr), wParError(parErr))
}
