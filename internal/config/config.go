package config

import (
	"log"
	"unicode/utf8"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	BotToken    string `yaml:"bot_token"`
	GithubToken string `yaml:"github_access_token"`
}

// MustConfig
func MustConfig(path string) *Config {
	var config Config

	err := cleanenv.ReadConfig(path, &config)
	if err != nil {
		log.Fatal("ConfigReader", err)
	}

	if utf8.RuneCountInString(config.BotToken) == 0 || utf8.RuneCountInString(config.GithubToken) == 0 {
		log.Fatal("token is nil")
	}

	return &config
}
