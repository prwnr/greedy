package world

import (
	"fmt"
	"math/rand"
	"swarm/npc"
	"time"
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
	notX := (l.Size / 2) - 1
	notY := l.Size - 1

	monsters := int(l.Size / 3)

	posX := randomNumber(l.Size, notX)
	posY := randomNumber(l.Size, notY)

	for i := 0; i < l.Size; i++ {
		tmp := make([]Place, 0)
		for j := 0; j < l.Size; j++ {
			p := Place{}
			if i == posX && j == posY {
				p.AddMonster(npc.NewMonster())
				if monsters > 0 {
					posX = randomNumber(l.Size, notX)
					posY = randomNumber(l.Size, notY)
					monsters--
				}
			}
			tmp = append(tmp, p)
		}
		l.Places = append(l.Places, tmp)
	}
}

func randomNumber(max, not int) int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	n := r.Intn(max)
	if n != not {
		return n
	}

	return randomNumber(max, not)
}
