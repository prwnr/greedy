package player

import (
	"strings"
	"testing"
)

func TestSetsStartingHeroPosition(t *testing.T) {
	h := NewHero()
	h.StartingPosition(1, 2)

	if h.Position.X != 1 || h.Position.Y != 2 {
		t.Errorf("got wrong hero position, expected 1;2. got %d;%d", h.Position.X, h.Position.Y)
	}
}

func TestShowsHeroLook(t *testing.T) {
	h := NewHero()

	got := h.Render()
	if strings.Compare(got, "*") != 0 {
		t.Errorf("got '%s', expected '*'", got)
	}
}

func TestHeroFighting(t *testing.T) {
	t.Run("hero is created with HP", func(t *testing.T) {
		h := NewHero()

		got := h.GetHP()
		if got != 100 {
			t.Errorf("hero should have 100 HP upon creation, has %d", got)
		}
	})

	t.Run("hero HP can be reduced", func(t *testing.T) {
		h := NewHero()
		h.ReduceHP(50)

		got := h.GetHP()
		if got != 50 {
			t.Errorf("hero should have 50 HP after reducing it by 50, has %d", got)
		}
	})

	t.Run("hero can attack", func(t *testing.T) {
		h := NewHero()

		got := h.Attack()
		if got != 100 {
			t.Errorf("hero attack should equal 100, got %d", got)
		}
	})

	t.Run("hero is alive when his HP is above 0", func(t *testing.T) {
		h := NewHero()

		if !h.IsAlive() {
			t.Error("hero should be alive")
		}
	})

	t.Run("hero is dead when his HP is below or at 0", func(t *testing.T) {
		h := NewHero()
		h.ReduceHP(500)

		if h.IsAlive() {
			t.Error("hero should be dead")
		}
	})
}
