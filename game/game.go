package game

import (
	"fmt"
	"swarm/common"
	"swarm/entity/npc"
	"swarm/entity/player"
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
	KillsCount      int
	TimeElapsed     int
	//Config for the game
	Config Config
	//Over defines if hero is defeated or time ran our
	Over          bool
	swarmReleased bool
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

	locReq := g.CurrentLocation.Requirements
	g.View.UpdateGoal(locReq.MonsterTarget, locReq.KillsCount, locReq.TimeFrame)

	g.View.UpdateHeroStats(g.Hero.GetStats())
	g.View.UpdateSkillBar(g.Hero.Skills())

	return g
}

// Cycle in game
func (g *Game) Cycle(second int64) {
	if g.CurrentLocation.Requirements.KillsCount == g.KillsCount && !g.TimeIsOver() {
		g.NextLocation()
	}

	if !g.TimeIsOver() {
		g.TimeElapsed++
		g.UpdateGoal()
	}

	if g.TimeIsOver() && !g.swarmReleased {
		go g.ReleaseSwarm()
		g.swarmReleased = true
	}

	if second%g.Config.MonsterSpawn == 0 {
		go g.CurrentLocation.PlaceMonsters(g.Config.MonstersSpawnNum)
	}

	if second%player.RegenTimeout == 0 {
		go g.Hero.Regenerate()
	}
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
				return
			}

			if !monster.IsAlive() {
				g.countKill(monster)

				currPlace.RemoveMonster()
				if hero.IsAlive() {
					res += hero.GainExperience(monster.GetExperienceValue())
				}
			}

			g.View.UpdateCombatLog(res)
		} else {
			if action == Heal {
				res := g.Hero.UseSkill(Heal, nil)
				g.View.UpdateCombatLog(res.Message)
			}
		}

		g.CheckHeroStatus()

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

	g.CheckHeroStatus()

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

func (g *Game) countKill(m *npc.Monster) {
	if m.Render() == g.CurrentLocation.Requirements.MonsterTarget {
		g.KillsCount++

		g.UpdateGoal()
	}
}

// UpdateGoal triggers view update on location goal
func (g *Game) UpdateGoal() {
	t := g.CurrentLocation.Requirements.TimeFrame - g.TimeElapsed
	k := g.CurrentLocation.Requirements.KillsCount - g.KillsCount
	g.View.UpdateGoal(g.CurrentLocation.Requirements.MonsterTarget, k, t)
}

// TimeIsOver checks if location time passed
func (g *Game) TimeIsOver() bool {
	return g.CurrentLocation.Requirements.TimeFrame-g.TimeElapsed <= 0
}

// CheckHeroStatus if its alive. If not, game is over
func (g *Game) CheckHeroStatus() {
	if !g.Hero.IsAlive() {
		g.View.UpdateCombatLog("Hero died. Press 'q' to quit or 'r' to restart.")
		g.Over = true
	}
}

// ReleaseSwarm monsters
func (g *Game) ReleaseSwarm() {
	t := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-t.C:
			g.CurrentLocation.PlaceMonsters(10)
			if !g.CurrentLocation.HasFreePlace() {
				t.Stop()
				return
			}
		}
	}
}

// NextLocation advance
func (g *Game) NextLocation() {
	g.CurrentLocation = world.NewLocation(g.Config.LocationSize, g.CurrentLocation.Level()+1)
	g.View.UpdateLocationTitle(g.CurrentLocation.Level())
	g.KillsCount = 0
	g.TimeElapsed = 0
	g.CurrentLocation.PlaceHero(g.Hero)
}
