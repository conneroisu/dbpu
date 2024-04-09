package dbpu

import "net/http"

// Client is a client for the dbpu API.
type Client struct {
	*http.Client
	BaseURL   string
	RegionURL string
	OrgToken  string
	OrgName   string
	GroupName string
	ApiToken  string
}

// NewClient returns a new client.
func NewClient() *Client {
	return &Client{
		Client:    http.DefaultClient,
		BaseURL:   "https://api.turso.tech/v1",
		RegionURL: "https://region.turso.io",
	}
}

func (c *Client) SetOrgToken(token string) { c.OrgToken = token }

func (c *Client) SetOrgName(name string) { c.OrgName = name }

func (c *Client) SetGroupName(name string) { c.GroupName = name }

func (c *Client) SetApiToken(token string) { c.ApiToken = token }
