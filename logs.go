package dbpu

// GetAuditLogs gets the audit logs for the given organization.
func (c *Client) GetAuditLogs(apiToken, orgName string) (AuditLogs, error) {
	req, reqErr := c.newGetAuditLogsRequest(apiToken, orgName)
	done, doErr := c.Do(req)
	response, parErr := parseResponse[AuditLogs](done)
	defer done.Body.Close()
	return resolveApi(response, wReqError(reqErr), wDoError(doErr), wParError(parErr))
}
