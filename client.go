package dbpu

import "net/http"

// Client is a client for the dbpu API.
type Client struct {
	*http.Client
	BaseURL   string
	RegionURL string
}

// NewClient returns a new client.
func NewClient() *Client {
	return &Client{
		Client:    http.DefaultClient,
		BaseURL:   "https://api.turso.tech/v1",
		RegionURL: "https://region.turso.io",
	}
}
