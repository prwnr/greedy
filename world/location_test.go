package world

import (
	"testing"
)

func TestLocationCreation(t *testing.T) {
	got := NewLocation(2)

	if len(got.Places) != 2 {
		t.Errorf("got %v elements, want 2", got)
	}

	if len(got.Places[0]) != 2 {
		t.Errorf("got %v elements on first row, want 2", got)
	}
}
