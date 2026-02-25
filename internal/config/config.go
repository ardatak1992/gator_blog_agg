package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const filename = ".gatorconfig.json"

// Config struct
type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

// Read reads the json file at the home directory and returns a Config struct
func Read() (Config, error) {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	configFile, err := os.Open(configFilePath)
	if err != nil {
		return Config{}, err
	}
	defer configFile.Close()

	cfg := Config{}
	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}

// SetUser sets user in the config file at home directory.
func (c *Config) SetUser(username string) error {

	c.CurrentUserName = username

	err := write(*c)
	if err != nil {
		return err
	}

	return nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("can't get the home directory: %s", err.Error())
	}

	return homeDir + "/" + filename, nil
}

func write(cfg Config) error {

	configFilePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	configFile, err := os.OpenFile(configFilePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer configFile.Close()

	encoder := json.NewEncoder(configFile)

	if err = encoder.Encode(cfg); err != nil {
		return err
	}

	return nil
}
