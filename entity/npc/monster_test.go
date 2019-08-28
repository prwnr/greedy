package npc

import (
	"swarm/modifiers"
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

		got := first.GetHealth()
		want := modifiers.MonsterBaseHealth
		if got != want {
			t.Errorf("monster level 1 should have %d HP upon creation, has %d", want, got)
		}

		second := NewMonster(2)

		got = second.GetHealth()
		want = modifiers.CalculateMonsterHealth(2)
		if got != want {
			t.Errorf("monster level 2 should have %d HP upon creation, has %d", want, got)
		}
	})

	t.Run("monster HP can be reduced", func(t *testing.T) {
		m := NewMonster(1)
		m.ReduceHealth(20)

		got := m.GetHealth()
		want := modifiers.MonsterBaseHealth - 20
		if got != want {
			t.Errorf("monster should have %d HP after reducing it by 20, has %d", want, got)
		}
	})

	t.Run("monster can attack", func(t *testing.T) {
		m := NewMonster(1)

		got := m.AttackPower()
		if got != modifiers.MonsterBaseAttack {
			t.Errorf("monster level 1 attack should equal %d, got %d", modifiers.MonsterBaseAttack, got)
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

func TestMonsterGetExperienceValue(t *testing.T) {
	tests := []struct {
		name  string
		level int
		want  int
	}{
		{"1 lvl mob gives 15 exp", 1, modifiers.MonsterBaseExperience * 1},
		{"2 lvl mob gives 30 exp", 2, modifiers.MonsterBaseExperience * 2},
		{"5 lvl mob gives 75 exp", 5, modifiers.MonsterBaseExperience * 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMonster(tt.level)

			if got := m.GetExperienceValue(); got != tt.want {
				t.Errorf("GetExperienceValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
