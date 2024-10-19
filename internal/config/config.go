package config

import (
	"log"
	"os"

	"github.com/go-yaml/yaml"
	"github.com/joho/godotenv"
)


func MustLoad() Config {

	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}
	configFile, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}

	var config Config

	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatal(err)
	}
	return config
}
