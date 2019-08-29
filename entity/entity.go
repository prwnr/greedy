package entity

import (
	"fmt"
	"sort"
	"swarm/common"
	"swarm/modifiers"
)

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

// hero a newborn hero
type Hero struct {
	//Position of the hero
	Position Position
	//Current values
	Entity
	level      *HeroLevel
	experience int
	mana       int
	//Maximum values
	maxHealth int
	maxMana   int
	//Skills
	skills map[string]castable
}

// NewHero creates new hero struct
func NewHero(x, y int) *Hero {
	h := &Hero{
		level:     NewHeroLevel(1, modifiers.HeroMaxLevel),
		mana:      modifiers.HeroBaseMana,
		maxHealth: modifiers.HeroBaseHealth,
		maxMana:   modifiers.HeroBaseMana,
	}

	h.Entity.Health = modifiers.HeroBaseHealth
	h.Entity.Attack = modifiers.HeroBaseAttack

	h.Position.X = x
	h.Position.Y = y

	h.skills = make(map[string]castable)
	h.skills["1"] = newBasicAttackSkill(h)
	h.skills["2"] = newHealingSkill(h)
	h.skills["3"] = newHeavyAttackSkill(h)

	return h
}

// AttackPower returns attack amount
func (h *Hero) AttackPower() int {
	return common.RandomMinNumber(h.Entity.AttackPower()-3, h.Entity.AttackPower())
}

// UseSkill selects and casts skill
func (h *Hero) UseSkill(num string, target Killable) Result {
	var res Result
	if skill, ok := h.skills[num]; ok {
		res = skill.cast(target)
		return res
	}

	return res
}

// Regenerate restores Health and Mana
func (h *Hero) Regenerate() {
	if h.Health < h.maxHealth {
		h.Health += modifiers.HeroHealthRegen
	}

	if h.mana < h.maxMana {
		h.mana += modifiers.HeroManaRegen
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
		{"Level", fmt.Sprintf("%d", h.level.Number)},
		{"Health", fmt.Sprintf("%d/%d", h.Entity.Health, h.maxHealth)},
		{"Mana", fmt.Sprintf("%d/%d", h.mana, h.maxMana)},
		{"Attack", fmt.Sprintf("%d", h.Entity.Attack)},
		{"Experience", fmt.Sprintf("%d/%d", h.experience, reqExp)},
	}
}

// Skills returns all available hero skills with their cool downs and mana cost
func (h *Hero) Skills() [][]string {
	var order []string
	for k := range h.skills {
		order = append(order, k)
	}

	sort.Strings(order)

	var names, cds, mana []string
	for _, i := range order {
		s := h.skills[i]
		names = append(names, fmt.Sprintf("%s:%s", i, s.getName()))
		mana = append(mana, fmt.Sprintf("%d mana", s.manaCost()))
		cds = append(cds, fmt.Sprintf("%.1f", s.currentCoolDown()))
	}

	return [][]string{
		names, mana, cds,
	}
}

func (h *Hero) levelUp() {
	if h.HasMaxLevel() {
		return
	}

	h.level = h.level.Next
	h.Attack = modifiers.CalculateHeroAttack(h.level.Number)
	h.maxHealth = modifiers.CalculateHeroHealth(h.level.Number)
	h.Health = h.maxHealth
	h.maxMana = modifiers.CalculateHeroMana(h.level.Number)
	h.mana = h.maxMana
}

// Position of a hero
type Position struct {
	X int
	Y int
}

var LevelLook = map[int]string{1: "#", 2: "$", 3: "@"}

// Monster NPC
type Monster struct {
	Entity
	look      string
	level     *MonsterLevel
	maxHealth int
}

// NewMonster returns new monster struct
func NewMonster(level int) *Monster {
	l := NewMonsterLevel(level)

	m := &Monster{
		look:  LevelLook[l.Number],
		level: l,
	}

	m.maxHealth = modifiers.CalculateMonsterHealth(m.Level())
	m.Entity.Health = m.maxHealth
	m.Entity.Attack = modifiers.CalculateMonsterAttack(m.Level())

	return m
}

// SetLook for monster
func (m *Monster) SetLook(look string) {
	m.look = look
}

// MonsterLevel of the monster
func (m *Monster) Level() int {
	return m.level.Number
}

// GetExperienceValue returns how much experience monster is worth.
func (m *Monster) GetExperienceValue() int {
	return modifiers.CalculateMonsterExperience(m.Level())
}

// Render monster look
func (m Monster) Render() string {
	return m.look
}

// GetStats returns current hero statistics
func (m *Monster) GetStats() [][]string {
	return [][]string{
		{"Level", fmt.Sprintf("%d", m.level.Number)},
		{"Health", fmt.Sprintf("%d/%d", m.GetHealth(), m.maxHealth)},
		{"Attack", fmt.Sprintf("%d", m.AttackPower())},
	}
}
