package dbpu

import (
	"fmt"
	"testing"
)

// TestCreateGetAuditLogsRequest tests the create get audit logs request.
func testCreateGetAuditLogsRequest(t *testing.T) {
	// Test the create get audit logs request
	t.Run("Test Create Get Audit Logs Request", func(t *testing.T) {
		c := NewClient()
		apiToken := "test"
		orgName := "test"
		req, err := c.newGetAuditLogsRequest(apiToken, orgName)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if req.Method != "GET" {
			t.Errorf("Expected method to be GET, got %v", req.Method)
		}
		if req.URL.String() != fmt.Sprintf(tursoEndpoint+"/organizations/%s/audit-logs", orgName) {
			t.Errorf("Expected URL to be %s, got %v", fmt.Sprintf(tursoEndpoint+"/organizations/%s/audit-logs", orgName), req.URL.String())
		}
		if req.Header.Get("Authorization") != fmt.Sprintf("Bearer %s", apiToken) {
			t.Errorf("Expected Authorization to be %s, got %v", fmt.Sprintf("Bearer %s", apiToken), req.Header.Get("Authorization"))
		}
	})
}

// TestParseAuditLogs tests the parse audit logs.
func testLogsParse(t *testing.T) {
	// Test the parse audit logs
	t.Run("Test Parse Audit Logs", func(t *testing.T) {
		body := []byte(`
		{
			"audit_logs":[ 
				{
					"author":"test",
					"code":"test",
					"created_at":"test",
					"data":"test",
					"message":"test",
					"origin":"test"
				}
			],
		"pagination":{"page":1,"page_size":1,"total_pages":1,"total_rows":1}}`)

		auditLogs, err := parseStruct[AuditLogs](body)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(auditLogs.AuditLogs) != 1 {
			t.Errorf("Expected AuditLogs to have 1 element, got %v", len(auditLogs.AuditLogs))
		}

		if auditLogs.AuditLogs[0].Author != "test" {
			t.Errorf("Expected Author to be test, got %v", auditLogs.AuditLogs[0].Author)
		}

		if auditLogs.AuditLogs[0].Code != "test" {
			t.Errorf("Expected Code to be test, got %v", auditLogs.AuditLogs[0].Code)
		}

		if auditLogs.AuditLogs[0].CreatedAt != "test" {
			t.Errorf("Expected CreatedAt to be test, got %v", auditLogs.AuditLogs[0].CreatedAt)
		}

		if auditLogs.AuditLogs[0].Data != "test" {
			t.Errorf("Expected Data to be test, got %v", auditLogs.AuditLogs[0].Data)
		}

		if auditLogs.AuditLogs[0].Message != "test" {
			t.Errorf("Expected Message to be test, got %v", auditLogs.AuditLogs[0].Message)
		}
	})
}
