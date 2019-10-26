package main

import (
	"greedy/entity"
	"greedy/game"
	"log"
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
			case <-entity.RechargeSkill:
				g.RunLocked(func() {
					g.View.UpdateSkillBar(g.Hero.Skills())
				})
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
				g.RunLocked(func() {
					g.Reset()
					g.UpdateView()
				})
			}

			continue
		}

		if g.GreedsReleased {
			g.RunLocked(func() {
				if !g.CurrentLocation.HasFreePlace() {
					g.EndGame("Greeds released. You lost.")
				}

				g.UpdateView()
			})

			continue
		}

		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				end <- true
				return
			default:
				g.RunLocked(func() {
					g.PlayerAction(e.ID)
					g.UpdateView()
				})
			}
		case <-tick:
			second++
			g.RunLocked(func() {
				g.Cycle(second)
				g.UpdateView()
			})
		}
	}
}
