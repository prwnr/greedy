package entity

import (
	"testing"
)

func TestCombatFight(t *testing.T) {
	tests := []struct {
		name     string
		attacker *Entity
		defender *Entity
		want     string
		wantErr  bool
	}{
		{"attacks for 5 power",
			&Entity{Health: 10, Attack: 5},
			&Entity{Health: 10, Attack: 5},
			"You hit monster for 5 damage, monster has 5 HP left \r\nMonster hit you for 5 damage. 5 HP left \r\n",
			false,
		},
		{"attacks for 10 power",
			&Entity{Health: 10, Attack: 10},
			&Entity{Health: 10, Attack: 5},
			"You hit monster for 10 damage, monster has 0 HP left \r\nMonster hit you for 5 damage. 5 HP left \r\n",
			false,
		},
		{"cannot attack dead defender",
			&Entity{Health: 10, Attack: 10},
			&Entity{Health: 0, Attack: 5},
			"",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCombat(tt.attacker, tt.defender)

			got, err := c.Fight()
			if (err != nil) != tt.wantErr {
				t.Errorf("Fight() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Fight() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCombatAttackBack(t *testing.T) {
	tests := []struct {
		name     string
		attacker *Entity
		defender *Entity
		want     string
	}{
		{"defender attacks back for 5 power",
			&Entity{Health: 10, Attack: 5},
			&Entity{Health: 10, Attack: 5},
			"Monster hit you for 5 damage. 5 HP left \r\n",
		},
		{"defender attacks back for 10 power",
			&Entity{Health: 10, Attack: 5},
			&Entity{Health: 10, Attack: 10},
			"Monster hit you for 10 damage. 0 HP left \r\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCombat(tt.attacker, tt.defender)

			if got := c.AttackBack(); got != tt.want {
				t.Errorf("AttackBack() = %v, want %v", got, tt.want)
			}
		})
	}
}
