package game

import (
	"swarm/entity/npc"
	"swarm/view"
	"swarm/world"
	"testing"
)

func assertViewChannelReceived(t *testing.T) {
	for {
		res := <-view.UIChange
		if res != true {
			t.Errorf("UIChange channel is not true")
		}
	}
}

func TestMovesHeroOnLocation(t *testing.T) {
	g := NewGame()
	g.Hero.Position.X = 0
	g.Hero.Position.Y = 0

	go assertViewChannelReceived(t)
	g.PlayerAction(MoveRight)
	if g.Hero.Position.X != 1 {
		t.Errorf("hero X position should be 1, was %d", g.Hero.Position.X)
	}

	g.PlayerAction(MoveDown)
	if g.Hero.Position.Y != 1 {
		t.Errorf("hero Y position should be 1, was %d", g.Hero.Position.Y)
	}
}

func TestMovingHeroToNewLocationRemovesHimFromOld(t *testing.T) {
	g := NewGame()
	g.Hero.Position.X = 0
	g.Hero.Position.Y = 0

	go assertViewChannelReceived(t)
	g.PlayerAction(MoveRight)
	got := g.CurrentLocation.Places[0][0].GetHero()
	if got != nil {
		t.Error("there should be no here on this place, got one")
	}
}

func TestWontMoveHeroOutsideLocation(t *testing.T) {
	g := NewGame()
	g.Hero.Position.X = 0
	g.Hero.Position.Y = 0

	go assertViewChannelReceived(t)
	g.PlayerAction(MoveUp)
	if g.Hero.Position.Y != 0 {
		t.Errorf("hero Y position should be 0, was %d", g.Hero.Position.Y)
	}

	g.PlayerAction(MoveLeft)
	if g.Hero.Position.X != 0 {
		t.Errorf("hero X position should be 0, was %d", g.Hero.Position.X)
	}
}

func TestGame_ReleasingSwarm(t *testing.T) {
	go assertViewChannelReceived(t)
	g := NewGame()
	g.CurrentLocation = world.NewLocation(2, 1)

	g.ReleaseSwarm()
	if g.CurrentLocation.HasFreePlace() {
		t.Errorf("current location should be full of monsters")
	}
}

func Test_isSkillAction(t *testing.T) {
	tests := []struct {
		name   string
		action string
		want   bool
	}{
		{"is a skill", "1", true},
		{"is a skill", "2", true},
		{"is not a skill", "5", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSkillAction(tt.action); got != tt.want {
				t.Errorf("isSkillAction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isMovement(t *testing.T) {
	tests := []struct {
		name   string
		action string
		want   bool
	}{
		{"is movement", "w", true},
		{"is movement", "a", true},
		{"is movement", "s", true},
		{"is movement", "d", true},
		{"is not a movement", "1", false},
		{"is not a movement", "t", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMovement(tt.action); got != tt.want {
				t.Errorf("isMovement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGame_countKill(t *testing.T) {
	go assertViewChannelReceived(t)
	g := NewGame()
	m := npc.NewMonster(3)
	g.countKill(m)

	if g.KillsCount != 1 {
		t.Errorf("kills count should be 1, got %d", g.KillsCount)
	}
}

func TestGame_TimeIsOver(t *testing.T) {
	g := NewGame()
	g.TimeElapsed = 200

	if got := g.TimeIsOver(); got != true {
		t.Errorf("TimeIsOver() = %v, want true", got)
	}
}

func TestGame_CheckHeroStatus(t *testing.T) {
	go assertViewChannelReceived(t)
	g := NewGame()

	g.CheckHeroStatus()
	if g.Over {
		t.Error("game cant be over when here is still alive")
	}

	g.Hero.ReduceHealth(1000)
	g.CheckHeroStatus()
	if !g.Over {
		t.Error("game should be over when here is dead")
	}
}

func TestGame_NextLocation(t *testing.T) {
	go assertViewChannelReceived(t)
	g := NewGame()

	assertLocationLevel := func(t *testing.T, got, want int) {
		if got != want {
			t.Errorf("current location should be level %d, got %d", want, got)
		}
	}

	assertLocationLevel(t, g.CurrentLocation.Level(), 1)

	g.NextLocation()
	assertLocationLevel(t, g.CurrentLocation.Level(), 2)
}
