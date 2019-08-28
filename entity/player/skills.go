package player

import (
	"fmt"
	"swarm/common"
	"time"
)

// Skill structure
type Skill struct {
	//Name of the skill displayed in UI
	Name string
	//CoolDown defined skill CD
	CoolDown float64
	//internalCD as current cool down after skill use
	internalCD float64
	//Hero which used the skill
	Hero *Hero
}

// RechargeSkill channel that triggers when skill CD is started
// and counts down
var RechargeSkill = make(chan bool)

// starts internal skill recharge cool down
func (s *Skill) startCoolDown() {
	s.internalCD = s.CoolDown
	RechargeSkill <- true
	go func() {
		ticker := time.NewTicker(time.Millisecond * 500)
		for range ticker.C {
			s.internalCD -= 0.5
			RechargeSkill <- true
			if s.internalCD <= 0 {
				ticker.Stop()
				return
			}
		}
	}()
}

// Name return skill name
func (s *Skill) GetName() string {
	return s.Name
}

// CurrentCoolDown returns internal recharge cool down
func (s *Skill) CurrentCoolDown() float64 {
	return s.internalCD
}

// Castable defines skill
type Castable interface {
	GetName() string
	CurrentCoolDown() float64
	Cast(target Killable) Result
	ManaCost() int
}

// Result returns what skill did
type Result struct {
	Message string
	Power   int
}

// HealingSkill structure
type HealingSkill struct {
	Skill
}

const healingSkillBaseAmount = 5

// NewHealingSkill creates healing skill.
func NewHealingSkill(h *Hero) *HealingSkill {
	return &HealingSkill{Skill{
		Name:     "Heal",
		CoolDown: 4,
		Hero:     h,
	}}
}

// Cast uses a skill and starts its cool down
func (s *HealingSkill) Cast(target Killable) Result {
	var r Result
	if s.internalCD > 0 {
		r.Message = "Cannot use skill, still recharging."
		return r
	}

	if s.Hero.mana <= 0 || s.ManaCost() > s.Hero.mana {
		r.Message = "Mana is too low."
		return r
	}

	if s.Hero.GetHealth() >= s.Hero.maxHealth {
		r.Message = fmt.Sprintf("Hero health restored by %d.", 0)
		return r
	}

	healAmount := healingSkillBaseAmount * s.Hero.level.Number
	s.Hero.Entity.Health += healAmount
	s.Hero.mana -= s.ManaCost()

	r.Message = fmt.Sprintf("Hero health restored by %d.", healAmount)

	s.startCoolDown()
	return r
}

// ManaCost returns cost of the skill
func (s *HealingSkill) ManaCost() int {
	return 10 - s.Hero.level.Number
}

// BasicAttackSkill
type BasicAttackSkill struct {
	Skill
}

type Killable interface {
	ReduceHealth(amount int)
	GetHealth() int
}

// NewBasicAttackSkill creates basic attack skill.
func NewBasicAttackSkill(h *Hero) *BasicAttackSkill {
	return &BasicAttackSkill{Skill{
		Name:     "Attack",
		CoolDown: 0.5,
		Hero:     h,
	}}
}

// Cast uses a skill and starts its cool down
func (s *BasicAttackSkill) Cast(target Killable) Result {
	var r Result
	if s.internalCD > 0 {
		r.Message = "Cannot use skill, still recharging."
		return r
	}
	r.Power = common.RandomMinNumber(s.Hero.Entity.AttackPower()-5, s.Hero.Entity.AttackPower())
	if target != nil {
		target.ReduceHealth(r.Power)
		r.Message = fmt.Sprintf("You hit monster for %d damage using basic attack, monster has %d HP left \r\n", r.Power, target.GetHealth())
	}

	s.startCoolDown()

	return r
}

// ManaCost returns cost of the skill
func (s *BasicAttackSkill) ManaCost() int {
	return 0
}

// HeavyAttackSkill
type HeavyAttackSkill struct {
	Skill
}

const heavyAttackBasePower = 1.4

// NewHeavyAttackSkill creates heavy attack skill.
func NewHeavyAttackSkill(h *Hero) *HeavyAttackSkill {
	return &HeavyAttackSkill{Skill{
		Name:     "Heavy Attack",
		CoolDown: 8,
		Hero:     h,
	}}
}

// Cast uses a skill and starts its cool down
func (s *HeavyAttackSkill) Cast(target Killable) Result {
	var r Result
	if s.internalCD > 0 {
		r.Message = "Cannot use skill, still recharging."
		return r
	}

	if s.Hero.mana <= 0 || s.ManaCost() > s.Hero.mana {
		r.Message = "Mana is too low."
		return r
	}

	r.Power = common.RandomMinNumber(s.Hero.Entity.AttackPower()-5, s.Hero.Entity.AttackPower())
	r.Power += int(float64(r.Power) * heavyAttackBasePower)
	if target != nil {
		target.ReduceHealth(r.Power)
		r.Message = fmt.Sprintf("You hit monster for %d damage using heavy attack, monster has %d HP left \r\n", r.Power, target.GetHealth())
	}

	s.Hero.mana -= s.ManaCost()
	s.startCoolDown()

	return r
}

// ManaCost returns cost of the skill
func (s *HeavyAttackSkill) ManaCost() int {
	return 12 - s.Hero.level.Number
}
