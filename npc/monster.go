package npc

import (
	"strconv"
	"swarm/common"
	"swarm/player"
)

// Monster NPC
type Monster struct {
	look   string
	health int
	attack int
	level  *player.Level
}

// NewMonster returns new monster struct
func NewMonster() *Monster {
	bHealth := 30
	l := player.NewLevel(common.RandomNumber(3)+1, 1)

	m := &Monster{
		look:   "#",
		health: bHealth * l.Number,
		attack: 5,
		level:  l,
	}
	return m
}

// Attack returns attack amount
func (m *Monster) Attack() int {
	return m.attack * m.level.Number
}

// ReduceHealth subtracts given amount from current HP
func (m *Monster) ReduceHealth(amount int) {
	m.health -= amount
}

// GetHP returns current monster HP
func (m *Monster) GetHP() int {
	return m.health
}

// IsAlive checks if monster HP is not at or below 0
func (m *Monster) IsAlive() bool {
	return m.health > 0
}

// Render monster look
func (m Monster) Render() string {
	return m.look
}

// GetStats returns current hero statistics
func (m *Monster) GetStats() [][]string {
	return [][]string{
		[]string{"Level", strconv.FormatInt(int64(m.level.Number), 10)},
		[]string{"Health", strconv.FormatInt(int64(m.GetHP()), 10)},
		[]string{"Attack", strconv.FormatInt(int64(m.Attack()), 10)},
	}
}
