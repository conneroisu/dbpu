package dbpu

import (
	"testing"
)

func TestParseGroups(t *testing.T) {
	t.Run("Testing ListGroupsResponse Parsing", func(t *testing.T) {
		body := []byte(`{"groups":[{"archived":false,"locations":["us-east-1"],"name":"test","primary":"us-east-1","uuid":"1"}]}`)
		response, err := parseStruct[ListGroupsResp](body)
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

	t.Run("Testing ListGroupsResponse Parsing with bad body", func(t *testing.T) {
		body := []byte(`{"groups":[{"archived":false,"locations":["us-east-1"],"name":"test","primary":"us-east-1","uuid":"1"}]`)
		_, err := parseStruct[ListGroupsResp](body)
		if err == nil {
			t.Errorf("Expected error decoding body, got nil")
		}
	})

	t.Run("Testing GroupResponse Parsing", func(t *testing.T) {
		body := []byte(`{"group":{"archived":false,"locations":["us-east-1"],"name":"test","primary":"us-east-1","uuid":"1"}}`)
		response, err := parseStruct[GroupResp](body)
		if err != nil {
			t.Errorf("Error decoding body. %v", err)
		}
		if response.Group.Name != "test" {
			t.Errorf("Expected group name to be test, got %s", response.Group.Name)
		}
	})

	t.Run("Testing GroupResponse Parsing with bad body", func(t *testing.T) {
		body := []byte(`{"group":{"archived":false,"locations":["us-east-1"],"name":"test","primary":"us-east-1","uuid":"1"`)
		_, err := parseStruct[GroupResp](body)
		if err == nil {
			t.Errorf("Expected error decoding body, got nil")
		}
	})

	t.Run("Testing GroupResponse Parsing with bad body", func(t *testing.T) {
		body := []byte(`{"group":{"archived":false,"locations":["us-east-1"],"name":"test","primary":"us-east-1","uuid":"1"`)
		_, err := parseStruct[GroupResp](body)
		if err == nil {
			t.Errorf("Expected error decoding body, got nil")
		}
	})

	t.Run("Testing Group Parsing", func(t *testing.T) {
		body := []byte(`{"archived":false,"locations":["us-east-1"],"name":"test","primary":"us-east-1","uuid":"1"}`)
		response, err := parseStruct[Group](body)
		if err != nil {
			t.Errorf("Error decoding body. %v", err)
		}
		if response.Name != "test" {
			t.Errorf("Expected group name to be test, got %s", response.Name)
		}
	})

	t.Run("Testing Group Parsing with bad body", func(t *testing.T) {
		body := []byte(`{"archived":false,"locations":["us-east-1"],"name":"test","primary":"us-east-1","uuid":"1"`)
		_, err := parseStruct[Group](body)
		if err == nil {
			t.Errorf("Expected error decoding body, got nil")
		}
	})

}
