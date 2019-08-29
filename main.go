package main

import (
	"log"
	"swarm/entity"
	"swarm/game"
	"swarm/view"
	"time"

	ui "github.com/gizak/termui/v3"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	g := game.NewGame()

	end := make(chan bool)
	go func() {
		for {
			select {
			case <-view.UIChange:
				ui.Render(g.View.All()...)
			case <-entity.RechargeSkill:
				g.View.UpdateSkillBar(g.Hero.Skills())
				ui.Render(g.View.SkillsBar)
			case <-end:
				return
			}
		}
	}()

	g.InitViews()
	uiEvents := ui.PollEvents()

	tick := time.NewTicker(time.Second * 1).C
	second := int64(0)
	for {
		if g.Over {
			e := <-uiEvents
			if e.ID == "q" || e.ID == "<C-c>" {
				end <- true
				return
			}

			if e.ID == "r" {
				g = game.NewGame()
				g.Over = false
			}

			continue
		}

		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				end <- true
				return
			default:
				g.PlayerAction(e.ID)
				g.UpdateView()
			}
		case <-tick:
			second++
			g.Cycle(second)
			g.UpdateView()
		}
	}
}
