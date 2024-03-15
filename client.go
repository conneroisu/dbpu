package dbpu

import "net/http"

type Client struct {
	*http.Client
	BaseURL string
}
