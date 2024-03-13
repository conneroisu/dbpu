package dbpu

import (
	"fmt"
	"net/http"
)

// GetAuditLogs gets the audit logs for the given organization.
func GetAuditLogs(apiToken string, orgName string) (AuditLogs, error) {
	req, reqErr := newGetAuditLogsRequest(apiToken, orgName)
	resp, doErr := (&http.Client{}).Do(req)
	response, parErr := parseResponse[AuditLogs](resp)
	defer resp.Body.Close()
	return resolveApiCall(response, withReqError(reqErr), withDoError(doErr), withParError(parErr))
}

// newGetAuditLogsRequest creates a request for GetAuditLogs.
func newGetAuditLogsRequest(apiToken string, orgName string) (*http.Request, error) {
	url := fmt.Sprintf(tursoEndpoint+"/organizations/%s/audit-logs", orgName)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	return req, nil
}
