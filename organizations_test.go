package dbpu

import (
	"testing"
)

func TestUpdateOrganizationConfigs(t *testing.T) {
	// Test the update organization configs.
	// ...
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
