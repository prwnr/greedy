package player

import (
	"fmt"
	"strings"
	"swarm/modifiers"
	"testing"
)

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
	h.Entity.Health = 10
	h.mana = 10

	wantHP := h.Entity.Health + modifiers.HeroHealthRegen
	wantMana := h.mana + modifiers.HeroManaRegen

	h.Regenerate()
	if h.GetHealth() != wantHP {
		t.Errorf("Hero health at %d, want %d", h.GetHealth(), wantHP)
	}

	if h.mana != wantMana {
		t.Errorf("Hero health at %d, want %d", h.GetHealth(), wantMana)
	}
}
