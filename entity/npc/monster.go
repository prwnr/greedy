package npc

import (
	"strconv"
	"swarm/entity"
)

const (
	BaseHealth = 30
	BaseAttack = 5
)

var LevelLook = map[int]string{1: "#", 2: "$", 3: "@"}

// Monster NPC
type Monster struct {
	entity.Entity
	look  string
	level *Level
}

// NewMonster returns new monster struct
func NewMonster(level int) *Monster {
	l := NewLevel(level)

	m := &Monster{
		look:  LevelLook[l.Number],
		level: l,
	}

	m.Entity.Health = BaseHealth * l.Number
	m.Entity.Attack = BaseAttack

	return m
}

// Render monster look
func (m Monster) Render() string {
	return m.look
}

// GetStats returns current hero statistics
func (m *Monster) GetStats() [][]string {
	return [][]string{
		[]string{"Level", strconv.FormatInt(int64(m.level.Number), 10)},
		[]string{"Health", strconv.FormatInt(int64(m.GetHealth()), 10)},
		[]string{"AttackPower", strconv.FormatInt(int64(m.AttackPower()), 10)},
	}
}
