package dbpu

import (
	"bytes"
	"fmt"
	"io"
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
	url := fmt.Sprintf(TursoEndpoint+"/organizations/%s/groups/%s/locations/%s", orgName, groupName, location)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return GroupResponse{}, fmt.Errorf("Error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return GroupResponse{}, fmt.Errorf("Error sending request. %v", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return GroupResponse{}, fmt.Errorf("Error reading body. %v", err)
	}
	response, err := parseStruct[GroupResponse](body)
	if err != nil {
		return GroupResponse{}, fmt.Errorf("Error decoding body. %v", err)
	}
	defer resp.Body.Close()
	return response, nil
}

// ListGroups lists the groups in an organization.
func ListGroups(orgName string, apiToken string) (ListGroupsResponse, error) {
	url := fmt.Sprintf(TursoEndpoint+"/organizations/%s/groups", orgName)
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
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ListGroupsResponse{}, fmt.Errorf("Error reading body. %v", err)
	}
	response, err := parseStruct[ListGroupsResponse](body)
	if err != nil {
		return ListGroupsResponse{}, fmt.Errorf("Error decoding body. %v", err)
	}
	defer resp.Body.Close()
	return response, nil
}

// CreateGroup creates a group in an organization.
func CreateGroup(orgName string, apiToken string, groupName string, location string) (GroupResponse, error) {
	url := fmt.Sprintf(TursoEndpoint+"/organizations/%s/groups", orgName)
	payload := fmt.Sprintf(`{"name": "%s", "location": "%s"}`, groupName, location)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		return GroupResponse{}, fmt.Errorf("Error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	req.Header.Set("Content-Type", "application/json")
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return GroupResponse{}, fmt.Errorf("Error sending request. %v", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return GroupResponse{}, fmt.Errorf("Error reading body. %v", err)
	}
	response, err := parseStruct[GroupResponse](body)
	if err != nil {
		return GroupResponse{}, fmt.Errorf("Error decoding body. %v", err)
	}
	defer resp.Body.Close()
	return response, nil
}

// GetGroup gets a group in an organization.
func GetGroup(orgName string, apiToken string, groupName string) (GroupResponse, error) {
	url := fmt.Sprintf(TursoEndpoint+"/organizations/%s/groups/%s", orgName, groupName)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return GroupResponse{}, fmt.Errorf("Error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return GroupResponse{}, fmt.Errorf("Error sending request. %v", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return GroupResponse{}, fmt.Errorf("Error reading body. %v", err)
	}
	response, err := parseStruct[GroupResponse](body)
	if err != nil {
		return GroupResponse{}, fmt.Errorf("Error decoding body. %v", err)
	}
	defer resp.Body.Close()
	return response, nil
}

// DeleteGroup deletes a group in an organization.
func DeleteGroup(orgName string, apiToken string, groupName string) error {
	url := fmt.Sprintf(TursoEndpoint+"/organizations/%s/groups/%s", orgName, groupName)
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

// TransferGroup transfers a group to a new organization.
func TransferGroup(orgName string, apiToken string, groupName string, newOrgName string) (Group, error) {
	url := fmt.Sprintf(TursoEndpoint+"/organizations/%s/groups/%s/transfer", orgName, groupName)
	payload := fmt.Sprintf(`{"organization": "%s"}`, newOrgName)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		return Group{}, fmt.Errorf("Error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	req.Header.Set("Content-Type", "application/json")
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return Group{}, fmt.Errorf("Error sending request. %v", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Group{}, fmt.Errorf("Error reading body. %v", err)
	}
	response, err := parseStruct[Group](body)
	if err != nil {
		return Group{}, fmt.Errorf("Error decoding body. %v", err)
	}
	defer resp.Body.Close()
	return response, nil
}

// AddLocationToGroup adds a location to a group.
func AddLocationToGroup(orgName string, apiToken string, groupName string, location string) (GroupResponse, error) {
	url := fmt.Sprintf(TursoEndpoint+"/organizations/%s/groups/%s/locations/%s", orgName, groupName, location)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return GroupResponse{}, fmt.Errorf("Error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return GroupResponse{}, fmt.Errorf("Error sending request. %v", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return GroupResponse{}, fmt.Errorf("Error reading body. %v", err)
	}
	response, err := parseStruct[GroupResponse](body)
	if err != nil {
		return GroupResponse{}, fmt.Errorf("Error decoding body. %v", err)
	}
	defer resp.Body.Close()
	return response, nil
}

// RemoveLocationFromGroup removes a location from a group.
func RemoveLocationFromGroup(orgName string, apiToken string, groupName string, location string) (GroupResponse, error) {
	url := fmt.Sprintf(TursoEndpoint+"/organizations/%s/groups/%s/locations/%s", orgName, groupName, location)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return GroupResponse{}, fmt.Errorf("Error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return GroupResponse{}, fmt.Errorf("Error sending request. %v", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return GroupResponse{}, fmt.Errorf("Error reading body. %v", err)
	}
	response, err := parseStruct[GroupResponse](body)
	if err != nil {
		return GroupResponse{}, fmt.Errorf("Error decoding body. %v", err)
	}
	defer resp.Body.Close()
	return response, nil
}

// UpdateVersionGroup updates the version group.
func UpdateVersionGroup(orgName string, apiToken string, groupName string) error {
	url := fmt.Sprintf(TursoEndpoint+"/organizations/%s/groups/%s/update", orgName, groupName)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return fmt.Errorf("Error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return fmt.Errorf("Error sending request. %v", err)
	}
	defer resp.Body.Close()
	return nil
}

// CreateGroupToken creates a token for a group.
func CreateGroupToken(orgName string, apiToken string, groupName string, expiration string, authorization string) (Jwt, error) {
	url := fmt.Sprintf(TursoEndpoint+"/organizations/%s/groups/%s/auth/tokens?expiration=%s&authorization=%s", orgName, groupName, expiration, authorization)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return Jwt{}, fmt.Errorf("Error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return Jwt{}, fmt.Errorf("Error sending request. %v", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Jwt{}, fmt.Errorf("Error reading response body: %v", err)
	}
	jwt, err := parseStruct[Jwt](body)
	if err != nil {
		return Jwt{}, fmt.Errorf("Error decoding response body: %v", err)
	}
	defer resp.Body.Close()
	return jwt, nil
}

// InvalidateGroupTokens invalidates all tokens for a group.
func InvalidateGroupTokens(orgName string, apiToken string, groupName string) error {
	url := fmt.Sprintf("%s/organizations/%s/groups/%s/auth/rotate", TursoEndpoint, orgName, groupName)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return fmt.Errorf("Error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return fmt.Errorf("Error sending request. %v", err)
	}
	defer resp.Body.Close()
	return nil
}
