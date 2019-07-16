package npc

import "testing"

func TestCreatesNewMonster(t *testing.T) {
	m := NewMonster()

	if m.Render() != "#" {
		t.Errorf("expected monster with '#' look, got '%s'", m.Render())
	}
}
