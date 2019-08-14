package player

import (
	"testing"
	"time"
)

func TestSkillStartCoolDown(t *testing.T) {
	s := &Skill{
		Name:     "Foo",
		CoolDown: 0,
	}

	s.startCoolDown(1)
	assertNumberEquals(t, 1, s.CoolDown)

	time.Sleep(time.Millisecond * 1200)
	assertNumberEquals(t, 0, s.CoolDown)
}

func TestSkillCurrentCoolDown(t *testing.T) {
	type fields struct {
		Name     string
		CoolDown int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
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
			if got := s.CurrentCoolDown(); got != tt.want {
				t.Errorf("CurrentCoolDown() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSkillGetName(t *testing.T) {
	type fields struct {
		Name     string
		CoolDown int
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

func assertNumberEquals(t *testing.T, want, got int) {
	if want != got {
		t.Errorf("want %d, got %d", want, got)
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

			if got := s.Cast(nil); got.Message != tt.wantMessage {
				t.Errorf("Cast() = %v, want %v", got, tt.wantMessage)
			}
			assertHealthEquals(t, tt.wantHealth, h.Health)
		})
	}
}
