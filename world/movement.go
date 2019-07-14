package world

import "swarm/player"

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
func Move(b *player.Hero, l *Location, direction string) {
	l.Places[b.Position.Y][b.Position.X].RemoveAvatars()

	switch direction {
	case MoveUp:
		if b.Position.Y > 0 {
			b.Position.Y--
		}
	case MoveDown:
		if b.Position.Y < l.Size-1 {
			b.Position.Y++
		}
	case MoveLeft:
		if b.Position.X > 0 {
			b.Position.X--
		}
	case MoveRight:
		if b.Position.X < l.Size-1 {
			b.Position.X++
		}
	}

	l.Places[b.Position.Y][b.Position.X].AddAvatar(b)
}
