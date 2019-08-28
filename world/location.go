package world

import (
	"strings"
	"swarm/common"
	"swarm/entity/npc"
	"swarm/entity/player"
	"swarm/modifiers"
)

//Location of the game
type Location struct {
	Size         int
	Places       [][]Place
	level        int
	Requirements levelRequirements
}

// Level of the location
func (l *Location) Level() int {
	return l.level
}

//NewLocation creates new game Map with places
func NewLocation(size, level int) *Location {
	req := levelRequirements{
		TimeFrame:     modifiers.LocationBaseTimeFrame * int(float64(level)*1.5),
		MonsterTarget: modifiers.LocationMonsterTarget,
		KillsCount:    modifiers.LocationBaseKillsCount * level,
	}

	l := &Location{
		Size:         size,
		level:        level,
		Requirements: req,
	}

	l.build()

	monsters := size / 2
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

// GetHeroPlace returns place on which hero currently is
func (l *Location) GetHeroPlace(h *player.Hero) *Place {
	return &l.Places[h.Position.Y][h.Position.X]
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

	level := common.RandomMinNumber(l.level+0, l.level+2)
	m := npc.NewMonster(level)
	if l.level > 1 {
		m.SetLook(npc.LevelLook[level-l.level+1])
	}

	place.AddMonster(m)

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

type levelRequirements struct {
	TimeFrame     int
	MonsterTarget string
	KillsCount    int
}
