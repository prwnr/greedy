package player

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
		ReqExperience: calculateLevelExperience(number),
		Next:          nextLevel,
	}

	return level
}

func calculateLevelExperience(num int) int {
	i := (num - 1) * 100

	return i + int(float64(i)*0.2)*(num-1)
}
