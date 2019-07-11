package board

import (
	"swarm/hero"
	"testing"
)

func TestMovesHeroOnMap(t *testing.T) {
	m := NewMap(4)
	h := hero.Bee{}
	h.Start(0, 0)

	Move(&h, &m, MoveRight)
	if h.Position.X != 1 {
		t.Errorf("hero X position should be 1, was %d", h.Position.X)
	}

	Move(&h, &m, MoveDown)
	if h.Position.Y != 1 {
		t.Errorf("hero Y position should be 1, was %d", h.Position.Y)
	}
}

func TestWontMoveHeroOutsideMap(t *testing.T) {
	m := NewMap(4)
	h := hero.Bee{}
	h.Start(0, 0)

	Move(&h, &m, MoveUp)
	if h.Position.Y != 0 {
		t.Errorf("hero Y position should be 0, was %d", h.Position.Y)
	}

	Move(&h, &m, MoveLeft)
	if h.Position.X != 0 {
		t.Errorf("hero X position should be 1, was %d", h.Position.X)
	}
}
