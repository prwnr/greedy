package combat

import (
	"fmt"
)

// Character interface defines object that has HP and can attack
type Character interface {
	IsAlive() bool
	ReduceHP(amount int)
	GetHP() int
	Attack() int
}

// Combat struct
type Combat struct {
	attacker Character
	defender Character
}

// NewCombat creates new combat with attacker and defender
func NewCombat(attacker, defender Character) Combat {
	return Combat{
		attacker: attacker,
		defender: defender,
	}
}

// Fight action between attacker and defender
func (c Combat) Fight() (string, error) {
	if !c.defender.IsAlive() {
		return "", &FightError{err: "Cannot attack dead opponent."}
	}

	nextAtt := c.attacker.Attack()
	c.defender.ReduceHP(nextAtt)

	result := fmt.Sprintf("Hitting opponent with %d power, opponent has %d HP left \r\n", nextAtt, c.defender.GetHP())
	return result, nil
}

// FightError occurs when attacker tries to attack dead opponent
type FightError struct {
	err string
}

// Error returns error string
func (e *FightError) Error() string {
	return e.err
}
