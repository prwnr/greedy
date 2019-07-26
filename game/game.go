package game

import (
	"fmt"
	"swarm/combat"
	"swarm/player"
	"swarm/view"
	"swarm/world"
)

// Game struct
type Game struct {
	//View represents game UI
	View *view.View
	//Hero is a player
	Hero *player.Hero
	//CurrentLocation is a location/map on which Hero is walking
	CurrentLocation *world.Location
	//Config for the game
	Config Config
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
	game := Game{}

	loadConfig(&game)

	game.View = view.NewView()
	game.Hero = player.NewHero(0, game.Config.LocationSize-1)
	game.CurrentLocation = world.NewLocation(game.Config.LocationSize)

	currPlace := &game.CurrentLocation.Places[game.Hero.Position.Y][game.Hero.Position.X]
	currPlace.SetHero(game.Hero)
	game.View.UpdateLocation(game.CurrentLocation.RenderPlaces())

	game.View.UpdateHeroStats(game.Hero.GetStats())

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
				if !currPlace.GetMonster().IsAlive() {
					currPlace.RemoveMonster()
				}
				g.View.UpdateCombatLog(res)
			}
		}

		return
	}

	currPlace.RemoveHero()
	newPlace := &g.CurrentLocation.Places[g.Hero.Position.Y][g.Hero.Position.X]
	newPlace.SetHero(g.Hero)
}

// UpdateView updates main views of the game
func (g *Game) UpdateView() {
	g.View.UpdateLocation(g.CurrentLocation.RenderPlaces())
	g.View.UpdateHeroStats(g.Hero.GetStats())
}
