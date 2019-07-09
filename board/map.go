package board

import (
	"fmt"
	"swarm/hero"
)

//Map of the game
type Map struct {
	Fields [5][5]string
}

//NewMap creates new game Map with first position
func NewMap() Map {
	m := Map{}
	m.Fields = [5][5]string{
		{"_", "_", "_", "_", "_"},
		{"_", "_", "_", "_", "_"},
		{"_", "_", "_", "_", "_"},
		{"_", "_", "_", "_", "_"},
		{"_", "_", "_", "_", "_"},
	}

	return m
}

// Display current positions on Map
func (m *Map) Display(b *hero.Bee) {
	prevPos := hero.Position{}
	for i, x := range m.Fields {
		for j, y := range x {
			if y == "*" {
				prevPos.X = j
				prevPos.Y = i
			}
		}
	}

	currPos := b.GetPosition()

	m.Fields[prevPos.Y][prevPos.X] = "_"
	m.Fields[currPos.Y][currPos.X] = "*"

	for _, s := range m.Fields {
		fmt.Println(s)
	}
}
