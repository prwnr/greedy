package main

import (
	"flag"
	"fmt"
	"swarm/board"
	"swarm/hero"
)

func main() {
	size := flag.Int("size", 10, "Size of each map for current game.")

	h := hero.Bee{}
	h.Start(2, 4)
	m := board.NewMap(size)
	m.Display(&h)
	var move string

	for {
		fmt.Print("Next move: ")

		_, err := fmt.Scanf("%s\n", &move)
		if err != nil {
			fmt.Println("Failed to read your move. Try again.")
			continue
		}

		h.Move(move)
		m.Display(&h)
	}
}
