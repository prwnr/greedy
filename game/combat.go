package game

import (
	"errors"
	"fmt"
	"swarm/entity/npc"
	"swarm/entity/player"
)

func fight(h *player.Hero, m *npc.Monster, skill string) (string, error) {
	if !m.IsAlive() {
		return "", errors.New("cannot attack dead monster")
	}

	res := h.UseSkill(skill, m)

	result := res.Message
	result += fightBack(h, m)

	return result, nil
}

func fightBack(h *player.Hero, m *npc.Monster) string {
	monsterHit := m.AttackPower()
	h.ReduceHealth(monsterHit)

	return fmt.Sprintf("Monster hit you for %d damage. %d HP left \r\n", monsterHit, h.GetHealth())
}
