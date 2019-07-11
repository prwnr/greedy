package main

import (
	"flag"
	"fmt"
	"swarm/board"
	"swarm/hero"
)

func main() {
	size := flag.Int("size", 10, "Size of each map for current game.")
	flag.Parse()

	h := hero.Bee{}
	h.Start((*size/2)-1, *size-1)
	m := board.NewMap(*size)
	board.Move(&h, &m, "init")

	m.Display()
	var move string

	for {
		fmt.Print("Next move: ")

		_, err := fmt.Scanf("%s\n", &move)
		if err != nil {
			fmt.Println("Failed to read your move. Try again.")
			continue
		}

		board.Move(&h, &m, move)
		m.Display()
	}
}
