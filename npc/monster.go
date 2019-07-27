package npc

import "strconv"

// Monster NPC
type Monster struct {
	look   string
	hp     int
	attack int
}

// NewMonster returns new monster struct
func NewMonster() *Monster {
	m := &Monster{
		look:   "#",
		hp:     100,
		attack: 5,
	}
	return m
}

// Attack returns attack amount
func (m *Monster) Attack() int {
	return m.attack
}

// ReduceHealth subtracts given amount from current HP
func (m *Monster) ReduceHealth(amount int) {
	m.hp -= amount
}

// GetHP returns current monster HP
func (m *Monster) GetHP() int {
	return m.hp
}

// IsAlive checks if monster HP is not at or below 0
func (m *Monster) IsAlive() bool {
	return m.hp > 0
}

// Render monster look
func (m Monster) Render() string {
	return m.look
}

// GetStats returns current hero statistics
func (m *Monster) GetStats() [][]string {
	return [][]string{
		[]string{"Health", strconv.FormatInt(int64(m.hp), 10)},
		[]string{"Attack", strconv.FormatInt(int64(m.attack), 10)},
	}
}
