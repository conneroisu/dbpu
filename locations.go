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

// Locations is a list of locations.
type Locations struct {
	Locations map[string]string `json:"locations"`
}

// ClosestLocation returns the closest location to the given latitude and longitude.
func ClosestLocation() (ServerClient, error) {
	req, reqErr := newClosestLocationRequest()
	done, doErr := (&http.Client{}).Do(req)
	parsed, parErr := parseResponse[ServerClient](done)
	defer done.Body.Close()
	return resolveApiCall(parsed, wReqError(reqErr), wDoError(doErr), wParError(parErr))
}

// ListLocations returns a list of locations.
func ListLocations(apiToken string) (Locations, error) {
	req, reqErr := newListLocationsReq(apiToken)
	done, doErr := (&http.Client{}).Do(req)
	parsed, parErr := parseResponse[Locations](done)
	defer done.Body.Close()
	return resolveApiCall(parsed, wReqError(reqErr), wDoError(doErr), wParError(parErr))
}

// newListLocationsReq creates a request for ListLocations.
func newListLocationsReq(apiToken string) (*http.Request, error) {
	url := fmt.Sprintf("https://api.turso.tech/v1/locations")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Error creating request. %v", err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	return req, nil
}

// newClosestLocationRequest creates a request for ClosestLocation.
func newClosestLocationRequest() (*http.Request, error) {
	url := fmt.Sprintf("https://region.turso.io/")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Error creating request. %v", err)
	}
	return req, nil
}
