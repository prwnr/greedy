package entity

import (
	"regexp"
	"testing"
	"time"
)

func TestSkillStartCoolDown(t *testing.T) {
	s := &skill{
		name:       "Foo",
		internalCD: 0,
		coolDown:   0.5,
	}

	go assertRechargeChannelReceived(t)
	s.startCoolDown()

	assertNumberEquals(t, 0.5, s.currentCoolDown())

	time.Sleep(time.Millisecond * 600)
	assertNumberEquals(t, 0, s.currentCoolDown())
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
			s := &skill{
				name:       tt.fields.Name,
				internalCD: tt.fields.CoolDown,
			}
			go assertRechargeChannelReceived(t)
			if got := s.currentCoolDown(); got != tt.want {
				t.Errorf("currentCoolDown() = %v, want %v", got, tt.want)
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
			s := &skill{
				name:       tt.fields.Name,
				internalCD: tt.fields.CoolDown,
			}
			if got := s.getName(); got != tt.want {
				t.Errorf("getName() = %v, want %v", got, tt.want)
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
		{"hero heals himself", args{heroHealth: 10, heroMana: 10}, "hero health restored by 5.", 15},
		{"hero cannot over heal", args{heroHealth: 150, heroMana: 10}, "hero health restored by 0.", 150},
		{"hero cannot heal when mana is too low", args{heroHealth: 10, heroMana: 0}, "Mana is too low.", 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHero(0, 0)
			h.Health = tt.args.heroHealth
			h.mana = tt.args.heroMana

			s := newHealingSkill(h)

			go assertRechargeChannelReceived(t)
			if got := s.cast(nil); got.Message != tt.wantMessage {
				t.Errorf("cast() = %v, want %v", got, tt.wantMessage)
			}
			_ = s.cast(nil)
			assertHealthEquals(t, tt.wantHealth, h.Health)
		})
	}
}

func TestHeroAttackSkill(t *testing.T) {
	tests := []struct {
		name         string
		skill        string
		wantMessage  string
		wantMinPower int
		wantMaxPower int
		target       killable
	}{
		{"hero basic attack",
			"1",
			"You hit monster for \\d* damage using basic attack, monster has \\d* HP left \r\n",
			12,
			15,
			&Entity{
				Health: 100,
				Attack: 1,
			},
		},
		{"hero heavy attack",
			"3",
			"You hit monster for \\d* damage using heavy attack, monster has \\d* HP left \r\n",
			21,
			36,
			&Entity{
				Health: 100,
				Attack: 1,
			},
		},
		{"hero cant attack nil target",
			"1",
			"",
			12,
			15,
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHero(0, 0)

			go assertRechargeChannelReceived(t)
			got := h.UseSkill(tt.skill, tt.target)
			if got.Power < tt.wantMinPower || got.Power > tt.wantMaxPower {
				t.Errorf("cast() power = %v, want power between %v, %v", got.Power, tt.wantMinPower, tt.wantMaxPower)
			}

			res, err := regexp.MatchString(tt.wantMessage, got.Message)
			if res != true || err != nil {
				t.Errorf("cast() = %v, want %v", got.Message, tt.wantMessage)
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
