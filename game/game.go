package game

import (
	"fmt"
	"swarm/combat"
	"swarm/player"
	"swarm/view"
	"swarm/world"
	"time"
)

// Game struct
type Game struct {
	//View represents game UI
	View *view.View
	//Hero is a player
	Hero *player.Hero
	//CurrentLocation is a location/map on which Hero is walking
	CurrentLocation *world.Location
	//MobsRespawn time in seconds
	MobsRespawn time.Duration
}

const (
	//MoveUp key
	MoveUp = "w"
	//MoveDown key
	MoveDown = "s"
	//MoveLeft key
	MoveLeft = "a"
	//MoveRight key
	MoveRight = "d"
	//Attack key
	Attack = "1"
)

// NewGame starts new game with all requirements.
func NewGame() Game {
	game := Game{
		MobsRespawn: 15 * time.Second,
	}

	game.View = view.NewView()
	game.Hero = player.NewHero(0, 10-1)
	game.CurrentLocation = world.NewLocation(10)

	currPlace := &game.CurrentLocation.Places[game.Hero.Position.Y][game.Hero.Position.X]
	currPlace.SetHero(game.Hero)
	game.View.UpdateLocation(game.CurrentLocation.RenderPlaces())

	return game
}

// MoveHero changes hero position
func (g *Game) MoveHero(direction string) {
	currPlace := &g.CurrentLocation.Places[g.Hero.Position.Y][g.Hero.Position.X]

	switch direction {
	case MoveUp:
		if g.Hero.Position.Y > 0 {
			g.Hero.Position.Y--
		}
	case MoveDown:
		if g.Hero.Position.Y < g.CurrentLocation.Size-1 {
			g.Hero.Position.Y++
		}
	case MoveLeft:
		if g.Hero.Position.X > 0 {
			g.Hero.Position.X--
		}
	case MoveRight:
		if g.Hero.Position.X < g.CurrentLocation.Size-1 {
			g.Hero.Position.X++
		}
	case Attack:
		if currPlace.IsOccupied() {
			c := combat.NewCombat(currPlace.GetHero(), currPlace.GetMonster())
			res, err := c.Fight()
			if err != nil {
				fmt.Errorf("Fight error: %v", err)
			} else {
				g.View.UpdateCombatLog(res)
				g.View.UpdateLocation(g.CurrentLocation.RenderPlaces())
			}
		}
	}

	currPlace.RemoveHero()
	newPlace := &g.CurrentLocation.Places[g.Hero.Position.Y][g.Hero.Position.X]
	newPlace.SetHero(g.Hero)

	g.View.UpdateLocation(g.CurrentLocation.RenderPlaces())
}
