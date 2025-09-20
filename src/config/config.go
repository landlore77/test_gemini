package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	GoogleClientID string `yaml:"GOOGLE_CLIENT_ID"`
	MySQL          struct {
		ID   string `yaml:"ID"`
		PASS string `yaml:"PASS"`
	} `yaml:"MYSQL"`
}

var Cfg Config

func LoadConfig() {
	log.Println("Loading configuration...")
	f, err := os.Open("src/config.yaml")
	if err != nil {
		log.Fatalf("Failed to open config file: %v", err)
	}
	defer f.Close()

	log.Println("Config file opened successfully.")
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&Cfg)
	if err != nil {
		log.Fatalf("Failed to decode config file: %v", err)
	}
	log.Println("Configuration loaded successfully.")
}
