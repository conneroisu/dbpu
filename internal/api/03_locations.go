package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ServerClient is a struct that contains the server and client locations.
type ServerClient struct {
	Server string `json:"server"`
	Client string `json:"client"`
}

// ClosestLocation returns the closest location to the given latitude and longitude.
func ClosestLocation() (ServerClient, error) {
	url := fmt.Sprintf("https://region.turso.io/")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ServerClient{}, fmt.Errorf("Error creating request. %v", err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return ServerClient{}, fmt.Errorf("Error sending request. %v", err)
	}
	defer resp.Body.Close()
	response, err := parseServerClient(resp.Body, err)
	if err != nil {
		return ServerClient{}, fmt.Errorf("Error decoding body. %v", err)
	}
	return response, nil
}

// parseServerClient parses the response from the server into a ServerClient.
func parseServerClient(body io.Reader, err error) (ServerClient, error) {
	var response ServerClient
	err = json.NewDecoder(body).Decode(&response)
	return response, err
}
