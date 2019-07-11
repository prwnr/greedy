package board

import (
	"fmt"
	"swarm/hero"
)

//Map of the game
type Map struct {
	Fields [][]string
}

//NewMap creates new game Map with first position
func NewMap(size int) Map {
	m := Map{}
	m.build(size)

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

func (m *Map) build(size int) {
	for i := 0; i < size; i++ {
		tmp := make([]string, 0)
		for j := 0; j < size; j++ {
			tmp = append(tmp, "_")
		}
		m.Fields = append(m.Fields, tmp)

	}
}
