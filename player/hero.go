package player

import (
	"fmt"
	"strconv"
	"swarm/common"
)

const (
	BaseHealth = 100
	BaseMana   = 50
	BaseAttack = 15
)

// Hero a newborn hero
type Hero struct {
	Position Position
	level    *Level
	health   int
	mana     int
	attack   int
}

// NewHero creates new hero struct
func NewHero(x, y int) *Hero {
	h := &Hero{
		level:  NewLevel(1, 5),
		health: BaseHealth,
		mana:   BaseMana,
		attack: BaseAttack,
	}

	h.Position.X = x
	h.Position.Y = y

	return h
}

// Attack returns attack amount
func (h *Hero) Attack() int {
	return common.RandomMinNumber(h.attack-5, h.attack)
}

// ReduceHealth subtracts given amount from current HP
func (h *Hero) ReduceHealth(amount int) {
	h.health -= amount
}

// GetHP returns current hero HP
func (h *Hero) GetHP() int {
	return h.health
}

// IsAlive checks if monster HP is not at or below 0
func (h *Hero) IsAlive() bool {
	return h.health > 0
}

// UseHeal activates given hero skill
func (h *Hero) UseHeal() string {
	if h.mana <= 0 {
		return fmt.Sprint("Mana is too low.")
	}

	if h.health == 100 {
		return fmt.Sprintf("Hero health restored by %d.", 0)
	}

	healAmount := 5 * h.level.Number
	h.health += healAmount
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
		[]string{"Health", strconv.FormatInt(int64(h.health), 10)},
		[]string{"Mana", strconv.FormatInt(int64(h.mana), 10)},
		[]string{"Attack", strconv.FormatInt(int64(h.attack), 10)},
	}
}

// Position of a hero
type Position struct {
	X int
	Y int
}
