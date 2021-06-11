package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type TelegramConfig struct {
	Token          string `env:"TELEGRAM_TOKEN" env-required:"true"`
	CollectionName string `env:"COLLECTION_NAME"`
}

func (dc TelegramConfig) GetToken() string {
	return dc.Token
}

var telegramConfig TelegramConfig

func GetTelegramConfig() TelegramConfig {
	return telegramConfig
}

func initTelegram() {
	log.Println("parsing configuration for telegram")
	if err := cleanenv.ReadEnv(&telegramConfig); err != nil {
		log.Fatalf("error parsing configuration for telegram: %e", err)
	}
}
