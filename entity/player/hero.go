package player

import (
	"fmt"
	"strconv"
	"swarm/common"
	"swarm/entity"
)

const (
	BaseHealth = 100
	BaseMana   = 50
	BaseAttack = 15
)

// Hero a newborn hero
type Hero struct {
	entity.Entity
	Position Position
	level    *Level
	mana     int
}

// NewHero creates new hero struct
func NewHero(x, y int) *Hero {
	h := &Hero{
		level: NewLevel(1, 5),
		mana:  BaseMana,
	}

	h.Entity.Health = BaseHealth
	h.Entity.Attack = BaseAttack

	h.Position.X = x
	h.Position.Y = y

	return h
}

// AttackPower returns attack amount
func (h *Hero) AttackPower() int {
	return common.RandomMinNumber(h.Entity.AttackPower()-5, h.Entity.AttackPower())
}

// UseHeal activates given hero skill
func (h *Hero) UseHeal() string {
	if h.mana <= 0 {
		return fmt.Sprint("Mana is too low.")
	}

	if h.GetHealth() == 100 {
		return fmt.Sprintf("Hero health restored by %d.", 0)
	}

	healAmount := 5 * h.level.Number
	h.Entity.Health += healAmount
	h.mana -= 11 - h.level.Number

	return fmt.Sprintf("Hero health restored by %d.", healAmount)
}

// Render shows how hero looks like on Location
func (h Hero) Render() string {
	return "*"
}

// GetStats returns current hero statistics
func (h *Hero) GetStats() [][]string {
	return [][]string{
		[]string{"Level", strconv.FormatInt(int64(h.level.Number), 10)},
		[]string{"Health", strconv.FormatInt(int64(h.Entity.Health), 10)},
		[]string{"Mana", strconv.FormatInt(int64(h.mana), 10)},
		[]string{"Attack", strconv.FormatInt(int64(h.Entity.Attack), 10)},
	}
}

// Position of a hero
type Position struct {
	X int
	Y int
}
