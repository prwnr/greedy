package world

import (
	"swarm/npc"
	"swarm/player"
)

// Place is a single element on a Location
type Place struct {
	hero     *player.Hero
	monsters []npc.Monster
}

// SetHero puts a Hero on a Place
func (p *Place) SetHero(h *player.Hero) {
	p.hero = h
}

// AddMonster to a place
func (p *Place) AddMonster(m *npc.Monster) {
	p.monsters = append(p.monsters, *m)
}

// IsOccupied checks if place is occupied by monsters
func (p *Place) IsOccupied() bool {
	return len(p.monsters) > 0
}

// RemoveHero removes Hero from current place
func (p *Place) RemoveHero() {
	p.hero = nil
}

// GetHero returns stands on Place
func (p *Place) GetHero() *player.Hero {
	return p.hero
}

// Render Place look
func (p *Place) Render() string {
	var look string
	if p.hero != nil {
		look += p.hero.Render()
	}

	if len(p.monsters) > 0 {
		for _, m := range p.monsters {
			look += m.Render()
		}
	}

	if look != "" {
		return look
	}

	return "_"
}
