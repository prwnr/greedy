package world

import (
	"strings"
	"swarm/npc"
	"swarm/player"
	"testing"
)

func TestPlaceHero(t *testing.T) {
	t.Run("can set new Hero on place", func(t *testing.T) {
		p := Place{}
		h := &player.Hero{}

		p.SetHero(h)

		if p.GetHero() == nil {
			t.Error("expected a hero, got nil")
		}
	})

	t.Run("can remove all Hero from place", func(t *testing.T) {
		p := Place{}
		h := &player.Hero{}

		p.SetHero(h)
		p.RemoveHero()

		if p.GetHero() != nil {
			t.Error("expected no hero, got one")
		}
	})
}

func TestPlaceMonsters(t *testing.T) {
	t.Run("can add new Monster on place", func(t *testing.T) {
		p := Place{}
		m := &npc.Monster{}

		p.AddMonster(m)

		if !p.IsOccupied() {
			t.Error("expected a monster, but got none")
		}
	})
}

func TestPlaceRendering(t *testing.T) {
	t.Run("can render Hero from place", func(t *testing.T) {
		p := Place{}
		h := &player.Hero{}

		p.SetHero(h)

		got := p.Render()
		if strings.Compare(got, "*") != 0 {
			t.Errorf("expected to have '*' Hero rendered, got %s", got)
		}
	})

	t.Run("renders empy place when there are no Heroes", func(t *testing.T) {
		p := Place{}

		got := p.Render()
		if strings.Compare(got, "_") != 0 {
			t.Errorf("expected to have empty place '_' rendered, got %s", got)
		}
	})
}
