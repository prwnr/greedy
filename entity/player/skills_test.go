package player

import (
	"regexp"
	"swarm/entity"
	"testing"
	"time"
)

func TestSkillStartCoolDown(t *testing.T) {
	s := &Skill{
		Name:     "Foo",
		CoolDown: 0,
	}

	go assertRechargeChannelReceived(t)
	s.startCoolDown(0.5)

	assertNumberEquals(t, 0.5, s.CoolDown)

	time.Sleep(time.Millisecond * 600)
	assertNumberEquals(t, 0, s.CoolDown)
}

func TestSkillCurrentCoolDown(t *testing.T) {
	type fields struct {
		Name     string
		CoolDown float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{"returns current cool down", fields{Name: "Foo", CoolDown: 5}, 5},
		{"returns current cool down", fields{Name: "Foo", CoolDown: 2}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Skill{
				Name:     tt.fields.Name,
				CoolDown: tt.fields.CoolDown,
			}
			go assertRechargeChannelReceived(t)
			if got := s.CurrentCoolDown(); got != tt.want {
				t.Errorf("CurrentCoolDown() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSkillGetName(t *testing.T) {
	type fields struct {
		Name     string
		CoolDown float64
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"returns name", fields{Name: "Foo", CoolDown: 5}, "Foo"},
		{"returns name", fields{Name: "Bar", CoolDown: 2}, "Bar"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Skill{
				Name:     tt.fields.Name,
				CoolDown: tt.fields.CoolDown,
			}
			if got := s.GetName(); got != tt.want {
				t.Errorf("GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeroHealingSkill(t *testing.T) {
	assertHealthEquals := func(t *testing.T, want, got int) {
		if want != got {
			t.Errorf("expected hero health to be at %d, but got %d", want, got)
		}
	}

	type args struct {
		heroHealth int
		heroMana   int
	}

	tests := []struct {
		name        string
		args        args
		wantMessage string
		wantHealth  int
	}{
		{"hero heals himself", args{heroHealth: 10, heroMana: 10}, "Hero health restored by 5.", 15},
		{"hero cannot over heal", args{heroHealth: 150, heroMana: 10}, "Hero health restored by 0.", 150},
		{"hero cannot heal when mana is too low", args{heroHealth: 10, heroMana: 0}, "Mana is too low.", 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHero(0, 0)
			h.Health = tt.args.heroHealth
			h.mana = tt.args.heroMana

			s := NewHealingSkill(h)

			go assertRechargeChannelReceived(t)
			if got := s.Cast(nil); got.Message != tt.wantMessage {
				t.Errorf("Cast() = %v, want %v", got, tt.wantMessage)
			}
			assertHealthEquals(t, tt.wantHealth, h.Health)
		})
	}
}

func TestHeroAttackSkill(t *testing.T) {
	type args struct {
		heroHealth int
		heroMana   int
	}

	tests := []struct {
		name         string
		wantMessage  string
		wantMinPower int
		wantMaxPower int
		target       Killable
	}{
		{"hero attacks monster",
			"You hit monster for \\d* damage, monster has \\d* HP left \r\n",
			10,
			15,
			&entity.Entity{
				Health: 100,
				Attack: 1,
			},
		},
		{"hero cant attack nil target",
			"",
			10,
			15,
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHero(0, 0)
			s := NewBasicAttackSkill(h)

			go assertRechargeChannelReceived(t)
			got := s.Cast(tt.target)
			if got.Power < tt.wantMinPower || got.Power > tt.wantMaxPower {
				t.Errorf("Cast() power = %v, want power between %v, %v", got.Power, tt.wantMinPower, tt.wantMaxPower)
			}

			res, err := regexp.MatchString(tt.wantMessage, got.Message)
			if res != true || err != nil {
				t.Errorf("Cast() = %v, want %v", got.Message, tt.wantMessage)
			}
		})
	}
}

func assertNumberEquals(t *testing.T, want, got float64) {
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func assertRechargeChannelReceived(t *testing.T) {
	for {
		res := <-RechargeSkill
		if res != true {
			t.Errorf("RechargeSkill channel is not true")
		}
	}
}
