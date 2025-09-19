package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	GoogleClientID string `yaml:"GOOGLE_CLIENT_ID"`
}

var Cfg Config

func LoadConfig() {
	f, err := os.Open("config.yaml")
	if err != nil {
		log.Fatalf("Failed to open config file: %v", err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&Cfg)
	if err != nil {
		log.Fatalf("Failed to decode config file: %v", err)
	}
}
