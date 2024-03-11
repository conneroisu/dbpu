package dbpu

import (
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
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return ServerClient{}, fmt.Errorf("Error sending request. %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ServerClient{}, fmt.Errorf("Error decoding body. %v", err)
	}
	response, err := parseStruct[ServerClient](body)
	return response, nil
}
