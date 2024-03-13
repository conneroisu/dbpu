package dbpu

import (
	"net/http"
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
	req, err := newCreateDatabaseReq(org.Token, org.Name, db.Name, db.Group)
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

	// make the request and check the response if it is not 401
	resp, err := (&http.Client{}).Do(req)
	if resp.StatusCode != 401 {
		t.Errorf("expected status code 401. Got %d", resp.StatusCode)
	}
	if err != nil {
		t.Errorf("error making request. %v", err)
	}
}
