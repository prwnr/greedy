package npc

import (
	"testing"
)

func TestCreatesNewMonster(t *testing.T) {
	m := NewMonster(1)

	if m.Render() != "#" {
		t.Errorf("expected monster with '#' look, got '%s'", m.Render())
	}
}

func TestMonsterFighting(t *testing.T) {
	t.Run("monster is created with HP", func(t *testing.T) {
		first := NewMonster(1)

		firstGot := first.GetHealth()
		if firstGot != 30 {
			t.Errorf("monster level 1 should have 30 HP upon creation, has %d", firstGot)
		}

		second := NewMonster(2)

		secondGot := second.GetHealth()
		if secondGot != 60 {
			t.Errorf("monster level 2 should have 60 HP upon creation, has %d", secondGot)
		}
	})

	t.Run("monster HP can be reduced", func(t *testing.T) {
		m := NewMonster(1)
		m.ReduceHealth(20)

		got := m.GetHealth()
		if got != 10 {
			t.Errorf("monster should have 10 HP after reducing it by 20, has %d", got)
		}
	})

	t.Run("monster can attack", func(t *testing.T) {
		m := NewMonster(1)

		got := m.AttackPower()
		if got != 5 {
			t.Errorf("monster level 1 attack should equal 5, got %d", got)
		}
	})

	t.Run("monster is alive when his HP is above 0", func(t *testing.T) {
		m := NewMonster(1)

		if !m.IsAlive() {
			t.Error("monster should be alive")
		}
	})

	t.Run("monster is dead when his HP is below or at 0", func(t *testing.T) {
		m := NewMonster(1)
		m.ReduceHealth(500)

		if m.IsAlive() {
			t.Error("monster should be dead")
		}
	})
}
