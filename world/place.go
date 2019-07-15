package world

import "swarm/player"

// Place is a single element on a Location
type Place struct {
	hero *player.Hero
}

// SetHero puts a Hero on a Place
func (p *Place) SetHero(h *player.Hero) {
	p.hero = h
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
	if p.hero == nil {
		return "_"
	}
	return p.hero.RenderLook()
}
