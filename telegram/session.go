package telegram

import (
	"log"
	"telegram-discord-pipe-bot/interfaces"
	"telegram-discord-pipe-bot/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type TelegramSession struct {
	Bot *tgbotapi.BotAPI

	config      interfaces.BotConfiger
	commStorage interfaces.CommunicationStorager

	FromTelegramChan chan models.Message
	ToTelegramChan   chan models.Message
}

func New(config interfaces.BotConfiger, commStorage interfaces.CommunicationStorager, fromTelegram, toTelegram chan models.Message) TelegramSession {
	bot, err := tgbotapi.NewBotAPI(config.GetToken())
	if err != nil {
		log.Fatalf("unable to create telegram bot with err: %e", err)
	}

	return TelegramSession{
		Bot:              bot,
		config:           config,
		commStorage:      commStorage,
		FromTelegramChan: fromTelegram,
		ToTelegramChan:   toTelegram,
	}
}

func (ts TelegramSession) Start() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := ts.Bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatalf("unable to get telegram bot update chan with err: %e", err)
	}

	log.Println("start telegram bot handling messages")
	ts.handleMessage(updates)
}
