package entity

import "testing"

func TestHeroMovesDown(t *testing.T) {
	h := NewHero(0, 0)

	h.MoveDown(2)
	assertPosition(t, 1, h.Position.Y)

	h.MoveDown(2)
	h.MoveDown(2)
	assertPosition(t, 2, h.Position.Y)
}

func TestHeroMovesUp(t *testing.T) {
	h := NewHero(0, 2)

	h.MoveUp()
	assertPosition(t, 1, h.Position.Y)

	h.MoveUp()
	h.MoveUp()
	assertPosition(t, 0, h.Position.Y)
}

func TestHeroMovesLeft(t *testing.T) {
	h := NewHero(2, 0)

	h.MoveLeft()
	assertPosition(t, 1, h.Position.X)

	h.MoveLeft()
	h.MoveLeft()
	assertPosition(t, 0, h.Position.X)
}

func TestHeroMovesRight(t *testing.T) {
	h := NewHero(0, 0)

	h.MoveRight(2)
	assertPosition(t, 1, h.Position.X)

	h.MoveRight(2)
	h.MoveRight(2)
	assertPosition(t, 2, h.Position.X)
}

func assertPosition(t *testing.T, want, got int) {
	if want != got {
		t.Errorf("expected hero position %d, but got %d", want, got)
	}
}
