package player

import "swarm/modifiers"

// Level of a hero/monster
type Level struct {
	Number        int
	ReqExperience int
	Next          *Level
}

// NewLevel constructs level with hierarchy for next
func NewLevel(number, max int) *Level {
	var nextLevel *Level

	if number >= max {
		nextLevel = nil
	} else {
		nextLevel = NewLevel(number+1, max)
	}

	level := &Level{
		Number:        number,
		ReqExperience: modifiers.CalculateHeroLevelExperience(number),
		Next:          nextLevel,
	}

	return level
}
