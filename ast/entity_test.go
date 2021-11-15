package ast

import (
	"testing"
)

func TestEntityCheckForErrors(t *testing.T) {
	// Add a new entity
	if err := EntityCheckForErrors(Entity{Name: "Alpha"}); err == nil {
		t.Errorf("Creation of entity should fail: %v", err)
	}

	// Add entity with too short name
	if err := EntityCheckForErrors(Entity{Name: "A1"}); err == nil {
		t.Error("Creation of entity with too short name must fail")
	}

	// Add entity with too short name
	if err := EntityCheckForErrors(Entity{Name: "A!$§6", Kind: "default"}); err == nil {
		t.Error("Name of entity must contain only letters and numbers")
	}

	// Add entity with wring kind
	if err := EntityCheckForErrors(Entity{Name: "A1", Kind: "dinosaur"}); err == nil {
		t.Error("Creation of entity with unknown kind must fail")
	}

	// Add entity with too short name
	if err := EntityCheckForErrors(Entity{Name: "Albert", Kind: "default"}); err != nil {
		t.Error("Entity has to be created")
	}

}
