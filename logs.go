package dbpu

import "fmt"

// GetAuditLogs gets the audit logs for the given organization.
func (c *Client) GetAuditLogs() (AuditLogs, error) {
	req, err := c.newGetAuditLogsRequest()
	if err != nil {
		return AuditLogs{}, fmt.Errorf("failed to create request for audit logs: %v", err)
	}
	done, err := c.Do(req)
	if err != nil {
		return AuditLogs{}, fmt.Errorf("failed to get audit logs: %v", err)
	}
	par, err := parseResponse[AuditLogs](done)
	if err != nil {
		return AuditLogs{}, fmt.Errorf("failed to parse response: %v", err)
	}
	defer done.Body.Close()
	return par, nil
}
