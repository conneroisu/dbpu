package dbpu

import (
	"testing"
)

// TestParsing tests the parsing functions in respect to the strcutures declared in tokens.go.
func TestParsing(t *testing.T) {
	// Test the parse api token response with a test id, name and token.
	t.Run("Test Parse Api Token", func(t *testing.T) {
		body := []byte(`
			{
				"id":"test",
				"name":"test",
				"token":"test"
			}
		`)
		apiToken, err := parseStruct[ApiToken](body)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if apiToken.ID != "test" {
			t.Errorf("Expected ID to be test, got %v", apiToken.ID)
		}
		if apiToken.Name != "test" {
			t.Errorf("Expected Name to be test, got %v", apiToken.Name)
		}
		if apiToken.Token != "test" {
			t.Errorf("Expected Token to be test, got %v", apiToken.Token)
		}
	})

	// Test the parse token response with a test id and name.
	t.Run("Test Parse Token", func(t *testing.T) {
		body := []byte(`{"id":"test","name":"test"}`)
		token, err := parseStruct[Token](body)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if token.Id != "test" {
			t.Errorf("Expected Id to be test, got %v", token.Id)
		}
		if token.Name != "test" {
			t.Errorf("Expected Name to be test, got %v", token.Name)
		}
	})

	// Test the parse list tokens response with a test id and name.
	t.Run("Test Parse List Tokens Response", func(t *testing.T) {
		body := []byte(`{"tokens":[{"id":"test","name":"test"}]}`)
		listTokensResponse, err := parseStruct[ListToksResp](body)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if len(listTokensResponse.Tokens) != 1 {
			t.Errorf("Expected Tokens to have 1 element, got %v", len(listTokensResponse.Tokens))
		}
		if listTokensResponse.Tokens[0].Id != "test" {
			t.Errorf("Expected Id to be test, got %v", listTokensResponse.Tokens[0].Id)
		}
		if listTokensResponse.Tokens[0].Name != "test" {
			t.Errorf("Expected Name to be test, got %v", listTokensResponse.Tokens[0].Name)
		}
	})

	// Test the parse validate token response.
	t.Run("Test Parse Validate Token Response", func(t *testing.T) {
		body := []byte(`{"exp":1}`)
		validateTokenResponse, err := parseStruct[ValidTokResp](body)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if validateTokenResponse.Exp != 1 {
			t.Errorf("Expected Exp to be 1, got %v", validateTokenResponse.Exp)
		}
	})

	// Test the parse revoke token response.
	t.Run("Test Parse Revoke Token Response", func(t *testing.T) {
		body := []byte(`{"token":"test"}`)
		revokeTokenResponse, err := parseStruct[RevokeTokResp](body)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if revokeTokenResponse.Token != "test" {
			t.Errorf("Expected Token to be test, got %v", revokeTokenResponse.Token)
		}
	})
}

// TestCreateCreateTokenRquest tests the create create token request function.
func TestCreateCreateTokenRquest(t *testing.T) {
	// Test the create create token request.
	t.Run("Test Create Create Token Request", func(t *testing.T) {
		tokenName := "test"
		req, err := newCreateTokenRequest(tokenName)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if req.Method != "POST" {
			t.Errorf("Expected method to be POST, got %v", req.Method)
		}
		if req.URL.String() != tursoEndpoint+"/auth/api-tokens/"+tokenName {
			t.Errorf("Expected URL to be %s, got %v", tursoEndpoint+"/auth/api-tokens/"+tokenName, req.URL.String())
		}
	})
}
