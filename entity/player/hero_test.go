package player

import (
	"fmt"
	"strings"
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
		if got != 100 {
			t.Errorf("hero should have 100 HP upon creation, has %d", got)
		}
	})

	t.Run("hero HP can be reduced", func(t *testing.T) {
		h := NewHero(0, 0)
		h.ReduceHealth(50)

		got := h.GetHealth()
		if got != 50 {
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

func TestHeroSkills(t *testing.T) {
	assertHealthEquals := func(t *testing.T, want, got int) {
		if want != got {
			t.Errorf("expected hero health to be at %d, but got %d", want, got)
		}
	}

	t.Run("hero heals himself", func(t *testing.T) {
		h := NewHero(0, 0)

		h.ReduceHealth(10)
		assertHealthEquals(t, 90, h.GetHealth())

		got := h.UseHeal()

		assertHealthEquals(t, 95, h.GetHealth())
		if got != "Hero health restored by 5." {
			t.Errorf("healing message is invalid, got '%s'", got)
		}
	})

	t.Run("hero cannot over heal", func(t *testing.T) {
		h := NewHero(0, 0)
		assertHealthEquals(t, 100, h.GetHealth())

		got := h.UseHeal()

		assertHealthEquals(t, 100, h.GetHealth())
		if got != "Hero health restored by 0." {
			t.Errorf("healing message is invalid, got '%s'", got)
		}
	})

	t.Run("hero cannot heal when mana is low", func(t *testing.T) {
		h := NewHero(0, 0)

		h.ReduceHealth(10)
		h.mana = 0

		assertHealthEquals(t, 90, h.GetHealth())

		got := h.UseHeal()

		assertHealthEquals(t, 90, h.GetHealth())
		if got != "Mana is too low." {
			t.Errorf("returned message should be about low mana, got '%s'", got)
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
		for i := 1; i < 5; i++ {
			exp := h.level.Next.ReqExperience - h.experience
			h.GainExperience(exp)
			assertHeroLevel(h, i+1)
		}

		h.GainExperience(100)
		assertHeroLevel(h, 5)

		if !h.MaxLevel() {
			t.Error("Hero should be at max level")
		}
	})
}
