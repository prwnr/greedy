package player

// Hero a newborn hero
type Hero struct {
	Position Position
	hp       int
	attack   int
}

// NewHero creates new hero struct
func NewHero() *Hero {
	h := &Hero{
		hp:     100,
		attack: 50,
	}

	return h
}

// StartingPosition sets starting position of a hero
func (h *Hero) StartingPosition(x, y int) {
	h.Position.X = x
	h.Position.Y = y
}

// Attack returns attack amount
func (h *Hero) Attack() int {
	return h.attack
}

// ReduceHP substracts given amount from current HP
func (h *Hero) ReduceHP(amount int) {
	h.hp -= amount
}

// GetHP returns current hero HP
func (h *Hero) GetHP() int {
	return h.hp
}

// IsAlive checks if monster HP is not at or below 0
func (h *Hero) IsAlive() bool {
	return h.hp > 0
}

// Render shows how hero looks like on Location
func (h Hero) Render() string {
	return "*"
}

// Position of a hero
type Position struct {
	X int
	Y int
}
