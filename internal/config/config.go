package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	configPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	data, err := os.ReadFile(configPath)
	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return Config{}, fmt.Errorf("Unable to decode to a json:%s\n Error:%w", string(data), err)
	}

	return config, nil
}

func (c *Config) SetUser(username string) error {
	c.CurrentUserName = username
	err := write(*c)
	if err != nil {
		return err
	}
	return nil
}

func getConfigFilePath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Unable to read the home path")
	}

	configPath := filepath.Join(homePath, configFileName)
	return configPath, nil
}

func write(cfg Config) error {
	dat, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("Unable to decode json %w", err)
	}

	configPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	err = os.WriteFile(configPath, dat, 0644)

	if err != nil {
		return fmt.Errorf("Unable to write to file %s\n%w", configPath, err)
	}

	return nil
}
