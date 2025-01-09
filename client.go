package dbpu

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/conneroisu/dbpu/internal/builders"
	"github.com/conneroisu/dbpu/internal/tursoerr"
)

type (
	// Client is a client for the turso API.
	Client struct {
		client  *http.Client
		BaseURL string // Base URL for API requests.
		OrgName string // Name of organization.

		header   builders.Header
		apiToken string // Token for API.
	}
	// option is a functional option for configuring a Client.
	option func(*Client)
)

// WithClient sets the client for the Client.
func WithClient(client *http.Client) func(*Client) {
	return func(c *Client) { c.client = client }
}

// NewClient returns a new client.
//
// Base URL is the base URL for API requests.
// Region URL is the base URL for region requests.
func NewClient(apiToken, orgName string, opts ...option) *Client {
	client := &Client{
		client:   http.DefaultClient,
		BaseURL:  "https://api.turso.tech/v1",
		apiToken: apiToken,
		OrgName:  orgName,
	}
	client.header.SetCommonHeaders = func(req *http.Request) {
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf(
			"Bearer %s",
			client.apiToken,
		))
	}
	for _, opt := range opts {
		opt(client)
	}
	return client
}

func (c *Client) sendRequest(req *http.Request, v any) error {
	req.Header.Set("Accept", "application/json")
	// Check whether Content-Type is already set, Upload Files API requires
	// Content-Type == multipart/form-data
	contentType := req.Header.Get("Content-Type")
	if contentType == "" {
		req.Header.Set("Content-Type", "application/json")
	}
	res, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if isFailureStatusCode(res) {
		return c.handleErrorResp(res)
	}
	return decode(res.Body, v)
}

func (c *Client) handleErrorResp(resp *http.Response) error {
	var errRes tursoerr.ErrorResponse
	err := json.NewDecoder(resp.Body).Decode(&errRes)
	if err != nil || errRes.Error == nil {
		reqErr := &tursoerr.ErrRequest{
			HTTPStatusCode: resp.StatusCode,
			Err:            err,
		}
		if errRes.Error != nil {
			reqErr.Err = errRes.Error
		}
		return reqErr
	}
	errRes.Error.HTTPStatusCode = resp.StatusCode
	return errRes.Error
}

func isFailureStatusCode(resp *http.Response) bool {
	return resp.StatusCode < http.StatusOK ||
		resp.StatusCode >= http.StatusBadRequest
}

func decode[T any](r io.Reader, v T) error {
	return json.NewDecoder(r).Decode(&v)
}
