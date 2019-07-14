package world

import (
	"fmt"
)

//Location of the game
type Location struct {
	Size   int
	Places [][]Place
}

//NewLocation creates new game Map with places
func NewLocation(size int) Location {
	l := Location{Size: size}
	l.build()

	return l
}

// Render current positions on Location
func (l *Location) Render() {
	loc := make([][]string, l.Size)
	for i := 0; i < l.Size; i++ {
		loc[i] = make([]string, l.Size)
	}

	for i, row := range l.Places {
		for j, el := range row {
			loc[i][j] = el.Render()
		}
	}

	for _, l := range loc {
		fmt.Println(l)
	}
}

func (l *Location) build() {
	for i := 0; i < l.Size; i++ {
		p := Place{}
		tmp := make([]Place, 0)
		for j := 0; j < l.Size; j++ {
			tmp = append(tmp, p)
		}
		l.Places = append(l.Places, tmp)
	}
}
