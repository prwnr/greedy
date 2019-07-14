package world

import "swarm/player"

// Place is a single element on a Location
type Place struct {
	avatars []player.Avatar
}

// AddAvatar puts a new Avatar on a Place
func (p *Place) AddAvatar(a player.Avatar) {
	p.avatars = append(p.avatars, a)
}

// RemoveAvatars resets Place
func (p *Place) RemoveAvatars() {
	p.avatars = []player.Avatar{}
}

// GetAvatars returns who stands on Place
func (p *Place) GetAvatars() []player.Avatar {
	return p.avatars
}

// Render Place look
func (p *Place) Render() string {
	if len(p.avatars) == 0 {
		return "_"
	}

	var place string
	for _, a := range p.avatars {
		place += a.RenderLook()
	}

	return place
}
