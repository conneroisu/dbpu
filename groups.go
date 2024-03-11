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

// GroupResponse is a response to adding a location to a group.
type GroupResponse struct {
	Group Group `json:"group"`
}

// ListGroupsResponse is a response to listing groups.
type ListGroupsResponse struct {
	Groups []Group `json:"groups"`
}

// AddLocation adds a location to a group.
func AddLocation(orgName string, apiToken string, groupName string, location string) (GroupResponse, error) {
	url := fmt.Sprintf(tursoEndpoint+"/organizations/%s/groups/%s/locations/%s", orgName, groupName, location)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return GroupResponse{}, fmt.Errorf("Error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return GroupResponse{}, fmt.Errorf("Error sending request. %v", err)
	}
	response, err := parseResponse[GroupResponse](resp)
	if err != nil {
		return GroupResponse{}, fmt.Errorf("Error decoding body. %v", err)
	}
	defer resp.Body.Close()
	return response, nil
}

// ListGroups lists the groups in an organization.
func ListGroups(orgName string, apiToken string) (ListGroupsResponse, error) {
	url := fmt.Sprintf(tursoEndpoint+"/organizations/%s/groups", orgName)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ListGroupsResponse{}, fmt.Errorf("Error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return ListGroupsResponse{}, fmt.Errorf("Error sending request. %v", err)
	}
	response, err := parseResponse[ListGroupsResponse](resp)
	if err != nil {
		return ListGroupsResponse{}, fmt.Errorf("Error decoding body. %v", err)
	}
	defer resp.Body.Close()
	return response, nil
}

// createCreateGroupRequest creates a request for creating a group in an organization.
func createCreateGroupRequest(orgName string, apiToken string, groupName string, location string) (*http.Request, error) {
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

// CreateGroup creates a group in an organization.
func CreateGroup(orgName string, apiToken string, groupName string, location string) (GroupResponse, error) {
	req, err := createCreateGroupRequest(orgName, apiToken, groupName, location)
	if err != nil {
		return GroupResponse{}, fmt.Errorf("Error creating request. %v", err)
	}
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return GroupResponse{}, fmt.Errorf("Error sending request. %v", err)
	}
	response, err := parseResponse[GroupResponse](resp)
	if err != nil {
		return GroupResponse{}, fmt.Errorf("Error decoding body. %v", err)
	}
	defer resp.Body.Close()
	return response, nil
}

func createGetGroupRequest(orgName string, apiToken string, groupName string) (*http.Request, error) {
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

// GetGroup gets a group in an organization.
func GetGroup(orgName string, apiToken string, groupName string) (GroupResponse, error) {
	req, err := createGetGroupRequest(orgName, apiToken, groupName)
	if err != nil {
		return GroupResponse{}, fmt.Errorf("Error creating request. %v", err)
	}
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return GroupResponse{}, fmt.Errorf("Error sending request. %v", err)
	}
	response, err := parseResponse[GroupResponse](resp)
	if err != nil {
		return GroupResponse{}, fmt.Errorf("Error decoding body. %v", err)
	}
	defer resp.Body.Close()
	return response, nil
}

// DeleteGroup deletes a group in an organization.
func DeleteGroup(orgName string, apiToken string, groupName string) error {
	url := fmt.Sprintf(tursoEndpoint+"/organizations/%s/groups/%s", orgName, groupName)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("Error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Error sending request. %v", err)
	}
	defer resp.Body.Close()
	return nil
}

func createTransferGroupRequest(orgName string, apiToken string, groupName string, newOrgName string) (*http.Request, error) {
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

// TransferGroup transfers a group to a new organization.
func TransferGroup(orgName string, apiToken string, groupName string, newOrgName string) (Group, error) {
	req, err := createTransferGroupRequest(orgName, apiToken, groupName, newOrgName)
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
func AddLocationToGroup(orgName string, apiToken string, groupName string, location string) (GroupResponse, error) {
	url := fmt.Sprintf(tursoEndpoint+"/organizations/%s/groups/%s/locations/%s", orgName, groupName, location)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return GroupResponse{}, fmt.Errorf("Error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return GroupResponse{}, fmt.Errorf("Error sending request. %v", err)
	}
	response, err := parseResponse[GroupResponse](resp)
	if err != nil {
		return GroupResponse{}, fmt.Errorf("Error decoding body. %v", err)
	}
	defer resp.Body.Close()
	return response, nil
}

// createRemoveLocationFromGroupRequest creates a request for removing a location from a group.
func createRemoveLocationFromGroupRequest(orgName string, apiToken string, groupName string, location string) (*http.Request, error) {
	url := fmt.Sprintf(tursoEndpoint+"/organizations/%s/groups/%s/locations/%s", orgName, groupName, location)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	return req, nil
}

// RemoveLocationFromGroup removes a location from a group.
func RemoveLocationFromGroup(orgName string, apiToken string, groupName string, location string) (GroupResponse, error) {
	req, err := createRemoveLocationFromGroupRequest(orgName, apiToken, groupName, location)
	if err != nil {
		return GroupResponse{}, fmt.Errorf("Error creating request. %v", err)
	}
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return GroupResponse{}, fmt.Errorf("Error sending request. %v", err)
	}
	response, err := parseResponse[GroupResponse](resp)
	if err != nil {
		return GroupResponse{}, fmt.Errorf("Error decoding body. %v", err)
	}
	defer resp.Body.Close()
	return response, nil
}

// UpdateVersionGroup updates the version group.
func UpdateVersionGroup(orgName string, apiToken string, groupName string) error {
	url := fmt.Sprintf(tursoEndpoint+"/organizations/%s/groups/%s/update", orgName, groupName)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return fmt.Errorf("error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return fmt.Errorf("error sending request. %v", err)
	}
	defer resp.Body.Close()
	return nil
}

// createGroupTokenRequest creates a request for creating a token for a group.
func createGroupTokenRequest(orgName string, apiToken string, groupName string, expiration string, authorization string) (*http.Request, error) {
	url := fmt.Sprintf(tursoEndpoint+"/organizations/%s/groups/%s/auth/tokens?expiration=%s&authorization=%s", orgName, groupName, expiration, authorization)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	return req, nil
}

// CreateGroupToken creates a token for a group.
func CreateGroupToken(orgName string, apiToken string, groupName string, expiration string, authorization string) (Jwt, error) {
	req, err := createGroupTokenRequest(orgName, apiToken, groupName, expiration, authorization)
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
