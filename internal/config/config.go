package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	DBURL string `json:"db_url"`
}

func Read() Config {
	byteData, err := os.ReadFile(filepath.Join("gatorconfig.json"))
	if err != nil {
		return Config{}
	}

	var config Config

	if err = json.Unmarshal(byteData, &config); err != nil {
		return Config{}
	}

	return config
}
