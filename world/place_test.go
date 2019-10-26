package world

import (
	"greedy/entity"
	"strings"
	"testing"
)

func TestPlaceHero(t *testing.T) {
	assertHasHero := func(t *testing.T, p *Place) {
		if p.GetHero() == nil {
			t.Error("expected a hero, got nil")
		}
	}

	assertNotHasHero := func(t *testing.T, p *Place) {
		if p.GetHero() != nil {
			t.Error("expected no hero, but got one")
		}
	}

	t.Run("can set new Hero on place", func(t *testing.T) {
		p := Place{}
		h := &entity.Hero{}

		p.SetHero(h)

		assertHasHero(t, &p)
	})

	t.Run("can remove all Hero from place", func(t *testing.T) {
		p := Place{}
		h := &entity.Hero{}

		p.SetHero(h)
		assertHasHero(t, &p)

		p.RemoveHero()
		assertNotHasHero(t, &p)
	})
}

func TestPlaceMonsters(t *testing.T) {
	t.Run("can add new Monster on place", func(t *testing.T) {
		p := Place{}
		m := entity.NewMonster(1)

		p.AddMonster(m)

		if !p.IsOccupied() {
			t.Error("expected a monster, but got none")
		}

		if p.GetMonster() != m {
			t.Error("monster on a place is different than expected")
		}
	})
}

func TestPlaceRendering(t *testing.T) {
	t.Run("renders Hero from place", func(t *testing.T) {
		p := Place{}
		h := entity.NewHero(0, 0)

		p.SetHero(h)

		got := p.Render()
		if strings.Compare(got, "*") != 0 {
			t.Errorf("expected to have '*' Hero rendered, got %s", got)
		}
	})

	t.Run("renders empty place when there are no Heroes", func(t *testing.T) {
		p := Place{}

		got := p.Render()
		if strings.Compare(got, "_") != 0 {
			t.Errorf("expected to have empty place '_' rendered, got %s", got)
		}
	})

	t.Run("renders Monster on place", func(t *testing.T) {
		p := Place{}
		m := entity.NewMonster(1)
		p.AddMonster(m)

		got := p.Render()
		if strings.Compare(got, "#") != 0 {
			t.Errorf("expected to have monster on place rendered, got %s", got)
		}
	})

	t.Run("renders Hero and Monster on place", func(t *testing.T) {
		p := Place{}
		m := entity.NewMonster(1)
		h := entity.NewHero(0, 0)

		p.SetHero(h)
		p.AddMonster(m)

		got := p.Render()
		if strings.Compare(got, "*#") != 0 {
			t.Errorf("expected to have hero and monster on place rendered, got %s", got)
		}
	})
}
