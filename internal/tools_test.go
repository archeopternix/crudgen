package internal

import (
	"testing"
)

func TestIsLetter(t *testing.T) {
	if !IsLetter("Azz019aZ") {
		t.Errorf("Contains only letters 'Azz019aZ'")
	}

	if IsLetter("Az$!_z019aZ") {
		t.Errorf("Contains not only letters 'Az$!_z019aZ'")
	}

	if IsLetter(" AB") {
		t.Errorf("Contains not only letters ' AB'")
	}
}
