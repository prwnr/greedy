package game

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"
)

// Config for current game
type Config struct {
	MonsterSpawn time.Duration `json:"monster_spawn,omitempty"`
	LocationSize int           `json:"location_size,omitempty"`
}

func loadConfig(g *Game) {
	jsonFile, err := os.Open("config.json")
	config := Config{
		MonsterSpawn: 10,
		LocationSize: 18,
	}

	if err == nil {
		defer jsonFile.Close()
		bytes, _ := ioutil.ReadAll(jsonFile)
		json.Unmarshal(bytes, &config)
	}

	config.MonsterSpawn = config.MonsterSpawn * time.Second
	g.Config = config
}
