package main

import (
	"log"
	"swarm/game"

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
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		default:
			g.MoveHero(e.ID)
			if !g.Hero.IsAlive() {
				g.View.UpdateCombatLog("Hero died")
				gameOver = true
			}

			ui.Render(g.View.All()...)
		}

		if gameOver {
			break
		}
	}
}
