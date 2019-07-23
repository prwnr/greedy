package main

import (
	"log"
	"swarm/player"
	"swarm/view"
	"swarm/world"

	ui "github.com/gizak/termui/v3"
)

func main() {
	size := 10
	view := view.NewView()
	h := player.NewHero()
	loc := world.NewLocation(size)

	h.StartingPosition((size/2)-1, size-1)
	world.Move(h, &loc, "init", view)

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	ui.Render(view.All()...)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		default:
			world.Move(h, &loc, e.ID, view)
			ui.Render(view.All()...)
		}
	}
}
