package world

import (
	"greedy/entity"
)

// Place is a single element on a Location
type Place struct {
	hero    *entity.Hero
	monster *entity.Monster
}

// SetHero puts a Hero on a Place
func (p *Place) SetHero(h *entity.Hero) {
	p.hero = h
}

// AddMonster to a place
func (p *Place) AddMonster(m *entity.Monster) {
	p.monster = m
}

// IsOccupied checks if place is occupied by monsters
func (p *Place) IsOccupied() bool {
	if p.monster == nil {
		return false
	}

	return p.monster.IsAlive()
}

// RemoveHero removes Hero from current place
func (p *Place) RemoveHero() {
	p.hero = nil
}

// RemoveMonster removes monster from current place
func (p *Place) RemoveMonster() {
	p.monster = nil
}

// GetHero returns stands on Place
func (p *Place) GetHero() *entity.Hero {
	return p.hero
}

// GetMonster returns monster from the place
func (p *Place) GetMonster() *entity.Monster {
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
