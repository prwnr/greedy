package world

import (
	"swarm/npc"
	"swarm/player"
)

// Place is a single element on a Location
type Place struct {
	hero    *player.Hero
	monster *npc.Monster
}

// SetHero puts a Hero on a Place
func (p *Place) SetHero(h *player.Hero) {
	p.hero = h
}

// AddMonster to a place
func (p *Place) AddMonster(m *npc.Monster) {
	p.monster = m
}

// IsOccupied checks if place is occupied by monsters
func (p *Place) IsOccupied() bool {
	return p.monster != nil
}

// RemoveHero removes Hero from current place
func (p *Place) RemoveHero() {
	p.hero = nil
}

// GetHero returns stands on Place
func (p *Place) GetHero() *player.Hero {
	return p.hero
}

// GetMonster returns monster from the place
func (p *Place) GetMonster() *npc.Monster {
	if p.monster == nil {
		return nil
	}

	return p.monster
}

// Render Place look
func (p *Place) Render() string {
	var look string
	if p.hero != nil {
		look += p.hero.Render()
	}

	if p.monster != nil && p.monster.IsAlive() {
		look += p.monster.Render()
	}

	if look != "" {
		return look
	}

	return "_"
}
