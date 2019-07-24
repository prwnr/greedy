package world

import (
	"testing"
)

func TestLocationCreation(t *testing.T) {
	got := NewLocation(5)

	if len(got.Places) != 5 {
		t.Errorf("got %v elements, want 2", got)
	}

	if len(got.Places[0]) != 5 {
		t.Errorf("got %v elements on first row, want 2", got)
	}
}
