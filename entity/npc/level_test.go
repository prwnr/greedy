package npc

import (
	"testing"
)

func TestNewLevel(t *testing.T) {
	l := NewLevel(3)
	assertLevelEquals(t, 3, l.Number)

	l = NewLevel(2)
	assertLevelEquals(t, 2, l.Number)
}

func assertLevelEquals(t *testing.T, want, got int) {
	if want != got {
		t.Errorf("level should be at %d, got %d", want, got)
	}
}
