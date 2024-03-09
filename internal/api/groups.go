package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/charmbracelet/log"
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

// parseGroupResponse parses the response from the server into an GroupResponse.
func parseGroupResponse(body io.Reader, err error) (GroupResponse, error) {
	var response GroupResponse
	err = json.NewDecoder(body).Decode(&response)
	return response, err
}

// parseGroup parses the response from the server into a Group.
func parseGroup(body io.Reader, err error) (Group, error) {
	var response Group
	err = json.NewDecoder(body).Decode(&response)
	return response, err
}

// ListGroupsResponse is a response to listing groups.
type ListGroupsResponse struct {
	Groups []Group `json:"groups"`
}

// parseListGroupsResponse parses the response from the server into a ListGroupsResponse.
func parseListGroupsResponse(body io.Reader, err error) (ListGroupsResponse, error) {
	var response ListGroupsResponse
	err = json.NewDecoder(body).Decode(&response)
	return response, err
}

// AddLocation adds a location to a group.
func AddLocation(orgName string, apiToken string, groupName string, location string) (GroupResponse, error) {
	url := fmt.Sprintf(TursoEndpoint+"/organizations/%s/groups/%s/locations/%s", orgName, groupName, location)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Errorf("Error creating request. %v", err)
		return GroupResponse{}, fmt.Errorf("Error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("Error sending request. %v", err)
		return GroupResponse{}, fmt.Errorf("Error sending request. %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("Error reading body. %v", err)
		return GroupResponse{}, fmt.Errorf("Error reading body. %v", err)
	}
	response, err := parseGroupResponse(bytes.NewReader(body), err)
	if err != nil {
		log.Errorf("Error decoding body into GroupResponse. %v", err)
		return GroupResponse{}, fmt.Errorf("Error decoding body. %v", err)
	}
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
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ListGroupsResponse{}, fmt.Errorf("Error reading body. %v", err)
	}
	response, err := parseListGroupsResponse(bytes.NewReader(body), err)
	if err != nil {
		return ListGroupsResponse{}, fmt.Errorf("Error decoding body. %v", err)
	}
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
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return GroupResponse{}, fmt.Errorf("Error sending request. %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return GroupResponse{}, fmt.Errorf("Error reading body. %v", err)
	}
	response, err := parseGroupResponse(bytes.NewReader(body), err)
	if err != nil {
		return GroupResponse{}, fmt.Errorf("Error decoding body. %v", err)
	}
	return response, nil
}

// curl -L 'https://api.turso.tech/v1/organizations/{organizationName}/groups/{groupName}' \
// -H 'Authorization: Bearer TOKEN'
func GetGroup(orgName string, apiToken string, groupName string) (GroupResponse, error) {
	url := fmt.Sprintf(TursoEndpoint+"/organizations/%s/groups/%s", orgName, groupName)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return GroupResponse{}, fmt.Errorf("Error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return GroupResponse{}, fmt.Errorf("Error sending request. %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return GroupResponse{}, fmt.Errorf("Error reading body. %v", err)
	}
	response, err := parseGroupResponse(bytes.NewReader(body), err)
	if err != nil {
		return GroupResponse{}, fmt.Errorf("Error decoding body. %v", err)
	}
	return response, nil
}

// curl -L -X DELETE 'https://api.turso.tech/v1/organizations/{organizationName}/groups/{groupName}' \
// -H 'Authorization: Bearer TOKEN'
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
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return Group{}, fmt.Errorf("Error sending request. %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Group{}, fmt.Errorf("Error reading body. %v", err)
	}
	response, err := parseGroup(bytes.NewReader(body), err)
	if err != nil {
		return Group{}, fmt.Errorf("Error decoding body. %v", err)
	}
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
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return GroupResponse{}, fmt.Errorf("Error sending request. %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return GroupResponse{}, fmt.Errorf("Error reading body. %v", err)
	}
	response, err := parseGroupResponse(bytes.NewReader(body), err)
	if err != nil {
		return GroupResponse{}, fmt.Errorf("Error decoding body. %v", err)
	}
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
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return GroupResponse{}, fmt.Errorf("Error sending request. %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return GroupResponse{}, fmt.Errorf("Error reading body. %v", err)
	}
	response, err := parseGroupResponse(bytes.NewReader(body), err)
	if err != nil {
		return GroupResponse{}, fmt.Errorf("Error decoding body. %v", err)
	}
	return response, nil
}

// curl -L -X POST 'https://api.turso.tech/v1/organizations/{organizationName}/groups/{groupName}/update' \
// -H 'Authorization: Bearer TOKEN'
func UpdateVersionGroup(orgName string, apiToken string, groupName string) error {
	url := fmt.Sprintf(TursoEndpoint+"/organizations/%s/groups/%s/update", orgName, groupName)
	req, err := http.NewRequest("POST", url, nil)
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

// curl -L -X POST 'https://api.turso.tech/v1/organizations/{organizationName}/groups/{groupName}/auth/tokens?expiration=2w&authorization=full-access' \
// -H 'Authorization: Bearer TOKEN'
func CreateGroupToken(orgName string, apiToken string, groupName string, expiration string, authorization string) (Jwt, error) {
	url := fmt.Sprintf(TursoEndpoint+"/organizations/%s/groups/%s/auth/tokens?expiration=%s&authorization=%s", orgName, groupName, expiration, authorization)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return Jwt{}, fmt.Errorf("Error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return Jwt{}, fmt.Errorf("Error sending request. %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return Jwt{}, fmt.Errorf("Error creating group token: %s", resp.Status)
	}
	var jwt Jwt
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Jwt{}, fmt.Errorf("Error reading response body: %v", err)
	}
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&jwt)
	if err != nil {
		return Jwt{}, fmt.Errorf("Error decoding response body: %v", err)
	}
	return jwt, nil
}

// InvalidateGroupTokens invalidates all tokens for a group.
func InvalidateGroupTokens(orgName string, apiToken string, groupName string) error {
	url := fmt.Sprintf(TursoEndpoint+"/organizations/%s/groups/%s/auth/rotate", orgName, groupName)
	req, err := http.NewRequest("POST", url, nil)
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
