package board

import "swarm/hero"

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
func Move(b *hero.Bee, m *Map, direction string) {
	m.Fields[b.Position.Y][b.Position.X] = "_"

	switch direction {
	case MoveUp:
		if b.Position.Y > 0 {
			b.Position.Y--
		}
	case MoveDown:
		if b.Position.Y < m.Size-1 {
			b.Position.Y++
		}
	case MoveLeft:
		if b.Position.X > 0 {
			b.Position.X--
		}
	case MoveRight:
		if b.Position.X < m.Size-1 {
			b.Position.X++
		}
	}

	m.Fields[b.Position.Y][b.Position.X] = "*"
}
