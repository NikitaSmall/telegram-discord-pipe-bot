package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type DiscordConfig struct {
	Token          string `env:"DISCORD_TOKEN" env-required:"true"`
	CollectionName string `env:"COLLECTION_NAME"`
}

func (dc DiscordConfig) GetToken() string {
	return dc.Token
}

var discordConfig DiscordConfig

func GetDiscordCondig() DiscordConfig {
	return discordConfig
}

func initDiscord() {
	log.Println("parsing configuration for discord")
	if err := cleanenv.ReadEnv(&discordConfig); err != nil {
		log.Fatalf("error parsing configuration for discord: %e", err)
	}
}
