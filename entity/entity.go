package entity

// Entity structure
type Entity struct {
	Health int
	Attack int
}

// AttackPower returns attack amount
func (e *Entity) AttackPower() int {
	return e.Attack
}

// ReduceHealth subtracts given amount from current health
func (e *Entity) ReduceHealth(amount int) {
	e.Health -= amount
}

// GetHealth returns current hero HP
func (e *Entity) GetHealth() int {
	return e.Health
}

// IsAlive checks if character health is not at or below 0
func (e *Entity) IsAlive() bool {
	return e.Health > 0
}
