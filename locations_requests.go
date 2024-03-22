package dbpu

import (
	"fmt"
	"net/http"
)

// newListLocationsReq creates a request for ListLocations.
func (c *Client) newListLocationsReq(apiToken string) (*http.Request, error) {
	url := fmt.Sprintf("%s/locations", c.BaseURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Error creating request. %v", err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	return req, nil
}

// newClosestLocationRequest creates a request for ClosestLocation.
func (c *Client) newClosestLocationRequest() (*http.Request, error) {
	url := fmt.Sprintf(c.RegionURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Error creating request. %v", err)
	}
	return req, nil
}
