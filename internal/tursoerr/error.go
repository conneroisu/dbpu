package tursoerr

import (
	"encoding/json"
	"fmt"
	"strings"
)

type (
	// ErrorResponse is the response returned by the Turso API.
	ErrorResponse struct {
		Error *APIError `json:"error,omitempty"`
	}

	// APIError provides error information returned by the Turso API.
	APIError struct {
		// Code is the code of the error.
		Code any `json:"code,omitempty"`
		// Message is the message of the error.
		Message string `json:"message"`
		// Param is the param of the error.
		Param *string `json:"param,omitempty"`
		// Type is the type of the error.
		Type string `json:"type"`
		// HTTPStatusCode is the status code of the error.
		HTTPStatusCode int `json:"-"`
	}

	// ErrRequest is an error returned by the Turso API.
	ErrRequest struct {
		Err            error
		HTTPStatusCode int
	}
)

// Error method implements the error interface on APIError.
func (e *APIError) Error() string {
	if e.HTTPStatusCode > 0 {
		return fmt.Sprintf(
			"error, status code: %d, message: %s",
			e.HTTPStatusCode,
			e.Message,
		)
	}
	return e.Message
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (e *APIError) UnmarshalJSON(data []byte) (err error) {
	var rawMap map[string]json.RawMessage
	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		return
	}
	err = json.Unmarshal(rawMap["message"], &e.Message)
	if err != nil {
		var messages []string
		err = json.Unmarshal(rawMap["message"], &messages)
		if err != nil {
			return
		}
		e.Message = strings.Join(messages, ", ")
	}
	// optional fields
	if _, ok := rawMap["param"]; ok {
		err = json.Unmarshal(rawMap["param"], &e.Param)
		if err != nil {
			return
		}
	}
	if _, ok := rawMap["code"]; !ok {
		return nil
	}
	// if the api returned a number, we need to force an integer
	// since the json package defaults to float64
	var intCode int
	err = json.Unmarshal(rawMap["code"], &intCode)
	if err == nil {
		e.Code = intCode
		return nil
	}
	return json.Unmarshal(rawMap["code"], &e.Code)
}

func (e *ErrRequest) Error() string {
	if e.Err == nil {
		return ""
	}
	return e.Err.Error()
}
