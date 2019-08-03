package main

import (
	"log"
	"swarm/entity/player"
	"swarm/game"
	"time"

	ui "github.com/gizak/termui/v3"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	g := game.NewGame()
	ui.Render(g.View.All()...)

	gameOver := false
	uiEvents := ui.PollEvents()

	tick := time.NewTicker(time.Second * 1).C
	second := int64(0)
	for {
		if gameOver {
			e := <-uiEvents
			if e.ID == "q" || e.ID == "<C-c>" {
				return
			}

			if e.ID == "r" {
				g = game.NewGame()
				ui.Render(g.View.All()...)
				gameOver = false
			}

			continue
		}

		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			default:
				g.PlayerAction(e.ID)
				if !g.Hero.IsAlive() {
					g.View.UpdateCombatLog("Hero died. Press 'q' to quit or 'r' to restart.")
					gameOver = true
				}
			}
		case <-tick:
			second++
			if second%g.Config.MonsterSpawn == 0 {
				go g.CurrentLocation.PlaceMonsters(g.Config.MonstersSpawnNum)
			}

			if second%player.RegenTimeout == 0 {
				go g.Hero.Regenerate()
			}
		}

		g.UpdateView()
		ui.Render(g.View.All()...)
	}
}
