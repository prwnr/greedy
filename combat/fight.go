package combat

import (
	"errors"
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
		return "", errors.New("cannot attack dead monster")
	}

	heroHit := c.attacker.Attack()
	monsterHit := c.defender.Attack()
	c.defender.ReduceHP(heroHit)
	c.attacker.ReduceHP(monsterHit)

	result := fmt.Sprintf("You hit monster for %d damage, monster has %d HP left \r\n", heroHit, c.defender.GetHP())
	result += fmt.Sprintf("Monster hit you for %d damage. %d HP left \r\n", monsterHit, c.attacker.GetHP())

	return result, nil
}
