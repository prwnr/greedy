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
	//Heal key
	Heal = "2"
)

// NewGame starts new game with all requirements.
func NewGame() Game {
	game := Game{}

	loadConfig(&game)

	game.View = view.NewView()
	game.Hero = player.NewHero(0, game.Config.LocationSize-1)
	game.CurrentLocation = world.NewLocation(game.Config.LocationSize)

	game.CurrentLocation.PlaceHero(game.Hero)
	game.View.UpdateLocation(game.CurrentLocation.RenderPlaces())

	game.View.UpdateHeroStats(game.Hero.GetStats())

	return game
}

// PlayerAction changes hero position
func (g *Game) PlayerAction(action string) {
	currPlace := g.CurrentLocation.GetHeroPlace(g.Hero)

	maxStep := g.CurrentLocation.Size - 1
	switch action {
	case MoveUp:
		g.Hero.MoveUp()
	case MoveDown:
		g.Hero.MoveDown(maxStep)
	case MoveLeft:
		g.Hero.MoveLeft()
	case MoveRight:
		g.Hero.MoveRight(maxStep)
	case Attack:
		if currPlace.IsOccupied() {
			c := combat.NewCombat(currPlace.GetHero(), currPlace.GetMonster())
			res, err := c.Fight()
			if err != nil {
				_ = fmt.Errorf("fight error: %v", err)
			} else {
				if !currPlace.GetMonster().IsAlive() {
					currPlace.RemoveMonster()
				}
				g.View.UpdateCombatLog(res)
			}
		}

		return
	case Heal:
		res := g.Hero.UseHeal()
		g.View.UpdateCombatLog(res)

		return
	}

	if currPlace.IsOccupied() {
		c := combat.NewCombat(currPlace.GetHero(), currPlace.GetMonster())
		res := c.AttackBack()
		g.View.UpdateCombatLog(res)
	}

	currPlace.RemoveHero()
	g.CurrentLocation.PlaceHero(g.Hero)
}

// UpdateView updates main views of the game
func (g *Game) UpdateView() {
	g.View.UpdateLocation(g.CurrentLocation.RenderPlaces())
	g.View.UpdateHeroStats(g.Hero.GetStats())

	currPlace := g.CurrentLocation.GetHeroPlace(g.Hero)
	if currPlace.IsOccupied() {
		g.View.ShowMonster(currPlace.GetMonster().GetStats())
	} else {
		g.View.HideMonster()
	}
}
