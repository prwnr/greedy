package combat

import (
	"fmt"
	"termui/v3/widgets"
)

var logger *widgets.Paragraph

// SetLogger for combat logs
func SetLogger(p *widgets.Paragraph) {
	logger = p
}

// Character interface defines object that has HP and can attack
type Character interface {
	IsAlive() bool
	ReduceHP(amount int)
	GetHP() int
	Attack() int
}

// Fight action between two objects
func Fight(attacker, defender Character) error {
	if !defender.IsAlive() {
		return &FightError{err: "Cannot attack dead opponent."}
	}

	nextAtt := attacker.Attack()
	defender.ReduceHP(nextAtt)
	logger.Text = fmt.Sprintf("Hitting opponent with %d power, opponent has %d HP left \r\n", nextAtt, defender.GetHP())
	return nil
}

// FightError occurs when attacker tries to attack dead opponent
type FightError struct {
	err string
}

// Error returns error string
func (e *FightError) Error() string {
	return e.err
}
