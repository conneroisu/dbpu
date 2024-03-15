package dbpu

import (
	"fmt"
	"net/http"
)

// GetAuditLogs gets the audit logs for the given organization.
func (c *Client) GetAuditLogs(apiToken, orgName string) (AuditLogs, error) {
	req, reqErr := newGetAuditLogsRequest(apiToken, orgName)
	done, doErr := c.Do(req)
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
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}
