package player

import (
	"fmt"
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
	Position   Position
	level      *Level
	experience int
	mana       int

	maxHealth int
	maxMana   int
}

// NewHero creates new hero struct
func NewHero(x, y int) *Hero {
	h := &Hero{
		level:     NewLevel(1, 5),
		mana:      BaseMana,
		maxHealth: BaseHealth,
		maxMana:   BaseMana,
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

	if h.GetHealth() == h.maxHealth {
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
	var reqExp int
	if !h.MaxLevel() {
		reqExp = h.level.Next.ReqExperience
	} else {
		reqExp = h.experience
	}

	return [][]string{
		[]string{"Level", fmt.Sprintf("%d", h.level.Number)},
		[]string{"Health", fmt.Sprintf("%d/%d", h.Entity.Health, h.maxHealth)},
		[]string{"Mana", fmt.Sprintf("%d/%d", h.mana, h.maxMana)},
		[]string{"Attack", fmt.Sprintf("%d", h.Entity.Attack)},
		[]string{"Experience", fmt.Sprintf("%d/%d", h.experience, reqExp)},
	}
}

// GainExperience adds exp to hero.
// Call LevelUp method once required experience is met.
func (h *Hero) GainExperience(amount int) string {
	if h.MaxLevel() {
		return ""
	}

	h.experience += amount
	if h.experience >= h.level.Next.ReqExperience {
		h.levelUp()
	}

	return fmt.Sprintf("Gained %d experience.\r\n", amount)
}

// MaxLevel determines if hero reached his maximum possible level.
func (h *Hero) MaxLevel() bool {
	if h.level.Next != nil {
		return false
	}

	return true
}

func (h *Hero) levelUp() {
	if h.MaxLevel() {
		return
	}

	h.level = h.level.Next
	h.Attack = BaseAttack + h.level.Number*2
	h.maxHealth = BaseHealth + h.level.Number*50
	h.Health = h.maxHealth
	h.maxMana = BaseMana + h.level.Number*20
	h.mana = h.maxMana
}

// Position of a hero
type Position struct {
	X int
	Y int
}
