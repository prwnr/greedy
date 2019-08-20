package player

import (
	"fmt"
	"sort"
	"swarm/common"
	"swarm/entity"
)

const (
	MaxLevel = 5
	//Base attributes of the hero
	BaseHealth = 150
	BaseMana   = 60
	BaseAttack = 15
	//Modifiers
	RegenTimeout = 8
	HealthRegen  = 4
	ManaRegen    = 2
)

// Hero a newborn hero
type Hero struct {
	//Position of the hero
	Position Position
	//Current values
	entity.Entity
	level      *Level
	experience int
	mana       int
	//Maximum values
	maxHealth int
	maxMana   int
	//Skills
	skills map[string]Castable
}

// NewHero creates new hero struct
func NewHero(x, y int) *Hero {
	h := &Hero{
		level:     NewLevel(1, MaxLevel),
		mana:      BaseMana,
		maxHealth: BaseHealth,
		maxMana:   BaseMana,
	}

	h.Entity.Health = BaseHealth
	h.Entity.Attack = BaseAttack

	h.Position.X = x
	h.Position.Y = y

	h.skills = make(map[string]Castable)
	h.skills["1"] = NewBasicAttackSkill(h)
	h.skills["2"] = NewHealingSkill(h)

	return h
}

// AttackPower returns attack amount
func (h *Hero) AttackPower() int {
	return common.RandomMinNumber(h.Entity.AttackPower()-5, h.Entity.AttackPower())
}

// UseSkill selects and casts skill
func (h *Hero) UseSkill(num string, target Killable) Result {
	var res Result
	if skill, ok := h.skills[num]; ok {
		res = skill.Cast(target)
		return res
	}

	return res
}

// Regenerate restores Health and Mana
func (h *Hero) Regenerate() {
	if h.Health < h.maxHealth {
		h.Health += HealthRegen
	}

	if h.mana < h.maxMana {
		h.mana += ManaRegen
	}
}

// GainExperience adds exp to hero.
// Call LevelUp method once required experience is met.
func (h *Hero) GainExperience(amount int) string {
	if h.HasMaxLevel() {
		return ""
	}

	h.experience += amount
	if h.experience >= h.level.Next.ReqExperience {
		h.levelUp()
	}

	return fmt.Sprintf("Gained %d experience.\r\n", amount)
}

// HasMaxLevel determines if hero reached his maximum possible level.
func (h *Hero) HasMaxLevel() bool {
	if h.level.Next != nil {
		return false
	}

	return true
}

// Render shows how hero looks like on Location
func (h Hero) Render() string {
	return "*"
}

// GetStats returns current hero statistics
func (h *Hero) GetStats() [][]string {
	var reqExp int
	if !h.HasMaxLevel() {
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

// Skills return all available hero skills with their cool downs
func (h *Hero) Skills() [][]string {
	var skills [][]string
	var order []string
	for k := range h.skills {
		order = append(order, k)
	}

	sort.Strings(order)

	for _, i := range order {
		s := h.skills[i]
		skill := []string{fmt.Sprintf("%s: %s", i, s.GetName()), fmt.Sprintf("%.1f", s.CurrentCoolDown())}
		skills = append(skills, skill)
	}

	return skills
}

func (h *Hero) levelUp() {
	if h.HasMaxLevel() {
		return
	}

	h.level = h.level.Next
	h.Attack = BaseAttack + h.level.Number*2
	h.maxHealth = BaseHealth + h.level.Number*25
	h.Health = h.maxHealth
	h.maxMana = BaseMana + h.level.Number*10
	h.mana = h.maxMana
}

// Position of a hero
type Position struct {
	X int
	Y int
}
