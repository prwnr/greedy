package player

import (
	"fmt"
	"swarm/common"
	"time"
)

// Skill structure
type Skill struct {
	Name     string
	CoolDown float64
	Type     SkillType
	Hero     *Hero
	Update   chan bool
}

// RechargeSkill channel that triggers when skill CD is started
// and counts down
var RechargeSkill = make(chan bool)

// SkillType defines type of the skill
type SkillType int

// Skill types
const (
	Offensive = 1
	Defensive = 2
)

// starts internal skill recharge cool down
func (s *Skill) startCoolDown(cd float64) {
	s.CoolDown = cd
	RechargeSkill <- true
	go func() {
		ticker := time.NewTicker(time.Millisecond * 500)
		for range ticker.C {
			s.CoolDown -= 0.5
			RechargeSkill <- true
			if s.CoolDown <= 0 {
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
	return s.CoolDown
}

// Castable defines skill
type Castable interface {
	GetName() string
	CurrentCoolDown() float64
	Cast(target Killable) Result
}

// Result returns what skill did
type Result struct {
	Message string
	Type    SkillType
	Power   int
}

// HealingSkill structure
type HealingSkill struct {
	Skill
}

// NewHealingSkill creates healing skill.
func NewHealingSkill(h *Hero) *HealingSkill {
	return &HealingSkill{Skill{
		Name:     "Heal",
		CoolDown: 0,
		Type:     Defensive,
		Hero:     h,
	}}
}

// Cast uses a skill and starts its cool down
func (s *HealingSkill) Cast(target Killable) Result {
	var r Result
	if s.CoolDown > 0 {
		r.Message = "Cannot use skill, still recharging."
		return r
	}

	reqMana := 11 - s.Hero.level.Number
	if s.Hero.mana <= 0 || reqMana > s.Hero.mana {
		r.Message = "Mana is too low."
		return r
	}

	if s.Hero.GetHealth() >= s.Hero.maxHealth {
		r.Message = fmt.Sprintf("Hero health restored by %d.", 0)
		return r
	}

	healAmount := 5 * s.Hero.level.Number
	s.Hero.Entity.Health += healAmount
	s.Hero.mana -= reqMana

	r.Message = fmt.Sprintf("Hero health restored by %d.", healAmount)

	r.Type = s.Type
	s.startCoolDown(5)
	return r
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
		CoolDown: 0,
		Type:     Offensive,
		Hero:     h,
	}}
}

// Cast uses a skill and starts its cool down
func (s *BasicAttackSkill) Cast(target Killable) Result {
	var r Result
	if s.CoolDown > 0 {
		r.Message = "Cannot use skill, still recharging."
		return r
	}
	r.Power = common.RandomMinNumber(s.Hero.Entity.AttackPower()-5, s.Hero.Entity.AttackPower())
	if target != nil {
		target.ReduceHealth(r.Power)
		r.Message = fmt.Sprintf("You hit monster for %d damage using basic attack, monster has %d HP left \r\n", r.Power, target.GetHealth())
	}

	s.startCoolDown(0.5)

	return r
}

// HeavyAttackSkill
type HeavyAttackSkill struct {
	Skill
}

// NewHeavyAttackSkill creates heavy attack skill.
func NewHeavyAttackSkill(h *Hero) *HeavyAttackSkill {
	return &HeavyAttackSkill{Skill{
		Name:     "Heavy Attack",
		CoolDown: 0,
		Type:     Offensive,
		Hero:     h,
	}}
}

// Cast uses a skill and starts its cool down
func (s *HeavyAttackSkill) Cast(target Killable) Result {
	var r Result
	if s.CoolDown > 0 {
		r.Message = "Cannot use skill, still recharging."
		return r
	}

	reqMana := 15 - s.Hero.level.Number
	if s.Hero.mana <= 0 || reqMana > s.Hero.mana {
		r.Message = "Mana is too low."
		return r
	}

	r.Power = common.RandomMinNumber(s.Hero.Entity.AttackPower()-5, s.Hero.Entity.AttackPower())
	r.Power += int(float64(r.Power) * 1.2)
	if target != nil {
		target.ReduceHealth(r.Power)
		r.Message = fmt.Sprintf("You hit monster for %d damage using heavy attack, monster has %d HP left \r\n", r.Power, target.GetHealth())
	}

	s.Hero.mana -= reqMana
	s.startCoolDown(10)

	return r
}
