package npc

import (
	"testing"
)

func TestCreatesNewMonster(t *testing.T) {
	m := NewMonster()

	if m.Render() != "#" {
		t.Errorf("expected monster with '#' look, got '%s'", m.Render())
	}
}

func TestMonsterFighting(t *testing.T) {
	t.Run("monster is created with HP", func(t *testing.T) {
		m := NewMonster()

		got := m.GetHP()
		if got != m.hp {
			t.Errorf("monster should have 100 HP upon creation, has %d", got)
		}
	})

	t.Run("monster HP can be reduced", func(t *testing.T) {
		m := NewMonster()
		m.ReduceHealth(50)

		got := m.GetHP()
		if got != m.hp {
			t.Errorf("monster should have 50 HP after reducing it by 50, has %d", got)
		}
	})

	t.Run("monster can attack", func(t *testing.T) {
		m := NewMonster()

		got := m.Attack()
		if got != m.attack {
			t.Errorf("monster attack should equal 40, got %d", got)
		}
	})

	t.Run("monster is alive when his HP is above 0", func(t *testing.T) {
		m := NewMonster()

		if !m.IsAlive() {
			t.Error("monster should be alive")
		}
	})

	t.Run("monster is dead when his HP is below or at 0", func(t *testing.T) {
		m := NewMonster()
		m.ReduceHealth(500)

		if m.IsAlive() {
			t.Error("monster should be dead")
		}
	})
}
