package game

import (
	"regexp"
	"swarm/entity"
	"testing"
)

//TODO improve code to get it testable
func Test_fight(t *testing.T) {
	type args struct {
		h     *entity.Hero
		m     *entity.Monster
		skill string
	}

	deadMonster := func() *entity.Monster {
		m := entity.NewMonster(1)
		m.ReduceHealth(100)
		return m
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"attacks 1st level monster",
			args{h: entity.NewHero(0, 0), m: entity.NewMonster(1), skill: "1"},
			"You hit monster for \\d* damage using basic attack, monster has \\d* HP left \r\nMonster hit you for \\d* damage. \\d* HP left \r\n",
			false,
		},
		{"attacks 3rd level monster",
			args{h: entity.NewHero(0, 0), m: entity.NewMonster(2), skill: "1"},
			"You hit monster for \\d* damage using basic attack, monster has \\d* HP left \r\nMonster hit you for \\d* damage. \\d* HP left \r\n",
			false,
		},
		{"cannot attack dead monster",
			args{h: entity.NewHero(0, 0), m: deadMonster(), skill: "1"},
			"",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go assertRechargeChannelReceived(t)
			got, err := fight(tt.args.h, tt.args.m, tt.args.skill)
			if (err != nil) != tt.wantErr {
				t.Errorf("fight() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			res, err := regexp.MatchString(tt.want, got)
			if res != true || err != nil {
				t.Errorf("fight() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func assertRechargeChannelReceived(t *testing.T) {
	for {
		res := <-entity.RechargeSkill
		if res != true {
			t.Errorf("RechargeSkill channel is not true")
		}
	}
}
