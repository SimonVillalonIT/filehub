package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Server      ServerConfig      `json:"server"`
	Credentials CredentialsConfig `json:"credentials"`
	DataFolder  string            `json:"data_folder"`
}

type ServerConfig struct {
	Port int `json:"port"`
}

type CredentialsConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Secret   string `json:"secret"`
}

func (c *Config) Load() error {
	file, err := os.Open("config.json")
	if err != nil {
		fmt.Println("Failed loading the config file")
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(c); err != nil {
		fmt.Println("Failed decoding the config file")
		return err
	}

	return nil
}
