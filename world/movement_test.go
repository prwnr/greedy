package world

import (
	"swarm/player"
	"testing"
)

func TestMovesHeroOnLocation(t *testing.T) {
	m := NewLocation(4)
	h := player.Hero{}
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

func TestMovingHeroToNewLocationRemovesHimFromOld(t *testing.T) {
	m := NewLocation(4)
	h := player.Hero{}
	h.Start(0, 0)

	Move(&h, &m, MoveRight)
	got := len(m.Places[0][0].GetAvatars())
	if got != 0 {
		t.Errorf("number of avatars on old place should be 0, got %d", got)
	}
}

func TestWontMoveHeroOutsideLocation(t *testing.T) {
	l := NewLocation(4)
	h := player.Hero{}
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
