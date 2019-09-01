package game

import (
	"fmt"
	"swarm/common"
	"swarm/entity"
	"swarm/modifiers"
	"swarm/view"
	"swarm/world"
	"sync"
	"time"
)

// Game struct
type Game struct {
	//View represents game UI
	View *view.View
	//Hero is a player
	Hero *entity.Hero
	//CurrentLocation is a location/map on which Hero is walking
	CurrentLocation *world.Location
	KillsCount      int
	TimeElapsed     int
	//Config for the game
	Config Config
	//Over defines if hero is defeated or time ran our
	Over          bool
	swarmReleased bool
	M             sync.Mutex
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
	g.Hero = entity.NewHero(0, g.Config.LocationSize-1)
	g.CurrentLocation = world.NewLocation(g.Config.LocationSize, 1)

	g.CurrentLocation.PlaceHero(g.Hero)

	return g
}

// RunLocked executes passed function in a locked mutex
func (g *Game) RunLocked(f func()) {
	g.M.Lock()
	defer g.M.Unlock()
	f()
}

// InitViews sets up first UI data.
func (g *Game) InitViews() {
	g.View.UpdateLocationTitle(g.CurrentLocation.Level())
	g.View.UpdateLocation(g.CurrentLocation.RenderPlaces())

	locReq := g.CurrentLocation.Requirements
	g.View.UpdateGoal(locReq.MonsterTarget, locReq.KillsCount, locReq.TimeFrame)

	g.View.UpdateHeroStats(g.Hero.GetStats())
	g.View.UpdateSkillBar(g.Hero.Skills())
}

// Cycle in game
func (g *Game) Cycle(second int64) {
	if g.CurrentLocation.Requirements.KillsCount == g.KillsCount && !g.timeIsOver() {
		g.nextLocation()
	}

	if !g.timeIsOver() {
		g.TimeElapsed++
		g.updateGoal()
	}

	if g.timeIsOver() && !g.swarmReleased {
		go g.releaseSwarm()
		g.swarmReleased = true
	}

	if second%g.Config.MonsterSpawn == 0 {
		g.CurrentLocation.PlaceMonsters(g.Config.MonstersSpawnNum)
	}

	if second%modifiers.HeroRegenTimeout == 0 {
		g.Hero.Regenerate()
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

		g.checkHeroStatus()
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

	g.checkHeroStatus()

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

func (g *Game) countKill(m *entity.Monster) {
	if m.Render() == g.CurrentLocation.Requirements.MonsterTarget {
		g.KillsCount++

		g.updateGoal()
	}
}

// updateGoal triggers view update on location goal
func (g *Game) updateGoal() {
	t := g.CurrentLocation.Requirements.TimeFrame - g.TimeElapsed
	k := g.CurrentLocation.Requirements.KillsCount - g.KillsCount
	g.View.UpdateGoal(g.CurrentLocation.Requirements.MonsterTarget, k, t)
}

// timeIsOver checks if location time passed
func (g *Game) timeIsOver() bool {
	return g.CurrentLocation.Requirements.TimeFrame-g.TimeElapsed <= 0
}

// checkHeroStatus if its alive. If not, game is over
func (g *Game) checkHeroStatus() {
	if !g.Hero.IsAlive() {
		g.View.UpdateCombatLog("Hero died. Press 'q' to quit or 'r' to restart.")
		g.Over = true
	}
}

// releaseSwarm monsters
func (g *Game) releaseSwarm() {
	t := time.NewTicker(time.Millisecond * 20)
	for {
		select {
		case <-t.C:
			g.CurrentLocation.PlaceMonsters(1)
			g.View.UpdateLocation(g.CurrentLocation.RenderPlaces())
			if !g.CurrentLocation.HasFreePlace() {
				t.Stop()
				return
			}
		}
	}
}

// nextLocation advance
func (g *Game) nextLocation() {
	g.CurrentLocation = world.NewLocation(g.Config.LocationSize, g.CurrentLocation.Level()+1)
	g.View.UpdateLocationTitle(g.CurrentLocation.Level())
	g.KillsCount = 0
	g.TimeElapsed = 0
	g.CurrentLocation.PlaceHero(g.Hero)
}
