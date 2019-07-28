package npc

// Level of a monster
type Level struct {
	Number int
}

// NewLevel constructs level for a monster
func NewLevel(number int) *Level {
	return &Level{Number: number}
}
