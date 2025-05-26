package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	DbUrl    *string `json:"db_url"`
	CurrUser *string `json:"current_user_name"`
}

func Read() Config {
	homePath, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	configPath := homePath + "/.gatorconfig.json"

	fileBody, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}

	config := Config{}
	err = json.Unmarshal(fileBody, &config)
	if err != nil {
		log.Fatal(err)
	}

	return config
}

func SetUser(user string) string {
	config := Read()
	config.CurrUser = &user

	homePath, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	configPath := homePath + "/.gatorconfig.json"
	if err = os.Truncate(configPath, 0); err != nil {
		log.Fatal(err)
	}

	dataToWrite, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	if err = os.Truncate(configPath, 0); err != nil {
		log.Fatal(err)
	}

	os.WriteFile(configPath, dataToWrite, 0644)
	return string(dataToWrite)
}
