package world

import (
	"fmt"
	"swarm/combat"
	"swarm/player"
)

const (
	//MoveUp key
	MoveUp = "w"
	//MoveDown key
	MoveDown = "s"
	//MoveLeft key
	MoveLeft = "a"
	//MoveRight key
	MoveRight = "d"
)

// Move changes hero position
func Move(h *player.Hero, l *Location, direction string) {
	l.Places[h.Position.Y][h.Position.X].RemoveHero()

	switch direction {
	case MoveUp:
		if h.Position.Y > 0 {
			h.Position.Y--
		}
	case MoveDown:
		if h.Position.Y < l.Size-1 {
			h.Position.Y++
		}
	case MoveLeft:
		if h.Position.X > 0 {
			h.Position.X--
		}
	case MoveRight:
		if h.Position.X < l.Size-1 {
			h.Position.X++
		}
	}

	p := &l.Places[h.Position.Y][h.Position.X]
	p.SetHero(h)
	if p.IsOccupied() {
		err := combat.Fight(p.GetHero(), p.GetMonster())
		if err != nil {
			fmt.Errorf("Fight error: %v", err)
		}
	}
}
