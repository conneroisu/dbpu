package dbpu

import (
	"testing"
)

// TestCreateClosestLocationRequest tests the create closest location request function.
func testCreateClosestLocationRequest(t *testing.T) {
	// Test the create closest location request.
	req, err := CreateClosestLocationRequest()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if req.Method != "GET" {
		t.Errorf("Expected method to be GET, got %v", req.Method)
	}
}

// TestServerClientParse tests the server client parse function.
func testServerClientParse(t *testing.T) {
	// Test the parse struct.
	body := []byte(`{"server":"test","client":"test"}`)
	serverClient, err := parseStruct[ServerClient](body)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if serverClient.Server != "test" {
		t.Errorf("Expected Server to be test, got %v", serverClient.Server)
	}
	if serverClient.Client != "test" {
		t.Errorf("Expected Client to be test, got %v", serverClient.Client)
	}
}
