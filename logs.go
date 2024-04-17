package dbpu

import (
	"fmt"
	"net/http"
)

// GetAuditLogs gets the audit logs for the given organization.
//
// It requires a valid base url, organization name, and api token.
//
// It returns a pointer to the audit logs object.
func (c *Client) GetAuditLogs() (*AuditLogs, error) {
	req, err := c.newGetAuditLogsRequest()
	if err != nil {
		return nil, fmt.Errorf("failed to create request for audit logs: %v", err)
	}
	done, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get audit logs: %v", err)
	}
	par, err := parseResponse[AuditLogs](done)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}
	defer done.Body.Close()
	return par, nil
}

// newGetAuditLogsRequest creates a request for GetAuditLogs.
func (c *Client) newGetAuditLogsRequest() (*http.Request, error) {
	url := fmt.Sprintf("%s/organizations/%s/audit-logs", c.BaseURL, c.OrgName)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.ApiToken))
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}
