package board

import "testing"

func TestMapCreation(t *testing.T) {
	got := NewMap(2)

	if len(got.Fields) != 2 {
		t.Errorf("got %v elements, want 2", got)
	}

	if len(got.Fields[0]) != 2 {
		t.Errorf("got %v elements on first row, want 2", got)
	}
}
