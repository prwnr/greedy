package entity

import (
	"errors"
	"fmt"
)

// Fightable interface defines object that can fight in combats.
type Fightable interface {
	IsAlive() bool
	ReduceHealth(amount int)
	GetHealth() int
	AttackPower() int
}

// Combat struct
type Combat struct {
	attacker Fightable
	defender Fightable
}

// NewCombat creates new combat with attacker and defender
func NewCombat(attacker, defender Fightable) Combat {
	return Combat{
		attacker: attacker,
		defender: defender,
	}
}

// Fight action between attacker and defender
func (c Combat) Fight() (string, error) {
	if !c.defender.IsAlive() {
		return "", errors.New("cannot attack dead monster")
	}

	heroHit := c.attacker.AttackPower()
	c.defender.ReduceHealth(heroHit)

	result := fmt.Sprintf("You hit monster for %d damage, monster has %d HP left \r\n", heroHit, c.defender.GetHealth())
	result += c.AttackBack()

	return result, nil
}

// AttackBack is an action where defender hits attacker
func (c Combat) AttackBack() string {
	monsterHit := c.defender.AttackPower()
	c.attacker.ReduceHealth(monsterHit)

	result := fmt.Sprintf("Monster hit you for %d damage. %d HP left \r\n", monsterHit, c.attacker.GetHealth())

	return result
}
