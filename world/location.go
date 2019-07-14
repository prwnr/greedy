package world

import (
	"fmt"
)

//Location of the game
type Location struct {
	Size   int
	Fields [][]string
}

//NewLocation creates new game Map with first position
func NewLocation(size int) Location {
	l := Location{Size: size}
	l.build()

	return l
}

// Display current positions on Location
func (l *Location) Display() {
	for _, s := range l.Fields {
		fmt.Println(s)
	}
}

func (l *Location) build() {
	for i := 0; i < l.Size; i++ {
		tmp := make([]string, 0)
		for j := 0; j < l.Size; j++ {
			tmp = append(tmp, "_")
		}
		l.Fields = append(l.Fields, tmp)
	}
}
