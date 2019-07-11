package hero

// Bee a newborn hero
type Bee struct {
	Position Position
}

// Start sets starting position of a hero
func (b *Bee) Start(x, y int) {
	b.Position.X = x
	b.Position.Y = y
}

// Position of a hero
type Position struct {
	X int
	Y int
}
