package dbpu

import (
	"testing"
)

func testParseStruct(t *testing.T) {
	t.Run("Test Parse Struct", func(t *testing.T) {
		// Test the parse struct.
		body := []byte(`{"jwt":"test"}`)
		jwt, err := parseStruct[Jwt](body)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if jwt.Jwt != "test" {
			t.Errorf("Expected Jwt to be test, got %v", jwt.Jwt)
		}
	})

	t.Run("Test new struct parse", func(t *testing.T) {
		type TestStruct struct {
			Test     string `json:"test"`
			Insomnia string `json:"insomnia"`
		}
		body := []byte(`{"test":"test","insomnia":"insomnia"}`)
		test, err := parseStruct[TestStruct](body)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if test.Test != "test" {
			t.Errorf("Expected Test to be test, got %v", test.Test)
		}
	})
}
