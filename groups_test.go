package dbpu

import (
	"testing"
)

func TestParseGroups(t *testing.T) {
	t.Run("ParseGroups", func(t *testing.T) {
		body := []byte(`{"groups":[{"archived":false,"locations":["us-east-1"],"name":"test","primary":"us-east-1","uuid":"1"}]}`)
		response, err := parseStruct[ListGroupsResponse](body)
		if err != nil {
			t.Errorf("Error decoding body. %v", err)
		}
		if len(response.Groups) != 1 {
			t.Errorf("Expected 1 group, got %d", len(response.Groups))
		}
		if response.Groups[0].Name != "test" {
			t.Errorf("Expected group name to be test, got %s", response.Groups[0].Name)
		}
	})

	t.Run("ParseGroupsError", func(t *testing.T) {
		body := []byte(`{"groups":[{"archived":false,"locations":["us-east-1"],"name":"test","primary":"us-east-1","uuid":"1"}]`)
		_, err := parseStruct[ListGroupsResponse](body)
		if err == nil {
			t.Errorf("Expected error decoding body, got nil")
		}
	})
}
