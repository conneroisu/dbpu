package dbpu

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
)

// Group is a group in the turso API.
type Group struct {
	Archived      bool     `json:"archived"`
	Locations     []string `json:"locations"`
	Name          string   `json:"name"`
	PrimaryRegion string   `json:"primary"`
	Uuid          string   `json:"uuid"`
}

// GroupResp is a response to adding a location to a group.
type GroupResp struct {
	Group Group `json:"group"`
}

// ListGroupsResp is a response to listing groups.
type ListGroupsResp struct {
	Groups []Group `json:"groups"`
}

// AddLocation adds a location to a group.
func AddLocation(orgName, apiToken, groupName, location string) (GroupResp, error) {
	req, reqErr := newAddLocationReq(orgName, apiToken, groupName, location)
	done, doErr := (&http.Client{}).Do(req)
	response, respErr := parseResponse[GroupResp](done)
	defer done.Body.Close()
	return resolveApiCall(response, wReqError(reqErr), wDoError(doErr), wParError(respErr))
}

// ListGroups lists the groups in an organization.
func ListGroups(orgName, apiToken string) (ListGroupsResp, error) {
	req, reqErr := newListGroupsReq(orgName, apiToken)
	done, doErr := (&http.Client{}).Do(req)
	response, parErr := parseResponse[ListGroupsResp](done)
	defer done.Body.Close()
	return resolveApiCall(response, wReqError(reqErr), wDoError(doErr), wParError(parErr))
}

// CreateGroup creates a group in an organization.
func CreateGroup(orgName, apiToken, groupName, location string) (GroupResp, error) {
	req, reqErr := newCreateGroupReq(orgName, apiToken, groupName, location)
	done, doErr := (&http.Client{}).Do(req)
	response, parErr := parseResponse[GroupResp](done)
	defer done.Body.Close()
	return resolveApiCall(response, wReqError(reqErr), wDoError(doErr), wParError(parErr))
}

// GetGroup gets a group in an organization.
func GetGroup(orgName, apiToken, groupName string) (GroupResp, error) {
	req, err := newGetGroupReq(orgName, apiToken, groupName)
	if err != nil {
		return GroupResp{}, fmt.Errorf("Error creating request. %v", err)
	}
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return GroupResp{}, fmt.Errorf("Error sending request. %v", err)
	}
	response, err := parseResponse[GroupResp](resp)
	if err != nil {
		return GroupResp{}, fmt.Errorf("Error decoding body. %v", err)
	}
	defer resp.Body.Close()
	return response, nil
}

// DeleteGroup deletes a group in an organization.
func DeleteGroup(orgName, apiToken, groupName string) error {
	req, err := newDeleteGroupReq(orgName, apiToken, groupName)
	if err != nil {
		return fmt.Errorf("Error creating request. %v", err)
	}
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return fmt.Errorf("Error sending request. %v", err)
	}
	defer resp.Body.Close()
	return nil
}

// TransferGroup transfers a group to a new organization.
func TransferGroup(orgName, apiToken, groupName, newOrgName string) (Group, error) {
	req, reqErr := newTransferGroupReq(orgName, apiToken, groupName, newOrgName)
	done, doErr := (&http.Client{}).Do(req)
	response, parErr := parseResponse[Group](done)
	defer done.Body.Close()
	return resolveApiCall(response, wReqError(reqErr), wDoError(doErr), wParError(parErr))
}

// AddLocationToGroup adds a location to a group.
func AddLocationToGroup(orgName, apiToken, groupName, location string) (GroupResp, error) {
	req, reqErr := newAddLocationToGroupReq(orgName, apiToken, groupName, location)
	done, doErr := (&http.Client{}).Do(req)
	response, parErr := parseResponse[GroupResp](done)
	defer done.Body.Close()
	return resolveApiCall(response, wReqError(reqErr), wDoError(doErr), wParError(parErr))
}

// CreateGroupToken creates a token for a group.
func CreateGroupToken(orgName, apiToken, groupName, expiration, authorization string) (Jwt, error) {
	req, reqErr := newGroupTokenReq(orgName, apiToken, groupName, expiration, authorization)
	done, doErr := (&http.Client{}).Do(req)
	jwt, parErr := parseResponse[Jwt](done)
	defer done.Body.Close()
	return resolveApiCall(jwt, wReqError(reqErr), wDoError(doErr), wParError(parErr))
}

// RemoveLocationFromGroup removes a location from a group.
func RemoveLocationFromGroup(orgName, apiToken, groupName, location string) (GroupResp, error) {
	req, reqErr := newRemoveLocationFromGroupReq(orgName, apiToken, groupName, location)
	done, doErr := (&http.Client{}).Do(req)
	response, parErr := parseResponse[GroupResp](done)
	defer done.Body.Close()
	return resolveApiCall(response, wReqError(reqErr), wDoError(doErr), wParError(parErr))
}

// InvalidateGroupTokens invalidates all tokens for a group.
func InvalidateGroupTokens(orgName, apiToken, groupName string) error {
	req, reqErr := newInvalidateGroupTokensRequest(orgName, apiToken, groupName)
	resp, resErr := (&http.Client{}).Do(req)
	if errors.Join(reqErr, resErr) != nil {
		return fmt.Errorf("error resolving API. %v", errors.Join(reqErr, resErr))
	}
	defer resp.Body.Close()
	return nil
}

// UpdateVersionGroup updates the version group.
func UpdateVersionGroup(orgName, apiToken, groupName string) (*http.Response, error) {
	req, reqErr := newUpdateGroupReq(orgName, apiToken, groupName)
	done, doErr := (&http.Client{}).Do(req)
	done.Body.Close()
	return resolveApiCall(done, wReqError(reqErr), wDoError(doErr))
}

// newRemoveLocationFromGroupReq creates a request for removing a location from a group.
func newRemoveLocationFromGroupReq(orgName, apiToken, groupName, location string) (*http.Request, error) {
	url := fmt.Sprintf(tursoEndpoint+"/organizations/%s/groups/%s/locations/%s", orgName, groupName, location)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	return req, nil
}

// newInvalidateGroupTokensRequest creates a request for invalidating all tokens for a group.
func newInvalidateGroupTokensRequest(orgName, apiToken, groupName string) (*http.Request, error) {
	url := fmt.Sprintf(tursoEndpoint+"/organizations/%s/groups/%s/auth/rotate", orgName, groupName)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	return req, nil
}

// newCreateGroupReq creates a request for creating a group in an organization.
func newCreateGroupReq(orgName, apiToken, groupName, location string) (*http.Request, error) {
	url := fmt.Sprintf(tursoEndpoint+"/organizations/%s/groups", orgName)
	payload := fmt.Sprintf(`{"name": "%s", "location": "%s"}`, groupName, location)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		return nil, fmt.Errorf("error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

// newGroupTokenReq creates a request for creating a token for a group.
func newGroupTokenReq(orgName, apiToken, groupName, expiration, authorization string) (*http.Request, error) {
	url := fmt.Sprintf(tursoEndpoint+"/organizations/%s/groups/%s/auth/tokens?expiration=%s&authorization=%s", orgName, groupName, expiration, authorization)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	return req, nil
}

// newUpdateGroupReq creates a request for updating a group.
func newUpdateGroupReq(orgName, apiToken, groupName string) (*http.Request, error) {
	url := fmt.Sprintf(tursoEndpoint+"/organizations/%s/groups/%s/update", orgName, groupName)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	return req, nil
}

// newGetGroupReq creates a request for getting a group.
func newGetGroupReq(orgName, apiToken, groupName string) (*http.Request, error) {
	url := fmt.Sprintf("%s/organizations/%s/groups/%s",
		tursoEndpoint, orgName, groupName,
	)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	return req, nil
}

// newAddLocationReq creates a request for adding a location to a group.
func newAddLocationReq(orgName, apiToken, groupName, location string) (*http.Request, error) {
	url := fmt.Sprintf(tursoEndpoint+"/organizations/%s/groups/%s/locations/%s", orgName, groupName, location)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	return req, nil
}

// newListGroupsReq creates a request for listing groups.
func newListGroupsReq(orgName, apiToken string) (*http.Request, error) {
	url := fmt.Sprintf(tursoEndpoint+"/organizations/%s/groups", orgName)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	return req, nil
}

// newDeleteGroupReq creates a request for deleting a group.
func newDeleteGroupReq(orgName, apiToken, groupName string) (*http.Request, error) {
	url := fmt.Sprintf(tursoEndpoint+"/organizations/%s/groups/%s", orgName, groupName)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	return req, nil
}

// newTransferGroupReq creates a request for transferring a group to a new organization.
func newTransferGroupReq(orgName, apiToken, groupName, newOrgName string) (*http.Request, error) {
	url := fmt.Sprintf("%s/organizations/%s/groups/%s/transfer", tursoEndpoint, orgName, groupName)
	payload := fmt.Sprintf(`{"organization": "%s"}`, newOrgName)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		return nil, fmt.Errorf("error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

// newAddLocationToGroupReq creates a request for adding a location to a group.
func newAddLocationToGroupReq(orgName, apiToken, groupName, location string) (*http.Request, error) {
	url := fmt.Sprintf("%s/organizations/%s/groups/%s/locations/%s", tursoEndpoint, orgName, groupName, location)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	return req, nil
}
