package npc

import (
	"fmt"
	"swarm/entity"
)

const (
	BaseHealth     = 30
	BaseAttack     = 5
	BaseExperience = 10
)

var LevelLook = map[int]string{1: "#", 2: "$", 3: "@"}

// Monster NPC
type Monster struct {
	entity.Entity
	look      string
	level     *Level
	maxHealth int
}

// NewMonster returns new monster struct
func NewMonster(level int) *Monster {
	l := NewLevel(level)

	m := &Monster{
		look:  LevelLook[l.Number],
		level: l,
	}

	m.maxHealth = BaseHealth*l.Number + (int(float64(l.Number-1) * 5))
	m.Entity.Health = m.maxHealth
	m.Entity.Attack = BaseAttack * l.Number

	return m
}

// GetExperienceValue returns how much experience monster is worth.
func (m *Monster) GetExperienceValue() int {
	return BaseExperience * m.level.Number
}

// Render monster look
func (m Monster) Render() string {
	return m.look
}

// GetStats returns current hero statistics
func (m *Monster) GetStats() [][]string {
	return [][]string{
		[]string{"Level", fmt.Sprintf("%d", m.level.Number)},
		[]string{"Health", fmt.Sprintf("%d/%d", m.GetHealth(), m.maxHealth)},
		[]string{"Attack", fmt.Sprintf("%d", m.AttackPower())},
	}
}
