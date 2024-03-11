package dbpu

import (
	"fmt"
	"io"
	"net/http"
)

// GetAuditLogs gets the audit logs for the given organization.
func GetAuditLogs(apiToken string, orgName string) (AuditLogs, error) {
	url := fmt.Sprintf(TursoEndpoint+"/organizations/%s/audit-logs", orgName)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return AuditLogs{}, fmt.Errorf("Error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return AuditLogs{}, fmt.Errorf("Error sending request. %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return AuditLogs{}, fmt.Errorf("Error reading body. %v", err)
	}
	response, err := parseStruct[AuditLogs](body)
	if err != nil {
		return AuditLogs{}, fmt.Errorf("Error decoding body. %v", err)
	}
	return response, nil
}
