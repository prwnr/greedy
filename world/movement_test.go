package world

import (
	"swarm/hero"
	"testing"
)

func TestMovesHeroOnLocation(t *testing.T) {
	m := NewLocation(4)
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

func TestWontMoveHeroOutsideLocation(t *testing.T) {
	l := NewLocation(4)
	h := hero.Bee{}
	h.Start(0, 0)

	Move(&h, &l, MoveUp)
	if h.Position.Y != 0 {
		t.Errorf("hero Y position should be 0, was %d", h.Position.Y)
	}

	Move(&h, &l, MoveLeft)
	if h.Position.X != 0 {
		t.Errorf("hero X position should be 1, was %d", h.Position.X)
	}
}
