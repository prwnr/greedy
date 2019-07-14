package main

import (
	"flag"
	"fmt"
	"swarm/hero"
	"swarm/world"
)

func main() {
	size := flag.Int("size", 10, "Size of each map for current game.")
	flag.Parse()

	h := hero.Bee{}
	h.Start((*size/2)-1, *size-1)
	l := world.NewLocation(*size)
	world.Move(&h, &l, "init")

	l.Display()
	var move string

	for {
		fmt.Print("Next move: ")

		_, err := fmt.Scanf("%s\n", &move)
		if err != nil {
			fmt.Println("Failed to read your move. Try again.")
			continue
		}

		world.Move(&h, &l, move)
		l.Display()
	}
}
