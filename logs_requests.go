package dbpu

import (
	"fmt"
	"net/http"
)

// newGetAuditLogsRequest creates a request for GetAuditLogs.
func (c *Client) newGetAuditLogsRequest(apiToken, orgName string) (*http.Request, error) {
	url := fmt.Sprintf("%s/organizations/%s/audit-logs", c.BaseURL, orgName)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}
