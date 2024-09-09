package dbpu

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// parseStruct parses the response from a byte array into the provided type.
// T is a type parameter that will be replaced by any type that satisfies the any interface.
func parseStruct[T any](body []byte) (T, error) {
	var data T
	err := json.NewDecoder(bytes.NewReader(body)).Decode(&data)
	if err != nil {
		return data, fmt.Errorf("error decoding body: %v", err)
	}
	return data, nil
}

// parseResponse parses the response from an HTTP request into the provided type.
// T is a type parameter that will be replaced by any type that satisfies the any interface.
func parseResponse[T any](response *http.Response) (*T, error) {
	body, err := io.ReadAll(response.Body)
	var data T
	if err != nil {
		return nil, fmt.Errorf("io failed to read response body: %v", err)
	}
	data, err = parseStruct[T](body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response body: %v", err)
	}
	defer response.Body.Close()
	return &data, nil
}
