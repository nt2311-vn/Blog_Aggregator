package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = "gatorconfig.json"

type Config struct {
	DbUrl    string `json:"db_url"`
	Username string `json:"current_user_name"`
}

func (c *Config) SetUser() {
	c.Username = "nt2311"

	byteData, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}

	if err = os.WriteFile(configFileName, byteData, 0666); err != nil {
		panic(err)
	}
}

func Read() Config {
	byteData, err := os.ReadFile(filepath.Join(configFileName))
	if err != nil {
		return Config{}
	}

	var config Config

	if err = json.Unmarshal(byteData, &config); err != nil {
		return Config{}
	}

	return config
}
