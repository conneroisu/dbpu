package dbpu

import (
	"fmt"
	"io"
	"net/http"
)

// CreateGetAuditLogsRequest creates a request for GetAuditLogs.
func CreateGetAuditLogsRequest(apiToken string, orgName string) (*http.Request, error) {
	url := fmt.Sprintf(TursoEndpoint+"/organizations/%s/audit-logs", orgName)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	return req, nil
}

// GetAuditLogs gets the audit logs for the given organization.
func GetAuditLogs(apiToken string, orgName string) (AuditLogs, error) {
	req, err := CreateGetAuditLogsRequest(apiToken, orgName)
	if err != nil {
		return AuditLogs{}, fmt.Errorf("Error reading request. %v", err)
	}
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return AuditLogs{}, fmt.Errorf("Error sending request. %v", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return AuditLogs{}, fmt.Errorf("Error reading body. %v", err)
	}
	response, err := parseStruct[AuditLogs](body)
	if err != nil {
		return AuditLogs{}, fmt.Errorf("Error decoding body. %v", err)
	}
	defer resp.Body.Close()
	return response, nil
}
