package npc

// Monster NPC
type Monster struct {
	look string
}

// NewMonster returns new monster struct
func NewMonster() *Monster {
	m := &Monster{look: "#"}
	return m
}

// Render monster look
func (m Monster) Render() string {
	return m.look
}
