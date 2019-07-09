package hero

// Bee a newborn hero
type Bee struct {
	position Position
}

// Start sets starting position of a hero
func (b *Bee) Start(x, y int) {
	b.position.X = x
	b.position.Y = y
}

// Move changes hero position
func (b *Bee) Move(direction string) {
	switch direction {
	case "l":
		if b.position.X > 0 {
			b.position.X--
		}
	case "r":
		if b.position.X < 4 {
			b.position.X++
		}
	case "u":
		if b.position.Y > 0 {
			b.position.Y--
		}
	case "d":
		if b.position.Y < 4 {
			b.position.Y++
		}
	}
}

// GetPosition returns current position of a hero
func (b *Bee) GetPosition() Position {
	return b.position
}

// Position of a hero
type Position struct {
	X int
	Y int
}
