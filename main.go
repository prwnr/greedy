package main

import (
	"flag"
	"log"
	"swarm/combat"
	"swarm/player"
	"swarm/world"
	"termui/v3/widgets"

	ui "github.com/gizak/termui/v3"
)

func main() {
	size := flag.Int("size", 10, "Size of each map for current game.")
	flag.Parse()

	h := player.NewHero()
	h.StartingPosition((*size/2)-1, *size-1)
	loc := world.NewLocation(*size)
	world.Move(h, &loc, "init")

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	worldLoc := loc.Create()
	p := widgets.NewParagraph()
	p.Title = "Combat log"
	p.SetRect(0, 13, 50, 18)
	combat.SetLogger(p)

	ui.Render(worldLoc, p)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		default:
			world.Move(h, &loc, e.ID)
			loc.Update()
			ui.Render(worldLoc, p)
		}
	}
}
