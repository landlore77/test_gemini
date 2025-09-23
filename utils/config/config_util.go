package config

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	MYSQL struct {
		ID   string `yaml:"ID"`
		PASS string `yaml:"PASS"`
	} `yaml:"MYSQL"`
}

var AppConfig Config

func LoadConfig(configPath string) {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	err = yaml.Unmarshal(data, &AppConfig)
	if err != nil {
		log.Fatalf("Error unmarshaling config: %v", err)
	}
}
