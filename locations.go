package dbpu

// ClosestLocation returns the closest location to the given latitude and longitude.
func (c *Client) ClosestLocation() (ServerClient, error) {
	req, reqErr := c.newClosestLocationRequest()
	done, doErr := c.Do(req)
	parsed, parErr := parseResponse[ServerClient](done)
	defer done.Body.Close()
	return resolveApi(parsed, wReqError(reqErr), wDoError(doErr), wParError(parErr))
}

// ListLocations returns a list of locations.
func (c *Client) ListLocations(apiToken string) (Locations, error) {
	req, reqErr := c.newListLocationsReq(apiToken)
	done, doErr := c.Do(req)
	parsed, parErr := parseResponse[Locations](done)
	defer done.Body.Close()
	return resolveApi(parsed, wReqError(reqErr), wDoError(doErr), wParError(parErr))
}
