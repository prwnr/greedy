package combat

import (
	"fmt"
	"swarm/npc"
	"swarm/player"
	"testing"
)

func TestHeroFightsMonster(t *testing.T) {
	t.Run("hero kills monster", func(t *testing.T) {
		m := npc.NewMonster()
		h := player.NewHero()

		err := Fight(h, m)
		if err != nil && m.IsAlive() {
			t.Error("monster should be dead, but is still alive")
		}
	})

	t.Run("hero cant kill dead monster", func(t *testing.T) {
		m := npc.NewMonster()
		m.ReduceHP(100)
		h := player.NewHero()

		err := Fight(h, m)
		fmt.Errorf("error %v", err)

		if err == nil {
			t.Error("fight sequence should return error because monster is dead")
		}
	})
}
