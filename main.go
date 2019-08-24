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

	uiEvents := ui.PollEvents()

	end := make(chan bool)
	go func() {
		for {
			select {
			case <-player.RechargeSkill:
				g.View.UpdateSkillBar(g.Hero.Skills())
				ui.Render(g.View.SkillsBar)
			case <-end:
				return
			}
		}
	}()

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
				ui.Render(g.View.All()...)
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
			}
		case <-tick:
			second++
			g.TimeElapsed++
			g.UpdateGoal()

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
