package dbpu

import "fmt"

// ClosestLocation returns the closest location to the given latitude and longitude.
func (c *Client) ClosestLocation() (*ServerClient, error) {
	req, err := c.newClosestLocationRequest()
	if err != nil {
		return nil, err
	}
	done, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error doing request: %v", err)
	}
	parsed, err := parseResponse[ServerClient](done)
	if err != nil {
		return nil, err
	}
	defer done.Body.Close()
	return &parsed, nil
}

// ListLocations returns a list of locations.
func (c *Client) ListLocations(apiToken string) (*Locations, error) {
	req, err := c.newListLocationsReq(apiToken)
	if err != nil {
		return nil, err
	}
	done, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error doing request: %v", err)
	}
	parsed, err := parseResponse[Locations](done)
	if err != nil {
		return nil, fmt.Errorf("error parsing response: %v", err)
	}
	defer done.Body.Close()
	return &parsed, nil
}
