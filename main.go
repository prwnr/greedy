package main

import (
	"log"
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
	ticker := time.NewTicker(g.Config.MonsterSpawn).C
	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			default:
				g.MoveHero(e.ID)
				if !g.Hero.IsAlive() {
					g.View.UpdateCombatLog("Hero died")
					gameOver = true
				}
			}
		case <-ticker:
			if g.CurrentLocation.HasFreePlace() {
				g.CurrentLocation.PlaceMonsters(1)
			}
		}

		g.UpdateView()
		ui.Render(g.View.All()...)

		if gameOver {
			break
		}
	}
}
