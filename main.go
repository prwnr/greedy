package main

import (
	"flag"
	"fmt"
	"swarm/player"
	"swarm/world"
)

func main() {
	size := flag.Int("size", 10, "Size of each map for current game.")
	flag.Parse()

	h := player.NewHero()
	h.StartingPosition((*size/2)-1, *size-1)
	l := world.NewLocation(*size)
	world.Move(h, &l, "init")

	l.Render()
	var move string

	for {
		fmt.Print("Next move: ")

		_, err := fmt.Scanf("%s\n", &move)
		if err != nil {
			fmt.Println("Failed to read your move. Try again.")
			continue
		}

		world.Move(h, &l, move)
		l.Render()
	}
}
