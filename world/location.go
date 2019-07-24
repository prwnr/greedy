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

	monsters := int(size / 2)
	l.PlaceMonsters(monsters)

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

// PlaceMonsters on location
func (l *Location) PlaceMonsters(num int) {
	for {
		if num <= 0 {
			break
		}

		placeMonster(l)
		num--
	}
}

func placeMonster(l *Location) error {
	x := common.RandomNumber(l.Size)
	y := common.RandomNumber(l.Size)

	place := &l.Places[x][y]
	if place.IsOccupied() || place.GetHero() != nil {
		return placeMonster(l)
	}

	place.AddMonster(npc.NewMonster())

	return nil
}

func (l *Location) build() {
	for i := 0; i < l.Size; i++ {
		tmp := make([]Place, 0)
		for j := 0; j < l.Size; j++ {
			p := Place{}
			tmp = append(tmp, p)
		}
		l.Places = append(l.Places, tmp)
	}
}
