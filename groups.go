package dbpu

import (
	"bytes"
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
func AddLocation(orgName string, apiToken string, groupName string, location string) (GroupResp, error) {
	req, err := newAddLocationReq(orgName, apiToken, groupName, location)
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

// ListGroups lists the groups in an organization.
func ListGroups(orgName string, apiToken string) (ListGroupsResp, error) {
	req, err := newListGroupsReq(orgName, apiToken)
	if err != nil {
		return ListGroupsResp{}, fmt.Errorf("Error creating request. %v", err)
	}
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return ListGroupsResp{}, fmt.Errorf("Error sending request. %v", err)
	}
	response, err := parseResponse[ListGroupsResp](resp)
	if err != nil {
		return ListGroupsResp{}, fmt.Errorf("Error decoding body. %v", err)
	}
	defer resp.Body.Close()
	return response, nil
}

// CreateGroup creates a group in an organization.
func CreateGroup(orgName string, apiToken string, groupName string, location string) (GroupResp, error) {
	req, err := newCreateGroupReq(orgName, apiToken, groupName, location)
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

// GetGroup gets a group in an organization.
func GetGroup(orgName string, apiToken string, groupName string) (GroupResp, error) {
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
func DeleteGroup(orgName string, apiToken string, groupName string) error {
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
func TransferGroup(orgName string, apiToken string, groupName string, newOrgName string) (Group, error) {
	req, err := newTransferGroupReq(orgName, apiToken, groupName, newOrgName)
	if err != nil {
		return Group{}, fmt.Errorf("Error creating request. %v", err)
	}
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return Group{}, fmt.Errorf("Error sending request. %v", err)
	}
	response, err := parseResponse[Group](resp)
	if err != nil {
		return Group{}, fmt.Errorf("Error decoding body. %v", err)
	}
	defer resp.Body.Close()
	return response, nil
}

// AddLocationToGroup adds a location to a group.
func AddLocationToGroup(orgName string, apiToken string, groupName string, location string) (GroupResp, error) {
	req, err := newAddLocationToGroupReq(orgName, apiToken, groupName, location)
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

// CreateGroupToken creates a token for a group.
func CreateGroupToken(orgName string, apiToken string, groupName string, expiration string, authorization string) (Jwt, error) {
	req, err := newGroupTokenReq(orgName, apiToken, groupName, expiration, authorization)
	if err != nil {
		return Jwt{}, fmt.Errorf("Error creating request. %v", err)
	}
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return Jwt{}, fmt.Errorf("Error sending request. %v", err)
	}
	jwt, err := parseResponse[Jwt](resp)
	if err != nil {
		return Jwt{}, fmt.Errorf("Error decoding response body: %v", err)
	}
	defer resp.Body.Close()
	return jwt, nil
}

// RemoveLocationFromGroup removes a location from a group.
func RemoveLocationFromGroup(orgName string, apiToken string, groupName string, location string) (GroupResp, error) {
	req, err := newRemoveLocationFromGroupReq(orgName, apiToken, groupName, location)
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

// UpdateVersionGroup updates the version group.
func UpdateVersionGroup(orgName string, apiToken string, groupName string) error {
	req, err := newUpdateGroupReq(orgName, apiToken, groupName)
	if err != nil {
		return fmt.Errorf("Error creating request. %v", err)
	}
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return fmt.Errorf("error sending request. %v", err)
	}
	defer resp.Body.Close()
	return nil
}

// newRemoveLocationFromGroupReq creates a request for removing a location from a group.
func newRemoveLocationFromGroupReq(orgName string, apiToken string, groupName string, location string) (*http.Request, error) {
	url := fmt.Sprintf(tursoEndpoint+"/organizations/%s/groups/%s/locations/%s", orgName, groupName, location)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	return req, nil
}

// createInvalidateGroupTokensRequest creates a request for invalidating all tokens for a group.
func createInvalidateGroupTokensRequest(orgName string, apiToken string, groupName string) (*http.Request, error) {
	url := fmt.Sprintf(tursoEndpoint+"/organizations/%s/groups/%s/auth/rotate", orgName, groupName)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	return req, nil
}

// InvalidateGroupTokens invalidates all tokens for a group.
func InvalidateGroupTokens(orgName string, apiToken string, groupName string) error {
	req, err := createInvalidateGroupTokensRequest(orgName, apiToken, groupName)
	if err != nil {
		return fmt.Errorf("Error reading request. %v", err)
	}
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return fmt.Errorf("Error sending request. %v", err)
	}
	defer resp.Body.Close()
	return nil
}

// newCreateGroupReq creates a request for creating a group in an organization.
func newCreateGroupReq(orgName string, apiToken string, groupName string, location string) (*http.Request, error) {
	url := fmt.Sprintf(tursoEndpoint+"/organizations/%s/groups", orgName)
	payload := fmt.Sprintf(`{"name": "%s", "location": "%s"}`, groupName, location)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		return nil, fmt.Errorf("Error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

// newGroupTokenReq creates a request for creating a token for a group.
func newGroupTokenReq(orgName string, apiToken string, groupName string, expiration string, authorization string) (*http.Request, error) {
	url := fmt.Sprintf(tursoEndpoint+"/organizations/%s/groups/%s/auth/tokens?expiration=%s&authorization=%s", orgName, groupName, expiration, authorization)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	return req, nil
}

// newUpdateGroupReq creates a request for updating a group.
func newUpdateGroupReq(orgName string, apiToken string, groupName string) (*http.Request, error) {
	url := fmt.Sprintf(tursoEndpoint+"/organizations/%s/groups/%s/update", orgName, groupName)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	return req, nil
}

// newGetGroupReq creates a request for getting a group.
func newGetGroupReq(orgName string, apiToken string, groupName string) (*http.Request, error) {
	url := fmt.Sprintf("%s/organizations/%s/groups/%s",
		tursoEndpoint, orgName, groupName,
	)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	return req, nil
}

// newAddLocationReq creates a request for adding a location to a group.
func newAddLocationReq(orgName string, apiToken string, groupName string, location string) (*http.Request, error) {
	url := fmt.Sprintf(tursoEndpoint+"/organizations/%s/groups/%s/locations/%s", orgName, groupName, location)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	return req, nil
}

// newListGroupsReq creates a request for listing groups.
func newListGroupsReq(orgName string, apiToken string) (*http.Request, error) {
	url := fmt.Sprintf(tursoEndpoint+"/organizations/%s/groups", orgName)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	return req, nil
}

// newDeleteGroupReq creates a request for deleting a group.
func newDeleteGroupReq(orgName string, apiToken string, groupName string) (*http.Request, error) {
	url := fmt.Sprintf(tursoEndpoint+"/organizations/%s/groups/%s", orgName, groupName)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	return req, nil
}

// newTransferGroupReq creates a request for transferring a group to a new organization.
func newTransferGroupReq(orgName string, apiToken string, groupName string, newOrgName string) (*http.Request, error) {
	url := fmt.Sprintf(tursoEndpoint+"/organizations/%s/groups/%s/transfer", orgName, groupName)
	payload := fmt.Sprintf(`{"organization": "%s"}`, newOrgName)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		return nil, fmt.Errorf("Error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

// newAddLocationToGroupReq creates a request for adding a location to a group.
func newAddLocationToGroupReq(orgName string, apiToken string, groupName string, location string) (*http.Request, error) {
	url := fmt.Sprintf(tursoEndpoint+"/organizations/%s/groups/%s/locations/%s", orgName, groupName, location)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	return req, nil
}
