package game

import (
	"fmt"
	"swarm/common"
	"swarm/entity/player"
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
	//Heal key
	Heal = "2"
)

// NewGame starts new game with all requirements.
func NewGame() Game {
	g := Game{}

	loadConfig(&g)

	g.View = view.NewView()
	g.Hero = player.NewHero(0, g.Config.LocationSize-1)
	g.CurrentLocation = world.NewLocation(g.Config.LocationSize, 1)

	g.CurrentLocation.PlaceHero(g.Hero)
	g.View.UpdateLocationTitle(g.CurrentLocation.Level())
	g.View.UpdateLocation(g.CurrentLocation.RenderPlaces())

	g.View.UpdateHeroStats(g.Hero.GetStats())
	g.View.UpdateSkillBar(g.Hero.Skills())

	return g
}

// PlayerAction changes hero position
func (g *Game) PlayerAction(action string) {
	currPlace := g.CurrentLocation.GetHeroPlace(g.Hero)

	maxStep := g.CurrentLocation.Size - 1
	hero := currPlace.GetHero()
	monster := currPlace.GetMonster()

	if isSkillAction(action) {
		if currPlace.IsOccupied() {
			res, err := fight(hero, monster, action)
			if err != nil {
				_ = fmt.Errorf("fight error: %v", err)
			} else {
				if !monster.IsAlive() {
					currPlace.RemoveMonster()
					if hero.IsAlive() {
						res += hero.GainExperience(monster.GetExperienceValue())
					}
				}
				g.View.UpdateCombatLog(res)
			}
		} else {
			if action == Heal {
				res := g.Hero.UseSkill(Heal, nil)
				g.View.UpdateCombatLog(res.Message)
			}
		}

		return
	}

	if isMovement(action) {
		switch action {
		case MoveUp:
			g.Hero.MoveUp()
		case MoveDown:
			g.Hero.MoveDown(maxStep)
		case MoveLeft:
			g.Hero.MoveLeft()
		case MoveRight:
			g.Hero.MoveRight(maxStep)
		}

		if currPlace.IsOccupied() {
			res := fightBack(hero, monster)
			g.View.UpdateCombatLog(res)
		}
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

func isSkillAction(action string) bool {
	skill := []string{"1", "3", Heal}
	return common.SliceContains(skill, action)
}

func isMovement(action string) bool {
	move := []string{MoveUp, MoveRight, MoveLeft, MoveDown}
	return common.SliceContains(move, action)
}
