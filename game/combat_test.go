package game

import (
	"swarm/entity/npc"
	"swarm/entity/player"
	"testing"
)

//TODO improve code to get it testable
func Test_fight(t *testing.T) {
	type args struct {
		h     *player.Hero
		m     *npc.Monster
		skill string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"",
			args{h: player.NewHero(0, 0), m: npc.NewMonster(1), skill: "1"},
			"You hit monster for 5 damage, monster has 5 HP left \r\nMonster hit you for 5 damage. 145 HP left \r\n",
			false,
		},
		{"",
			args{h: player.NewHero(0, 0), m: npc.NewMonster(2), skill: "1"},
			"You hit monster for 5 damage, monster has 5 HP left \r\nMonster hit you for 10 damage. 140 HP left \r\n",
			false,
		},
		{"",
			args{h: player.NewHero(0, 0), m: npc.NewMonster(3), skill: "1"},
			"You hit monster for 5 damage, monster has 5 HP left \r\nMonster hit you for 15 damage. 135 HP left \r\n",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fight(tt.args.h, tt.args.m, tt.args.skill)
			if (err != nil) != tt.wantErr {
				t.Errorf("fight() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("fight() got = %v, want %v", got, tt.want)
			}
		})
	}
}
