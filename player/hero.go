package player

// Hero a newborn hero
type Hero struct {
	Position Position
}

// Start sets starting position of a hero
func (b *Hero) Start(x, y int) {
	b.Position.X = x
	b.Position.Y = y
}

// RenderLook shows how hero looks like on Location
func (b Hero) RenderLook() string {
	return "*"
}

// Position of a hero
type Position struct {
	X int
	Y int
}

// Avatar defines character look and other attributes
type Avatar interface {
	// RenderLook shows character current look
	RenderLook() string
}
