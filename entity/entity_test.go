package entity

import "testing"

func TestEntityAttackPower(t *testing.T) {
	type fields struct {
		Health int
		Attack int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"has 10 attack power", fields{Health: 10, Attack: 10}, 10},
		{"has 15 attack power", fields{Health: 10, Attack: 15}, 15},
		{"has 25 attack power", fields{Health: 10, Attack: 25}, 25},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Entity{
				Health: tt.fields.Health,
				Attack: tt.fields.Attack,
			}
			if got := e.AttackPower(); got != tt.want {
				t.Errorf("AttackPower() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntityGetHealth(t *testing.T) {
	type fields struct {
		Health int
		Attack int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"has 10 health", fields{Health: 10, Attack: 1}, 10},
		{"has 15 health", fields{Health: 15, Attack: 1}, 15},
		{"has 25 health", fields{Health: 25, Attack: 1}, 25},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Entity{
				Health: tt.fields.Health,
				Attack: tt.fields.Attack,
			}
			if got := e.GetHealth(); got != tt.want {
				t.Errorf("GetHealth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntityIsAlive(t *testing.T) {
	type fields struct {
		Health int
		Attack int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"is alive with health above zero", fields{Health: 10, Attack: 10}, true},
		{"is dead with health at zero", fields{Health: 0, Attack: 10}, false},
		{"is dead with health below zero", fields{Health: -100, Attack: 10}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Entity{
				Health: tt.fields.Health,
				Attack: tt.fields.Attack,
			}
			if got := e.IsAlive(); got != tt.want {
				t.Errorf("IsAlive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntityReduceHealth(t *testing.T) {
	type fields struct {
		Health int
		Attack int
	}
	tests := []struct {
		name   string
		fields fields
		amount int
		want   int
	}{
		{"reduces health to zero", fields{Health: 10, Attack: 10}, 10, 0},
		{"reduces health to 50%", fields{Health: 20, Attack: 10}, 10, 10},
		{"reduces health below zero", fields{Health: 20, Attack: 10}, 30, -10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Entity{
				Health: tt.fields.Health,
				Attack: tt.fields.Attack,
			}

			e.ReduceHealth(tt.amount)

			if got := e.GetHealth(); got != tt.want {
				t.Errorf("After ReduceHealth(), got GetHealth() = %v, want %v", got, tt.want)
			}
		})
	}
}
