package npc

import (
	"fmt"
	"swarm/entity"
	"swarm/modifiers"
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

	m.maxHealth = modifiers.CalculateMonsterHealth(m.Level())
	m.Entity.Health = m.maxHealth
	m.Entity.Attack = modifiers.CalculateMonsterAttack(m.Level())

	return m
}

// SetLook for monster
func (m *Monster) SetLook(look string) {
	m.look = look
}

// Level of the monster
func (m *Monster) Level() int {
	return m.level.Number
}

// GetExperienceValue returns how much experience monster is worth.
func (m *Monster) GetExperienceValue() int {
	return modifiers.CalculateMonsterExperience(m.Level())
}

// Render monster look
func (m Monster) Render() string {
	return m.look
}

// GetStats returns current hero statistics
func (m *Monster) GetStats() [][]string {
	return [][]string{
		{"Level", fmt.Sprintf("%d", m.level.Number)},
		{"Health", fmt.Sprintf("%d/%d", m.GetHealth(), m.maxHealth)},
		{"Attack", fmt.Sprintf("%d", m.AttackPower())},
	}
}
