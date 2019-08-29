package entity

import (
	"fmt"
	"strings"
	"swarm/modifiers"
	"testing"
)

func TestEntityAttackPower(t *testing.T) {
	type fields struct {
		Health int
		Attack int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"has 10 attack power", fields{Health: 10, Attack: 10}, 10},
		{"has 15 attack power", fields{Health: 10, Attack: 15}, 15},
		{"has 25 attack power", fields{Health: 10, Attack: 25}, 25},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Entity{
				Health: tt.fields.Health,
				Attack: tt.fields.Attack,
			}
			if got := e.AttackPower(); got != tt.want {
				t.Errorf("AttackPower() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntityGetHealth(t *testing.T) {
	type fields struct {
		Health int
		Attack int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"has 10 health", fields{Health: 10, Attack: 1}, 10},
		{"has 15 health", fields{Health: 15, Attack: 1}, 15},
		{"has 25 health", fields{Health: 25, Attack: 1}, 25},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Entity{
				Health: tt.fields.Health,
				Attack: tt.fields.Attack,
			}
			if got := e.GetHealth(); got != tt.want {
				t.Errorf("GetHealth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntityIsAlive(t *testing.T) {
	type fields struct {
		Health int
		Attack int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"is alive with health above zero", fields{Health: 10, Attack: 10}, true},
		{"is dead with health at zero", fields{Health: 0, Attack: 10}, false},
		{"is dead with health below zero", fields{Health: -100, Attack: 10}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Entity{
				Health: tt.fields.Health,
				Attack: tt.fields.Attack,
			}
			if got := e.IsAlive(); got != tt.want {
				t.Errorf("IsAlive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntityReduceHealth(t *testing.T) {
	type fields struct {
		Health int
		Attack int
	}
	tests := []struct {
		name   string
		fields fields
		amount int
		want   int
	}{
		{"reduces health to zero", fields{Health: 10, Attack: 10}, 10, 0},
		{"reduces health to 50%", fields{Health: 20, Attack: 10}, 10, 10},
		{"reduces health below zero", fields{Health: 20, Attack: 10}, 30, -10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Entity{
				Health: tt.fields.Health,
				Attack: tt.fields.Attack,
			}

			e.ReduceHealth(tt.amount)

			if got := e.GetHealth(); got != tt.want {
				t.Errorf("After ReduceHealth(), got GetHealth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetsStartingHeroPosition(t *testing.T) {
	h := NewHero(1, 2)

	if h.Position.X != 1 || h.Position.Y != 2 {
		t.Errorf("got wrong hero position, expected 1;2. got %d;%d", h.Position.X, h.Position.Y)
	}
}

func TestShowsHeroLook(t *testing.T) {
	h := NewHero(0, 0)

	got := h.Render()
	if strings.Compare(got, "*") != 0 {
		t.Errorf("got '%s', expected '*'", got)
	}
}

func TestHeroFighting(t *testing.T) {
	assertAttackBetween := func(t *testing.T, actual, min, max int) {
		if actual < min || actual > max {
			t.Errorf("attack should be between %d and %d, got %d", min, max, actual)
		}
	}

	t.Run("hero is created with HP", func(t *testing.T) {
		h := NewHero(0, 0)

		got := h.GetHealth()
		if got != modifiers.HeroBaseHealth {
			t.Errorf("hero should have 100 HP upon creation, has %d", got)
		}
	})

	t.Run("hero HP can be reduced", func(t *testing.T) {
		h := NewHero(0, 0)
		reduce := 50
		h.ReduceHealth(reduce)

		got := h.GetHealth()
		if got != modifiers.HeroBaseHealth-reduce {
			t.Errorf("hero should have 50 HP after reducing it by 50, has %d", got)
		}
	})

	t.Run("hero can attack", func(t *testing.T) {
		h := NewHero(0, 0)

		got := h.AttackPower()
		assertAttackBetween(t, got, 10, 15)
	})

	t.Run("hero is alive when his HP is above 0", func(t *testing.T) {
		h := NewHero(0, 0)

		if !h.IsAlive() {
			t.Error("hero should be alive")
		}
	})

	t.Run("hero is dead when his HP is below or at 0", func(t *testing.T) {
		h := NewHero(0, 0)
		h.ReduceHealth(500)

		if h.IsAlive() {
			t.Error("hero should be dead")
		}
	})
}

func TestHeroGainExperience(t *testing.T) {
	tests := []struct {
		name      string
		gainedExp int
		want      string
	}{
		{"gains 15 experience", 15, fmt.Sprintf("Gained %d experience.\r\n", 15)},
		{"gains 40 experience", 40, fmt.Sprintf("Gained %d experience.\r\n", 40)},
		{"gains 80 experience", 80, fmt.Sprintf("Gained %d experience.\r\n", 80)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHero(0, 0)
			if got := h.GainExperience(tt.gainedExp); got != tt.want {
				t.Errorf("GainExperience() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeroLevelUp(t *testing.T) {
	assertHeroLevel := func(h *Hero, want int) {
		if h.level.Number != want {
			t.Errorf("Hero should be at level %d, got %d", want, h.level.Number)
		}
	}

	t.Run("hero levels up gaining experience", func(t *testing.T) {
		h := NewHero(0, 0)

		h.GainExperience(200)
		assertHeroLevel(h, 2)
	})

	t.Run("hero reaches maximum level", func(t *testing.T) {
		h := NewHero(0, 0)

		assertHeroLevel(h, 1)
		for i := 1; i < modifiers.HeroMaxLevel; i++ {
			exp := h.level.Next.ReqExperience - h.experience
			h.GainExperience(exp)
			assertHeroLevel(h, i+1)
		}

		h.GainExperience(100)
		assertHeroLevel(h, modifiers.HeroMaxLevel)

		if !h.HasMaxLevel() {
			t.Error("Hero should be at max level")
		}
	})
}

func TestHeroRegenerate(t *testing.T) {
	h := NewHero(0, 0)
	h.Health = 10
	h.mana = 10

	wantHP := h.Health + modifiers.HeroHealthRegen
	wantMana := h.mana + modifiers.HeroManaRegen

	h.Regenerate()
	if h.GetHealth() != wantHP {
		t.Errorf("Hero health at %d, want %d", h.GetHealth(), wantHP)
	}

	if h.mana != wantMana {
		t.Errorf("Hero health at %d, want %d", h.GetHealth(), wantMana)
	}
}

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
