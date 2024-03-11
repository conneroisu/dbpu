package dbpu

import (
	"fmt"
	"net/http"
)

// ServerClient is a struct that contains the server and client locations.
type ServerClient struct {
	Server string `json:"server"`
	Client string `json:"client"`
}

// CreateClosestLocationRequest creates a request for ClosestLocation.
func CreateClosestLocationRequest() (*http.Request, error) {
	url := fmt.Sprintf("https://region.turso.io/")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Error creating request. %v", err)
	}
	return req, nil
}

// ClosestLocation returns the closest location to the given latitude and longitude.
func ClosestLocation() (ServerClient, error) {
	req, err := CreateClosestLocationRequest()
	if err != nil {
		return ServerClient{}, fmt.Errorf("Error reading request. %v", err)
	}
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return ServerClient{}, fmt.Errorf("Error sending request. %v", err)
	}
	response, err := parseResponse[ServerClient](resp)
	defer resp.Body.Close()
	return response, nil
}
