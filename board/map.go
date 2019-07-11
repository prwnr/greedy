package board

import (
	"fmt"
)

//Map of the game
type Map struct {
	Size   int
	Fields [][]string
}

//NewMap creates new game Map with first position
func NewMap(size int) Map {
	m := Map{Size: size}
	m.build()

	return m
}

// Display current positions on Map
func (m *Map) Display() {
	for _, s := range m.Fields {
		fmt.Println(s)
	}
}

func (m *Map) build() {
	for i := 0; i < m.Size; i++ {
		tmp := make([]string, 0)
		for j := 0; j < m.Size; j++ {
			tmp = append(tmp, "_")
		}
		m.Fields = append(m.Fields, tmp)
	}
}
