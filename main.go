package main

import (
	"log"
	"swarm/player"
	"swarm/view"
	"swarm/world"

	ui "github.com/gizak/termui/v3"
)

// Size of a single location
const Size = 10

func main() {
	view := view.NewView()
	h := player.NewHero()
	loc := world.NewLocation(Size)

	h.StartingPosition((Size/2)-1, Size-1)
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
