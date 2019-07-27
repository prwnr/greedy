package world

import (
	"strings"
	"swarm/common"
	"swarm/npc"
	"swarm/player"
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
			return
		}

		_ = placeMonster(l)
		num--
	}
}

// PlaceHero on location (based on his current position)
func (l *Location) PlaceHero(h *player.Hero) {
	l.Places[h.Position.Y][h.Position.X].SetHero(h)
}

// HasFreePlace checks if every place on location is occupied
func (l *Location) HasFreePlace() bool {
	for _, row := range l.Places {
		for _, p := range row {
			if !p.IsOccupied() && p.GetHero() == nil {
				return true
			}
		}
	}

	return false
}

func placeMonster(l *Location) bool {
	if !l.HasFreePlace() {
		return false
	}

	x := common.RandomNumber(l.Size)
	y := common.RandomNumber(l.Size)

	place := &l.Places[x][y]
	if place.IsOccupied() || place.GetHero() != nil {
		return placeMonster(l)
	}

	place.AddMonster(npc.NewMonster())

	return true
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
