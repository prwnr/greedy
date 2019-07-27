package player

import (
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

func TestHeroFightingMethods(t *testing.T) {
	t.Run("hero is created with HP", func(t *testing.T) {
		h := NewHero(0, 0)

		got := h.GetHP()
		if got != h.health {
			t.Errorf("hero should have 100 HP upon creation, has %d", got)
		}
	})

	t.Run("hero HP can be reduced", func(t *testing.T) {
		h := NewHero(0, 0)
		h.ReduceHealth(50)

		got := h.GetHP()
		if got != h.health {
			t.Errorf("hero should have 50 HP after reducing it by 50, has %d", got)
		}
	})

	t.Run("hero can attack", func(t *testing.T) {
		h := NewHero(0, 0)

		got := h.Attack()
		if got != h.attack {
			t.Errorf("hero attack should equal 100, got %d", got)
		}
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
		assertHealthEquals(t, 90, h.GetHP())

		got := h.UseHeal()

		assertHealthEquals(t, 92, h.GetHP())
		if got != "Hero health restored by 2." {
			t.Errorf("healing message is invalid, got '%s'", got)
		}
	})

	t.Run("hero cannot over heal", func(t *testing.T) {
		h := NewHero(0, 0)
		assertHealthEquals(t, 100, h.GetHP())

		got := h.UseHeal()

		assertHealthEquals(t, 100, h.GetHP())
		if got != "Hero health restored by 0." {
			t.Errorf("healing message is invalid, got '%s'", got)
		}
	})

	t.Run("hero cannot heal when mana is low", func(t *testing.T) {
		h := NewHero(0, 0)

		h.ReduceHealth(10)
		h.mana = 0

		assertHealthEquals(t, 90, h.GetHP())

		got := h.UseHeal()

		assertHealthEquals(t, 90, h.GetHP())
		if got != "Mana is too low." {
			t.Errorf("returned message should be about low mana, got '%s'", got)
		}
	})
}
