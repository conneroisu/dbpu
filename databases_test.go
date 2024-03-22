package dbpu

import (
	"net/http"
	"testing"
)

func testParsing(t *testing.T) {
	t.Run("Test Parse Database", func(t *testing.T) {
		body := []byte(`
			{
				"id":"test",
				"name":"test",
				"group":"test",
				"type":"test",
				"regions":["test"]
			}
		`)
		db, err := parseStruct[Database](body)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if db.ID != "test" {
			t.Errorf("Expected ID to be test, got %v", db.ID)
		}
		if db.Name != "test" {
			t.Errorf("Expected Name to be test, got %v", db.Name)
		}
		if db.Group != "test" {
			t.Errorf("Expected Group to be test, got %v", db.Group)
		}
		if db.Type != "test" {
			t.Errorf("Expected Type to be test, got %v", db.Type)
		}
		if db.Regions[0] != "test" {
			t.Errorf("Expected Regions[0] to be test, got %v", db.Regions[0])
		}
	})
}

func testCreateCreateDatabaseRequest(t *testing.T) {
	org := Org{
		Name:  "org",
		Token: "token",
	}
	db := Database{
		Name:    "name",
		Group:   "group",
		ID:      "id",
		Type:    "type",
		Regions: []string{"region1", "region2"},
	}
	c := NewClient()
	req, err := c.newCreateDatabaseReq(org.Token, org.Name, db.Name, db.Group)
	if err != nil {
		t.Errorf("error creating request. %v", err)
	}
	if req.Method != "POST" {
		t.Errorf("expected POST method. Got %s", req.Method)
	}
	if req.Header.Get("Authorization") != "Bearer token" {
		t.Errorf("expected Authorization header to be Bearer token. Got %s", req.Header.Get("Authorization"))
	}
	if req.Header.Get("Content-Type") != "application/json" {
		t.Errorf("expected Content-Type header to be application/json. Got %s", req.Header.Get("Content-Type"))
	}
	resp, err := (&http.Client{}).Do(req)
	if resp.StatusCode != 401 {
		t.Errorf("expected 401 status code. Got %d", resp.StatusCode)
	}
	if err != nil {
		t.Errorf("error making request. %v", err)
	}
}
