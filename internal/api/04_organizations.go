package api

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/bytedance/sonic/decoder"
	"github.com/charmbracelet/log"
)

// Organization is a response to listing organizations.
type Organization struct {
	BlockedReads  bool   `json:"blocked_reads"`
	BlockedWrites bool   `json:"blocked_writes"`
	Name          string `json:"name"`
	Overages      bool   `json:"overages"`
	Slug          string `json:"slug"`
	Type          string `json:"type"`
}

// ListOrganizations lists the organizations that the user has access to.
func ListOrganizations(apiToken string) ([]Organization, error) {
	url := fmt.Sprintf(TursoEndpoint + "/organizations")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Errorf("Error creating request. %v", err)
		return []Organization{}, fmt.Errorf("Error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("Error sending request. %v", err)
		return []Organization{}, fmt.Errorf("Error sending request. %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("Error reading body. %v", err)
		return []Organization{}, fmt.Errorf("Error reading body. %v", err)
	}
	listOrganizations, err := parseOrganizations(body, err)
	if err != nil {
		log.Errorf("Error decoding body into Organization. %v", err)
		return []Organization{}, fmt.Errorf("Error decoding body. %v", err)
	}
	return listOrganizations, nil
}

// parseOrganization parses the response from the server into a Organization.
func parseOrganizations(body []byte, err error) ([]Organization, error) {
	var response []Organization
	err = decoder.NewDecoder(string(body)).Decode(&response)
	return response, err
}

// UpdateOrganiation updates the organization with the given name.
func UpdateOrganiation(apiToken string, orgName string, orgType string) (Organization, error) {
	url := fmt.Sprintf(TursoEndpoint + "/organizations/" + orgName)
	reqJsonBody := fmt.Sprintf(`{"type": "%s"}`, orgType)
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer([]byte(reqJsonBody)))
	if err != nil {
		log.Errorf("Error creating request. %v", err)
		return Organization{}, fmt.Errorf("Error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("Error sending request. %v", err)
		return Organization{}, fmt.Errorf("Error sending request. %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("Error reading body. %v", err)
		return Organization{}, fmt.Errorf("Error reading body. %v", err)
	}
	updateOrganization, err := parseOrganization(body, err)
	if err != nil {
		log.Errorf("Error decoding body into Organization. %v", err)
		return Organization{}, fmt.Errorf("Error decoding body. %v", err)
	}
	return updateOrganization, nil
}

// parseOrganization parses the response from the server into a Organization.
func parseOrganization(body []byte, err error) (Organization, error) {
	var response Organization
	err = decoder.NewDecoder(string(body)).Decode(&response)
	return response, err
}
