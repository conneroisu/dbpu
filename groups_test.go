package dbpu

import (
	"net/http"
	"testing"
)

func testParseGroups(t *testing.T) {
	t.Run("Testing ListGroupsResponse Parsing", func(t *testing.T) {
		body := []byte(`{"groups":[{"archived":false,"locations":["us-east-1"],"name":"test","primary":"us-east-1","uuid":"1"}]}`)
		response, err := parseStruct[ListGroupsResp](body)
		if err != nil {
			t.Errorf("Error decoding body. %v", err)
		}
		if len(response.Groups) != 1 {
			t.Errorf("Expected 1 group, got %d", len(response.Groups))
		}
		if response.Groups[0].Name != "test" {
			t.Errorf("Expected group name to be test, got %s", response.Groups[0].Name)
		}
	})

	t.Run("Testing ListGroupsResponse Parsing with bad body", func(t *testing.T) {
		body := []byte(`{"groups":[{"archived":false,"locations":["us-east-1"],"name":"test","primary":"us-east-1","uuid":"1"}]`)
		_, err := parseStruct[ListGroupsResp](body)
		if err == nil {
			t.Errorf("Expected error decoding body, got nil")
		}
	})

	t.Run("Testing GroupResponse Parsing", func(t *testing.T) {
		body := []byte(`{"group":{"archived":false,"locations":["us-east-1"],"name":"test","primary":"us-east-1","uuid":"1"}}`)
		response, err := parseStruct[GroupResp](body)
		if err != nil {
			t.Errorf("Error decoding body. %v", err)
		}
		if response.Group.Name != "test" {
			t.Errorf("Expected group name to be test, got %s", response.Group.Name)
		}
	})

	t.Run("Testing GroupResponse Parsing with bad body", func(t *testing.T) {
		body := []byte(`{"group":{"archived":false,"locations":["us-east-1"],"name":"test","primary":"us-east-1","uuid":"1"`)
		_, err := parseStruct[GroupResp](body)
		if err == nil {
			t.Errorf("Expected error decoding body, got nil")
		}
	})

	t.Run("Testing GroupResponse Parsing with bad body", func(t *testing.T) {
		body := []byte(`{"group":{"archived":false,"locations":["us-east-1"],"name":"test","primary":"us-east-1","uuid":"1"`)
		_, err := parseStruct[GroupResp](body)
		if err == nil {
			t.Errorf("Expected error decoding body, got nil")
		}
	})

	t.Run("Testing Group Parsing", func(t *testing.T) {
		body := []byte(`{"archived":false,"locations":["us-east-1"],"name":"test","primary":"us-east-1","uuid":"1"}`)
		response, err := parseStruct[Group](body)
		if err != nil {
			t.Errorf("Error decoding body. %v", err)
		}
		if response.Name != "test" {
			t.Errorf("Expected group name to be test, got %s", response.Name)
		}
	})

	t.Run("Testing Group Parsing with bad body", func(t *testing.T) {
		body := []byte(`{"archived":false,"locations":["us-east-1"],"name":"test","primary":"us-east-1","uuid":"1"`)
		_, err := parseStruct[Group](body)
		if err == nil {
			t.Errorf("Expected error decoding body, got nil")
		}
	})

	t.Run("Testing newAddLocation request generation", func(t *testing.T) {
		req, err := newAddLocationReq("test", "test", "test", "test")
		if err != nil {
			t.Errorf("Error creating request. %v", err)
		}
		if req.Method != "POST" {
			t.Errorf("Expected POST method, got %s", req.Method)
		}
		if req.Header.Get("Authorization") != "Bearer test" {
			t.Errorf("Expected Authorization header to be Bearer test, got %s", req.Header.Get("Authorization"))
		}
		resp, err := (&http.Client{}).Do(req)
		if resp.StatusCode != 401 {
			t.Errorf("Expected 401 status code, got %d", resp.StatusCode)
		}
		if err != nil {
			t.Errorf("Error making request. %v", err)
		}
	})

	t.Run("Testing newCreateGroupReq request generation", func(t *testing.T) {
		req, err := newCreateGroupReq("test", "test", "test", "test")
		if err != nil {
			t.Errorf("Error creating request. %v", err)
		}
		if req.Method != "POST" {
			t.Errorf("Expected POST method, got %s", req.Method)
		}
		if req.Header.Get("Authorization") != "Bearer test" {
			t.Errorf("Expected Authorization header to be Bearer test, got %s", req.Header.Get("Authorization"))
		}
		resp, err := (&http.Client{}).Do(req)
		if resp.StatusCode != 401 {
			t.Errorf("Expected 401 status code, got %d", resp.StatusCode)
		}
		if err != nil {
			t.Errorf("Error making request. %v", err)
		}
	})

	t.Run("Testing newDeleteGroupReq request generation", func(t *testing.T) {
		req, err := newDeleteGroupReq("test", "test", "test")
		if err != nil {
			t.Errorf("Error creating request. %v", err)
		}
		if req.Method != "DELETE" {
			t.Errorf("Expected DELETE method, got %s", req.Method)
		}
		if req.Header.Get("Authorization") != "Bearer test" {
			t.Errorf("Expected Authorization header to be Bearer test, got %s", req.Header.Get("Authorization"))
		}
		resp, err := (&http.Client{}).Do(req)
		if resp.StatusCode != 401 {
			t.Errorf("Expected 401 status code, got %d", resp.StatusCode)
		}
		if err != nil {
			t.Errorf("Error making request. %v", err)
		}
	})

	t.Run("Testing newGetGroupReq request generation", func(t *testing.T) {
		req, err := newGetGroupReq("test", "test", "test")
		if err != nil {
			t.Errorf("Error creating request. %v", err)
		}
		if req.Method != "GET" {
			t.Errorf("Expected GET method, got %s", req.Method)
		}
		if req.Header.Get("Authorization") != "Bearer test" {
			t.Errorf("Expected Authorization header to be Bearer test, got %s", req.Header.Get("Authorization"))
		}
		resp, err := (&http.Client{}).Do(req)
		if resp.StatusCode != 401 {
			t.Errorf("Expected 401 status code, got %d", resp.StatusCode)
		}
		if err != nil {
			t.Errorf("Error making request. %v", err)
		}
	})

	t.Run("Testing newListGroupsReq request generation", func(t *testing.T) {
		req, err := newListGroupsReq("test", "test")
		if err != nil {
			t.Errorf("Error creating request. %v", err)
		}
		if req.Method != "GET" {
			t.Errorf("Expected GET method, got %s", req.Method)
		}
		if req.Header.Get("Authorization") != "Bearer test" {
			t.Errorf("Expected Authorization header to be Bearer test, got %s", req.Header.Get("Authorization"))
		}
		resp, err := (&http.Client{}).Do(req)
		if resp.StatusCode != 401 {
			t.Errorf("Expected 401 status code, got %d", resp.StatusCode)
		}
		if err != nil {
			t.Errorf("Error making request. %v", err)
		}
	})

	t.Run("Testing newDeleteLocationReq request generation", func(t *testing.T) {
		req, err := newDeleteGroupReq("test", "test", "test")
		if err != nil {
			t.Errorf("Error creating request. %v", err)
		}
		if req.Method != "DELETE" {
			t.Errorf("Expected DELETE method, got %s", req.Method)
		}
		if req.Header.Get("Authorization") != "Bearer test" {
			t.Errorf("Expected Authorization header to be Bearer test, got %s", req.Header.Get("Authorization"))
		}
		resp, err := (&http.Client{}).Do(req)
		if resp.StatusCode != 401 {
			t.Errorf("Expected 401 status code, got %d", resp.StatusCode)
		}
		if err != nil {
			t.Errorf("Error making request. %v", err)
		}
	})

	t.Run("Testing newUpdateGroupReq request generation", func(t *testing.T) {
		req, err := newUpdateGroupReq("test", "test", "test")
		if err != nil {
			t.Errorf("Error creating request. %v", err)
		}
		if req.Method != "POST" {
			t.Errorf("Expected PUT method, got %s", req.Method)
		}
		if req.Header.Get("Authorization") != "Bearer test" {
			t.Errorf("Expected Authorization header to be Bearer test, got %s", req.Header.Get("Authorization"))
		}
		resp, err := (&http.Client{}).Do(req)
		if resp.StatusCode != 401 {
			t.Errorf("Expected 401 status code, got %d", resp.StatusCode)
		}
		if err != nil {
			t.Errorf("Error making request. %v", err)
		}
	})
}
