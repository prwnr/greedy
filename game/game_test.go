package game

import (
	"testing"
)

func TestMovesHeroOnLocation(t *testing.T) {
	g := NewGame()
	g.Hero.Position.X = 0
	g.Hero.Position.Y = 0

	g.MoveHero(MoveRight)
	if g.Hero.Position.X != 1 {
		t.Errorf("hero X position should be 1, was %d", g.Hero.Position.X)
	}

	g.MoveHero(MoveDown)
	if g.Hero.Position.Y != 1 {
		t.Errorf("hero Y position should be 1, was %d", g.Hero.Position.Y)
	}
}

func TestMovingHeroToNewLocationRemovesHimFromOld(t *testing.T) {
	g := NewGame()
	g.Hero.Position.X = 0
	g.Hero.Position.Y = 0

	g.MoveHero(MoveRight)
	got := g.CurrentLocation.Places[0][0].GetHero()
	if got != nil {
		t.Error("there should be no here on this place, got one")
	}
}

func TestWontMoveHeroOutsideLocation(t *testing.T) {
	g := NewGame()
	g.Hero.Position.X = 0
	g.Hero.Position.Y = 0

	g.MoveHero(MoveUp)
	if g.Hero.Position.Y != 0 {
		t.Errorf("hero Y position should be 0, was %d", g.Hero.Position.Y)
	}

	g.MoveHero(MoveLeft)
	if g.Hero.Position.X != 0 {
		t.Errorf("hero X position should be 0, was %d", g.Hero.Position.X)
	}
}
