package entity

import (
	"fmt"
	"swarm/common"
	"sync"
	"time"
)

// skill structure
type skill struct {
	name       string
	key        string
	coolDown   float64
	internalCD float64
	hero       *Hero
	m          sync.Mutex
}

// RechargeSkill channel that triggers when skill CD is started
// and counts down
var RechargeSkill = make(chan bool)

// starts internal skill recharge cool down
func (s *skill) startCoolDown() {
	s.m.Lock()
	s.internalCD = s.coolDown
	s.m.Unlock()

	RechargeSkill <- true
	go func() {
		ticker := time.NewTicker(time.Millisecond * 500)
		for range ticker.C {
			s.m.Lock()
			s.internalCD -= 0.5
			s.m.Unlock()

			RechargeSkill <- true

			s.m.Lock()
			if s.internalCD <= 0 {
				ticker.Stop()
				s.m.Unlock()
				return
			}
			s.m.Unlock()
		}
	}()
}

// canCast skill only when internal cool down is at 0
func (s *skill) canCast() bool {
	s.m.Lock()
	defer s.m.Unlock()

	if s.internalCD > 0 {
		return false
	}

	return true
}

// name return skill name
func (s *skill) getName() string {
	return s.name
}

// key bind
func (s *skill) getKey() string {
	return s.key
}

// currentCoolDown returns internal recharge cool down
func (s *skill) currentCoolDown() float64 {
	s.m.Lock()
	defer s.m.Unlock()
	return s.internalCD
}

// castable defines skill
type castable interface {
	getName() string
	getKey() string
	currentCoolDown() float64
	cast(target killable) Result
	manaCost() int
}

// Result returns what skill did
type Result struct {
	Message string
	Power   int
}

// healingSkill structure
type healingSkill struct {
	skill
}

const healingSkillBaseAmount = 5

// newHealingSkill creates healing skill.
func newHealingSkill(h *Hero) *healingSkill {
	return &healingSkill{skill{
		name:     "Heal",
		key:      "2",
		coolDown: 4,
		hero:     h,
	}}
}

// cast uses a skill and starts its cool down
func (s *healingSkill) cast(target killable) Result {
	var r Result
	if !s.canCast() {
		r.Message = "Cannot use skill, still recharging."
		return r
	}

	if s.hero.mana <= 0 || s.manaCost() > s.hero.mana {
		r.Message = "Mana is too low."
		return r
	}

	if s.hero.GetHealth() >= s.hero.maxHealth {
		r.Message = fmt.Sprintf("hero health restored by %d.", 0)
		return r
	}

	healAmount := healingSkillBaseAmount * s.hero.level.Number
	s.hero.Health += healAmount
	s.hero.mana -= s.manaCost()

	r.Message = fmt.Sprintf("hero health restored by %d.", healAmount)

	s.startCoolDown()
	return r
}

// manaCost returns cost of the skill
func (s *healingSkill) manaCost() int {
	return 10 - s.hero.level.Number
}

// basicAttackSkill
type basicAttackSkill struct {
	skill
}

//killable contract for skills that are target health
type killable interface {
	ReduceHealth(amount int)
	GetHealth() int
}

// newBasicAttackSkill creates basic attack skill.
func newBasicAttackSkill(h *Hero) *basicAttackSkill {
	return &basicAttackSkill{skill{
		name:     "Attack",
		key:      "1",
		coolDown: 0.5,
		hero:     h,
	}}
}

// cast uses a skill and starts its cool down
func (s *basicAttackSkill) cast(target killable) Result {
	var r Result
	if !s.canCast() {
		r.Message = "Cannot use skill, still recharging."
		return r
	}
	r.Power = s.hero.AttackPower()
	if target != nil {
		target.ReduceHealth(r.Power)
		r.Message = fmt.Sprintf("You hit monster for %d damage using basic attack, monster has %d HP left \r\n", r.Power, target.GetHealth())
	}

	s.startCoolDown()

	return r
}

// manaCost returns cost of the skill
func (s *basicAttackSkill) manaCost() int {
	return 0
}

// heavyAttackSkill
type heavyAttackSkill struct {
	skill
}

const heavyAttackBasePower = 1.4

// newHeavyAttackSkill creates heavy attack skill.
func newHeavyAttackSkill(h *Hero) *heavyAttackSkill {
	return &heavyAttackSkill{skill{
		name:     "Heavy Attack",
		key:      "3",
		coolDown: 8,
		hero:     h,
	}}
}

// cast uses a skill and starts its cool down
func (s *heavyAttackSkill) cast(target killable) Result {
	var r Result
	if !s.canCast() {
		r.Message = "Cannot use skill, still recharging."
		return r
	}

	if s.hero.mana <= 0 || s.manaCost() > s.hero.mana {
		r.Message = "Mana is too low."
		return r
	}

	r.Power = common.RandomMinNumber(s.hero.AttackPower()-5, s.hero.AttackPower())
	r.Power += int(float64(r.Power) * heavyAttackBasePower)
	if target != nil {
		target.ReduceHealth(r.Power)
		r.Message = fmt.Sprintf("You hit monster for %d damage using heavy attack, monster has %d HP left \r\n", r.Power, target.GetHealth())
	}

	s.hero.mana -= s.manaCost()
	s.startCoolDown()

	return r
}

// manaCost returns cost of the skill
func (s *heavyAttackSkill) manaCost() int {
	return 12 - s.hero.level.Number
}
