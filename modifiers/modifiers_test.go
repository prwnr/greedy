package modifiers

import "testing"

func TestCalculateMonsterHealth(t *testing.T) {
	type args struct {
		l int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"monster level 1 health", args{l: 1}, 30},
		{"monster level 2 health", args{l: 2}, 65},
		{"monster level 3 health", args{l: 3}, 100},
		{"monster level 10 health", args{l: 10}, 345},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateMonsterHealth(tt.args.l); got != tt.want {
				t.Errorf("CalculateMonsterHealth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateMonsterAttack(t *testing.T) {
	type args struct {
		l int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"monster level 1 attack", args{l: 1}, 4},
		{"monster level 2 attack", args{l: 2}, 8},
		{"monster level 3 attack", args{l: 3}, 12},
		{"monster level 10 attack", args{l: 10}, 40},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateMonsterAttack(tt.args.l); got != tt.want {
				t.Errorf("CalculateMonsterAttack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateMonsterExperience(t *testing.T) {
	type args struct {
		l int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"monster level 1 experience", args{l: 1}, 15},
		{"monster level 2 experience", args{l: 2}, 30},
		{"monster level 3 experience", args{l: 3}, 45},
		{"monster level 10 experience", args{l: 10}, 150},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateMonsterExperience(tt.args.l); got != tt.want {
				t.Errorf("CalculateMonsterExperience() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateHeroAttack(t *testing.T) {
	type args struct {
		l int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"hero level 2 attack", args{l: 2}, 20},
		{"hero level 3 attack", args{l: 3}, 22},
		{"hero level 3 attack", args{l: 4}, 25},
		{"hero level 10 attack", args{l: 10}, 40},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateHeroAttack(tt.args.l); got != tt.want {
				t.Errorf("CalculateHeroAttack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateHeroHealth(t *testing.T) {
	type args struct {
		l int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"hero level 2 health", args{l: 2}, 200},
		{"hero level 3 health", args{l: 3}, 225},
		{"hero level 4 health", args{l: 4}, 250},
		{"hero level 10 health", args{l: 10}, 400},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateHeroHealth(tt.args.l); got != tt.want {
				t.Errorf("CalculateHeroHealth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateHeroMana(t *testing.T) {
	type args struct {
		l int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"hero level 2 mana", args{l: 2}, 84},
		{"hero level 3 mana", args{l: 3}, 96},
		{"hero level 4 mana", args{l: 4}, 108},
		{"hero level 10 mana", args{l: 10}, 180},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateHeroMana(tt.args.l); got != tt.want {
				t.Errorf("CalculateHeroMana() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateHeroLevelExperience(t *testing.T) {
	type args struct {
		l int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"level 2 required experience", args{l: 2}, 120},
		{"level 3 required experience", args{l: 3}, 320},
		{"level 4 required experience", args{l: 4}, 600},
		{"level 5 required experience", args{l: 5}, 960},
		{"level 6 required experience", args{l: 6}, 1400},
		{"level 10 required experience", args{l: 10}, 3960},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateHeroLevelExperience(tt.args.l); got != tt.want {
				t.Errorf("CalculateHeroLevelExperience() = %v, want %v", got, tt.want)
			}
		})
	}
}
