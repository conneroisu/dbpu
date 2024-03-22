package dbpu

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/bytedance/sonic/decoder"
)

// parseStruct parses the response from a byte array into the provided type.
// T is a type parameter that will be replaced by any type that satisfies the any interface.
func parseStruct[T any](body []byte) (T, error) {
	var data T
	err := decoder.NewDecoder(string(body)).Decode(&data)
	if err != nil {
		return data, fmt.Errorf("error decoding body: %v", err)
	}
	return data, nil
}

// parseResponse parses the response from an HTTP request into the provided type.
// T is a type parameter that will be replaced by any type that satisfies the any interface.
func parseResponse[T any](response *http.Response) (T, error) {
	body, err := io.ReadAll(response.Body)
	var data T
	if err != nil {
		return data, fmt.Errorf("io failed to read response body: %v", err)
	}
	data, err = parseStruct[T](body)
	if err != nil {
		return data, fmt.Errorf("failed to parse response body: %v", err)
	}
	defer response.Body.Close()
	return data, nil
}

// resolveApiConfig is a configuration for resolving an API call.
type resolveApiConfig struct {
	ReqError error
	DoError  error
	ParError error
}

// resolveApiOpt is a functional option for setting the request, do, and parse errors in the resolveApiConfig.
type resolveApiOpt func(*resolveApiConfig)

// wReqError is a functional option for setting the request error in the resolveApiConfig.
func wReqError(err error) resolveApiOpt {
	return func(c *resolveApiConfig) {
		c.ReqError = err
	}
}

// wDoError is a functional option for setting the do error in the resolveApiConfig.
func wDoError(err error) resolveApiOpt {
	return func(c *resolveApiConfig) {
		c.DoError = err
	}
}

// wParError is a functional option for setting the parse error in the resolveApiConfig.
func wParError(err error) resolveApiOpt {
	return func(c *resolveApiConfig) {
		c.ParError = err
	}
}

// newResolveApiConfig creates a new resolveApiConfig with the provided options.
func newResolveApiConfig(opts ...resolveApiOpt) *resolveApiConfig {
	config := &resolveApiConfig{
		ReqError: nil,
		DoError:  nil,
		ParError: nil,
	}
	for _, opt := range opts {
		opt(config)
	}
	return config
}

// resolveApi resolves the API call by joining the request, do, and parse errors.
func resolveApi[Obj any](obj Obj, opts ...resolveApiOpt) (Obj, error) {
	config := newResolveApiConfig(opts...)
	if err := errors.Join(config.ReqError, config.DoError, config.ParError); err != nil {
		return obj, fmt.Errorf("error resolving API. %v", err)
	}
	return obj, nil
}
