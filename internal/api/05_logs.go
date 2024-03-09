package api

import (
	"fmt"
	"io"
	"net/http"

	"github.com/bytedance/sonic/decoder"
)

// AuditLogs is a response to listing organizations.
type AuditLogs struct {
	AuditLogs  []AuditLog `json:"audit_logs"`
	Pagination Pagination `json:"pagination"`
}

// AuditLog is a response to listing organizations.
type AuditLog struct {
	Author    string `json:"author"`
	Code      string `json:"code"`
	CreatedAt string `json:"created_at"`
	Data      string `json:"data"`
	Message   string `json:"message"`
	Origin    string `json:"origin"`
}

// Pagination is a response to listing organizations.
type Pagination struct {
	Page       int `json:"page"`
	PageSize   int `json:"page_size"`
	TotalPages int `json:"total_pages"`
	TotalRows  int `json:"total_rows"`
}

// GetAuditLogs gets the audit logs for the given organization.
func GetAuditLogs(apiToken string, orgName string) (AuditLogs, error) {
	url := fmt.Sprintf(TursoEndpoint+"/organizations/%s/audit-logs", orgName)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return AuditLogs{}, fmt.Errorf("Error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return AuditLogs{}, fmt.Errorf("Error sending request. %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return AuditLogs{}, fmt.Errorf("Error reading body. %v", err)
	}
	auditLogs, err := parseAuditLogs(body, err)
	if err != nil {
		return AuditLogs{}, fmt.Errorf("Error decoding body. %v", err)
	}
	return auditLogs, nil
}

// parseAuditLogs parses the response from the server into a AuditLogs.
func parseAuditLogs(body []byte, err error) (AuditLogs, error) {
	var response AuditLogs
	err = decoder.NewDecoder(string(body)).Decode(&response)
	return response, err
}
