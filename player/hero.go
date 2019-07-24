package player

import "strconv"

// Hero a newborn hero
type Hero struct {
	Position Position
	hp       int
	attack   int
}

// NewHero creates new hero struct
func NewHero(x, y int) *Hero {
	h := &Hero{
		hp:     100,
		attack: 50,
	}

	h.Position.X = x
	h.Position.Y = y

	return h
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

// GetStats returns current hero statistics
func (h *Hero) GetStats() [][]string {
	return [][]string{
		[]string{"Health", strconv.FormatInt(int64(h.hp), 10)},
		[]string{"Attack", strconv.FormatInt(int64(h.attack), 10)},
	}
}

// Position of a hero
type Position struct {
	X int
	Y int
}
