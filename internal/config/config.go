package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

func Read() Config {
	configPath, err := getConfigFilePath()
	if err != nil {
		fmt.Println("Error:", err)
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Println("Error:", err)
	}

	var config Config

	if err = json.Unmarshal(data, &config); err != nil {
		fmt.Println("Error:", err)
	}
	return config
}

func getConfigFilePath() (string, error) {
	homepath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homepath, configFileName), nil
}

func write(c Config) error {
	configPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	data, err := json.Marshal(c)
	if err != nil {
		return err
	}

	os.WriteFile(configPath, data, 0644)
	return nil
}

func (c *Config) SetUser(username string) error {
	c.CurrentUserName = username
	err := write(*c)
	if err != nil {
		return err
	}
	return nil
}

func (c Config) GetDatabaseURL() string {
	return c.DBUrl
}
