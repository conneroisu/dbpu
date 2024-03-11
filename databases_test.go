package dbpu

import (
	"testing"
)

func TestCreateCreateDatabaseRequest(t *testing.T) {
	org := Org{
		Name:  "org",
		Token: "token",
	}
	db := Db{
		Name:    "name",
		Group:   "group",
		ID:      "id",
		Type:    "type",
		Regions: []string{"region1", "region2"},
	}
	req, err := createCreateDatabaseRequest(org, db)
	if err != nil {
		t.Errorf("Error creating request. %v", err)
	}
	if req.Method != "POST" {
		t.Errorf("Expected POST method. Got %s", req.Method)
	}

	if req.Header.Get("Authorization") != "Bearer token" {
		t.Errorf("Expected Authorization header to be Bearer token. Got %s", req.Header.Get("Authorization"))
	}

	if req.Header.Get("Content-Type") != "application/json" {
		t.Errorf("Expected Content-Type header to be application/json. Got %s", req.Header.Get("Content-Type"))
	}
}
