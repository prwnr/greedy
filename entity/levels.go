package entity

import "greedy/modifiers"

// HeroLevel of a hero/monster
type HeroLevel struct {
	Number        int
	ReqExperience int
	Next          *HeroLevel
}

// NewHeroLevel constructs level with hierarchy for next
func NewHeroLevel(number, max int) *HeroLevel {
	var nextLevel *HeroLevel

	if number >= max {
		nextLevel = nil
	} else {
		nextLevel = NewHeroLevel(number+1, max)
	}

	level := &HeroLevel{
		Number:        number,
		ReqExperience: modifiers.CalculateHeroLevelExperience(number),
		Next:          nextLevel,
	}

	return level
}

// MonsterLevel of a monster
type MonsterLevel struct {
	Number int
}

// NewMonsterLevel constructs level for a monster
func NewMonsterLevel(number int) *MonsterLevel {
	return &MonsterLevel{Number: number}
}
