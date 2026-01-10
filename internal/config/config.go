package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func Read() (Config, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(filepath.Join(home, configFileName))
	if err != nil {
		return Config{}, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return Config{}, err
	}

	return config, nil
}

func (c *Config) SetUser(name string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	c.CurrentUserName = name

	file, err := os.Create(filepath.Join(home, configFileName))
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder((file))
	return encoder.Encode(c)
}
