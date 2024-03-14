package dbpu

import (
	"testing"
)

func testUpdateOrganizationConfigs(t *testing.T) {

	t.Run("Test With Blocked Writes", func(t *testing.T) {
		org := NewUpdateOrganiationConfig(Org{}, WithBlockedWrites(true))
		if org.BlockedWrites != true {
			t.Errorf("Expected BlockedWrites to be true, got %v", org.BlockedWrites)
		}
	})

	t.Run("Test With Blocked Reads", func(t *testing.T) {
		org := NewUpdateOrganiationConfig(Org{}, WithBlockedReads(true))
		if org.BlockedReads != true {
			t.Errorf("Expected BlockedReads to be true, got %v", org.BlockedReads)
		}
	})

	t.Run("Test With Blocked Writes and Reads", func(t *testing.T) {
		org := NewUpdateOrganiationConfig(Org{}, WithBlockedWrites(true), WithBlockedReads(true))
		if org.BlockedWrites != true {
			t.Errorf("Expected BlockedWrites to be true, got %v", org.BlockedWrites)
		}
		if org.BlockedReads != true {
			t.Errorf("Expected BlockedReads to be true, got %v", org.BlockedReads)
		}
	})

	t.Run("Test With Name", func(t *testing.T) {
		org := NewUpdateOrganiationConfig(Org{}, WithName("test"))
		if org.Name != "test" {
			t.Errorf("Expected Name to be test, got %v", org.Name)
		}
	})

	t.Run("Test With Overages", func(t *testing.T) {
		org := NewUpdateOrganiationConfig(Org{}, WithOverages(true))
		if org.Overages != true {
			t.Errorf("Expected Overages to be true, got %v", org.Overages)
		}
	})

	t.Run("Test With Slug", func(t *testing.T) {
		org := NewUpdateOrganiationConfig(Org{}, WithSlug("test"))
		if org.Slug != "test" {
			t.Errorf("Expected Slug to be test, got %v", org.Slug)
		}
	})
}

// testOrganizationsParse tests the parsing of organization objects.
func testOrganizationsParse(t *testing.T) {
	t.Run("Test Parse Organizations", func(t *testing.T) {
		body := []byte(`[{"name":"test","slug":"test","type":"test","overages":true,"blocked_reads":true,"blocked_writes":true}]`)
		orgs, err := parseStruct[[]Org](body)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if len(orgs) != 1 {
			t.Errorf("Expected orgs to have 1 element, got %v", len(orgs))
		}
		if orgs[0].Name != "test" {
			t.Errorf("Expected Name to be test, got %v", orgs[0].Name)
		}
		if orgs[0].Slug != "test" {
			t.Errorf("Expected Slug to be test, got %v", orgs[0].Slug)
		}
		if orgs[0].Type != "test" {
			t.Errorf("Expected Type to be test, got %v", orgs[0].Type)
		}
		if orgs[0].Overages != true {
			t.Errorf("Expected Overages to be true, got %v", orgs[0].Overages)
		}
		if orgs[0].BlockedReads != true {
			t.Errorf("Expected BlockedReads to be true, got %v", orgs[0].BlockedReads)
		}
		if orgs[0].BlockedWrites != true {
			t.Errorf("Expected BlockedWrites to be true, got %v", orgs[0].BlockedWrites)
		}
	})
}
