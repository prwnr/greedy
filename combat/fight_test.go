package combat

import (
	"swarm/npc"
	"swarm/player"
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
