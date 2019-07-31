package combat

import (
	"swarm/character/npc"
	"swarm/character/player"
	"testing"
)

func TestHeroFightsMonster(t *testing.T) {
	t.Run("hero kills monster", func(t *testing.T) {
		m := npc.NewMonster(1)
		h := player.NewHero(0, 0)

		c := NewCombat(h, m)
		_, err := c.Fight()
		if err != nil && m.IsAlive() {
			t.Error("monster should be dead, but is still alive")
		}
		assertAttackerHealth(t, h, 95)
	})

	t.Run("hero cant kill dead monster", func(t *testing.T) {
		m := npc.NewMonster(1)
		m.ReduceHealth(100)
		h := player.NewHero(0, 0)

		c := NewCombat(h, m)
		_, err := c.Fight()

		if err == nil {
			t.Error("fight sequence should return error because monster is dead")
		}
	})
}

func TestDefenderAttacksAttackedBack(t *testing.T) {
	m := npc.NewMonster(1)
	h := player.NewHero(0, 0)

	c := NewCombat(h, m)

	_ = c.AttackBack()
	assertAttackerHealth(t, h, 95)
}

func assertAttackerHealth(t *testing.T, attacker Character, want int) {
	if attacker.GetHP() != want {
		t.Errorf("attacker should have %dHP after being hit by monster", want)
	}
}