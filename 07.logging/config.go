package main

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App struct {
		Name string `yaml:"name"`
	} `yaml:"app"`

	Validation struct {
		MinAge int `yaml:"min_age"`
	} `yaml:"validation"`

	Logging struct {
		Level string `yaml:"level"`
	} `yaml:"logging"`
}

func LoadConfig() (*Config, error) {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}
	var cfg Config
	err = yaml.Unmarshal(data, &cfg)

	return &cfg, err
}
