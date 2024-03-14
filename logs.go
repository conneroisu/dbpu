package dbpu

import (
	"fmt"
	"net/http"
)

// GetAuditLogs gets the audit logs for the given organization.
func GetAuditLogs(apiToken, orgName string) (AuditLogs, error) {
	req, reqErr := newGetAuditLogsRequest(apiToken, orgName)
	done, doErr := (&http.Client{}).Do(req)
	response, parErr := parseResponse[AuditLogs](done)
	defer done.Body.Close()
	return resolveApiCall(response, wReqError(reqErr), wDoError(doErr), wParError(parErr))
}

// newGetAuditLogsRequest creates a request for GetAuditLogs.
func newGetAuditLogsRequest(apiToken, orgName string) (*http.Request, error) {
	url := fmt.Sprintf(tursoEndpoint+"/organizations/%s/audit-logs", orgName)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	return req, nil
}
