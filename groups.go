package dbpu

import (
	"errors"
	"fmt"
	"net/http"
)

// AddLocation adds a location to a group.
func (c *Client) AddLocation(orgName, apiToken, groupName, location string) (GroupResp, error) {
	req, reqErr := c.newAddLocationReq(orgName, apiToken, groupName, location)
	done, doErr := c.Do(req)
	response, respErr := parseResponse[GroupResp](done)
	defer done.Body.Close()
	return resolveApi(response, wReqError(reqErr), wDoError(doErr), wParError(respErr))
}

// ListGroups lists the groups in an organization.
func (c *Client) ListGroups(orgName, apiToken string) (ListGroupsResp, error) {
	req, reqErr := c.newListGroupsReq(orgName, apiToken)
	done, doErr := c.Do(req)
	response, parErr := parseResponse[ListGroupsResp](done)
	defer done.Body.Close()
	return resolveApi(response, wReqError(reqErr), wDoError(doErr), wParError(parErr))
}

// CreateGroup creates a group in an organization.
func (c *Client) CreateGroup(orgName, apiToken, groupName, location string) (GroupResp, error) {
	req, reqErr := c.newCreateGroupReq(orgName, apiToken, groupName, location)
	done, doErr := (&http.Client{}).Do(req)
	response, parErr := parseResponse[GroupResp](done)
	defer done.Body.Close()
	return resolveApi(response, wReqError(reqErr), wDoError(doErr), wParError(parErr))
}

// GetGroup gets a group in an organization.
func (c *Client) GetGroup(orgName, apiToken, groupName string) (GroupResp, error) {
	req, reqErr := c.newGetGroupReq(orgName, apiToken, groupName)
	done, doErr := c.Do(req)
	parsed, parErr := parseResponse[GroupResp](done)
	defer done.Body.Close()
	return resolveApi(parsed, wReqError(reqErr), wDoError(doErr), wParError(parErr))
}

// TransferGroup transfers a group to a new organization.
func (c *Client) TransferGroup(orgName, apiToken, groupName, newOrgName string) (Group, error) {
	req, reqErr := c.newTransferGroupReq(orgName, apiToken, groupName, newOrgName)
	done, doErr := c.Do(req)
	parsed, parErr := parseResponse[Group](done)
	defer done.Body.Close()
	return resolveApi(parsed, wReqError(reqErr), wDoError(doErr), wParError(parErr))
}

// DeleteGroup deletes a group in an organization.
func (c *Client) DeleteGroup(orgName, apiToken, groupName string) (*http.Response, error) {
	req, err := c.newDeleteGroupReq(orgName, apiToken, groupName)
	done, err := c.Do(req)
	done.Body.Close()
	return resolveApi(done, wReqError(err), wDoError(err))
}

// AddLocationToGroup adds a location to a group.
func (c *Client) AddLocationToGroup(orgName, apiToken, groupName, location string) (GroupResp, error) {
	req, reqErr := c.newAddLocationToGroupReq(orgName, apiToken, groupName, location)
	done, doErr := c.Do(req)
	parsed, parErr := parseResponse[GroupResp](done)
	defer done.Body.Close()
	return resolveApi(parsed, wReqError(reqErr), wDoError(doErr), wParError(parErr))
}

// CreateGroupToken creates a token for a group.
func (c *Client) CreateGroupToken(orgName, apiToken, groupName, expiration, authorization string) (Jwt, error) {
	req, reqErr := c.newGroupTokenReq(orgName, apiToken, groupName, expiration, authorization)
	done, doErr := c.Do(req)
	jwt, parErr := parseResponse[Jwt](done)
	defer done.Body.Close()
	return resolveApi(jwt, wReqError(reqErr), wDoError(doErr), wParError(parErr))
}

// RemoveLocationFromGroup removes a location from a group.
func (c *Client) RemoveLocationFromGroup(orgName, apiToken, groupName, location string) (GroupResp, error) {
	req, reqErr := c.newRemoveLocationFromGroupReq(orgName, apiToken, groupName, location)
	done, doErr := c.Do(req)
	parsed, parErr := parseResponse[GroupResp](done)
	defer done.Body.Close()
	return resolveApi(parsed, wReqError(reqErr), wDoError(doErr), wParError(parErr))
}

// InvalidateGroupTokens invalidates all tokens for a group.
func (c *Client) InvalidateGroupTokens(orgName, apiToken, groupName string) error {
	req, reqErr := c.newInvalidateGroupTokensReq(orgName, apiToken, groupName)
	resp, resErr := c.Do(req)
	if errors.Join(reqErr, resErr) != nil {
		return fmt.Errorf("error resolving API. %v", errors.Join(reqErr, resErr))
	}
	defer resp.Body.Close()
	return nil
}

// UpdateVersionGroup updates the version group.
func (c *Client) UpdateVersionGroup(orgName, apiToken, groupName string) (*http.Response, error) {
	req, reqErr := c.newUpdateGroupReq(orgName, apiToken, groupName)
	done, doErr := (&http.Client{}).Do(req)
	done.Body.Close()
	return resolveApi(done, wReqError(reqErr), wDoError(doErr))
}
