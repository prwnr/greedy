package npc

// Monster NPC
type Monster struct {
	look string
	hp   int
}

// NewMonster returns new monster struct
func NewMonster() *Monster {
	m := &Monster{look: "#", hp: 100}
	return m
}

// Attack returns attack amount
func (m *Monster) Attack() int {
	return 0
}

// ReduceHP substracts given amount from current HP
func (m *Monster) ReduceHP(amount int) {
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
