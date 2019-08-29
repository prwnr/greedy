package game

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Config for current game
type Config struct {
	MonsterSpawn     int64 `json:"monster_spawn,omitempty"`
	MonstersSpawnNum int   `json:"monsters_spawn_num,omitempty"`
	LocationSize     int   `json:"location_size,omitempty"`
}

func loadConfig(g *Game) {
	jsonFile, err := os.Open("config.json")
	config := Config{
		MonsterSpawn:     5,
		MonstersSpawnNum: 2,
		LocationSize:     18,
	}

	if err == nil {
		defer jsonFile.Close()
		bytes, _ := ioutil.ReadAll(jsonFile)
		_ = json.Unmarshal(bytes, &config)
	}

	g.Config = config
}
