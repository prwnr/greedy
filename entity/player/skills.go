package player

import (
	"fmt"
	"swarm/common"
	"time"
)

// Skill structure
type Skill struct {
	Name     string
	CoolDown int
	Type     SkillType
}

// SkillType defines type of the skill
type SkillType int

// Skill types
const (
	Offensive = 1
	Defensive = 2
)

// starts internal skill recharge cool down
func (s *Skill) startCoolDown(cd int) {
	s.CoolDown = cd
	go func() {
		ticker := time.NewTicker(time.Second * 1)
		for {
			select {
			case <-ticker.C:
				s.CoolDown--
				if s.CoolDown == 0 {
					ticker.Stop()
					return
				}
			}
		}
	}()
}

// Name return skill name
func (s *Skill) GetName() string {
	return s.Name
}

// CurrentCoolDown returns internal recharge cool down
func (s *Skill) CurrentCoolDown() int {
	return s.CoolDown
}

// Castable defines skill
type Castable interface {
	GetName() string
	CurrentCoolDown() int
	Cast(h *Hero) Result
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
func NewHealingSkill() *HealingSkill {
	return &HealingSkill{Skill{
		Name:     "Heal",
		CoolDown: 0,
		Type:     Defensive,
	}}
}

// Cast uses a skill and starts its cool down
func (s *HealingSkill) Cast(h *Hero) Result {
	var r Result
	if s.CoolDown > 0 {
		r.Message = "Cannot use skill, still recharging."
		return r
	}

	reqMana := 11 - h.level.Number

	if h.mana <= 0 || reqMana > h.mana {
		r.Message = "Mana is too low."
		return r
	}

	if h.GetHealth() >= h.maxHealth {
		r.Message = fmt.Sprintf("Hero health restored by %d.", 0)
		return r
	}

	healAmount := 5 * h.level.Number
	h.Entity.Health += healAmount
	h.mana -= reqMana

	r.Message = fmt.Sprintf("Hero health restored by %d.", healAmount)

	r.Type = s.Type
	s.startCoolDown(5)
	return r
}

// BasicAttackSkill
type BasicAttackSkill struct {
	Skill
}

// NewBasicAttackSkill creates healing skill.
func NewBasicAttackSkill() *BasicAttackSkill {
	return &BasicAttackSkill{Skill{
		Name:     "Heal",
		CoolDown: 0,
		Type:     Offensive,
	}}
}

// Cast uses a skill and starts its cool down
func (s *BasicAttackSkill) Cast(h *Hero) Result {
	r := Result{Type: s.Type}
	r.Power = common.RandomMinNumber(h.Entity.AttackPower()-5, h.Entity.AttackPower())
	s.startCoolDown(1)

	return r
}
