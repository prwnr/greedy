package game

import "testing"

func TestLoadConfig(t *testing.T) {
	type args struct {
		MonsterSpawn     int64
		MonstersSpawnNum int
		LocationSize     int
	}

	assertConfigArgsMatch := func(t *testing.T, got Config, want args) {
		if got.MonsterSpawn != want.MonsterSpawn {
			t.Errorf("got MonsterSpawn = %d, want %d", got.MonsterSpawn, want.MonsterSpawn)
		}

		if got.MonstersSpawnNum != want.MonstersSpawnNum {
			t.Errorf("got MonstersSpawnNum = %d, want %d", got.MonstersSpawnNum, want.MonstersSpawnNum)
		}

		if got.LocationSize != want.LocationSize {
			t.Errorf("got LocationSize = %d, want %d", got.LocationSize, want.LocationSize)
		}
	}

	tests := []struct {
		name string
		game *Game
		want args
	}{
		{"loads default config", &Game{}, args{
			MonsterSpawn:     5,
			MonstersSpawnNum: 2,
			LocationSize:     18,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loadConfig(tt.game)
			assertConfigArgsMatch(t, tt.game.Config, tt.want)
		})
	}
}
