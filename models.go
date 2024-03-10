package main

import (
	"fmt"

	"github.com/bytedance/sonic/decoder"
)

// Jwt is a JSON Web Token.
type Jwt struct {
	Jwt string `json:"jwt"` // jwt is the JSON Web Token.
}

// parseStruct parses the response from a byte array into the provided type.
// T is a type parameter that will be replaced by any type that satisfies the AnyData interface.
func parseStruct[T any](body []byte) (T, error) {
	var data T
	err := decoder.NewDecoder(string(body)).Decode(&data)
	if err != nil {
		// Zero value of T is returned in case of an error. Use default(T) when Go 1.18+ syntax is not recognized.
		return data, fmt.Errorf("error decoding body: %v", err)
	}
	return data, nil
}
