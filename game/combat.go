package game

import (
	"errors"
	"fmt"
	"greedy/entity"
)

func fight(h *entity.Hero, m *entity.Monster, skill string) (string, error) {
	if !m.IsAlive() {
		return "", errors.New("cannot attack dead monster")
	}

	res := h.UseSkill(skill, m)

	result := res.Message
	if m.IsAlive() {
		result += fightBack(h, m)
	}

	return result, nil
}

func fightBack(h *entity.Hero, m *entity.Monster) string {
	monsterHit := m.AttackPower()
	h.ReduceHealth(monsterHit)

	return fmt.Sprintf("Monster hit you for %d damage. %d HP left \r\n", monsterHit, h.GetHealth())
}
