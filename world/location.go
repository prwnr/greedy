package world

import (
	"strings"
	"swarm/common"
	"swarm/npc"
)

//Location of the game
type Location struct {
	Size   int
	Places [][]Place
}

//NewLocation creates new game Map with places
func NewLocation(size int) *Location {
	l := &Location{Size: size}
	l.build()

	return l
}

// RenderPlaces on Location
func (l *Location) RenderPlaces() string {
	loc := make([][]string, l.Size)
	for i := 0; i < l.Size; i++ {
		loc[i] = make([]string, l.Size)
	}

	for i, row := range l.Places {
		for j, el := range row {
			loc[i][j] = el.Render()
		}
	}

	var render string
	for _, l := range loc {
		render += strings.Join(l, " ") + "\r\n"
	}

	return render
}

func (l *Location) build() {
	notX := 0
	notY := 0

	monsters := int(l.Size / 3)

	posX := randomUniqueNumber(l.Size, notX)
	posY := randomUniqueNumber(l.Size, notY)

	for i := 0; i < l.Size; i++ {
		tmp := make([]Place, 0)
		for j := 0; j < l.Size; j++ {
			p := Place{}
			if i == posX && j == posY {
				p.AddMonster(npc.NewMonster())
				if monsters > 0 {
					posX = randomUniqueNumber(l.Size, notX)
					posY = randomUniqueNumber(l.Size, notY)
					monsters--
				}
			}
			tmp = append(tmp, p)
		}
		l.Places = append(l.Places, tmp)
	}
}

func randomUniqueNumber(max, notIn int) int {
	n := common.RandomNumber(max)
	if n != notIn {
		return n
	}

	return randomUniqueNumber(max, notIn)
}
