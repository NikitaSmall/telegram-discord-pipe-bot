package cmd

import (
	"log"
	"os"
	"telegram-discord-pipe-bot/interfaces"
	"telegram-discord-pipe-bot/models"
	"telegram-discord-pipe-bot/storage"
	"telegram-discord-pipe-bot/telegram"
)

func TelegramStart(terminateSignal chan os.Signal, botConfig interfaces.BotConfiger, discordChan, telegramChan chan models.Message) {
	log.Println("start telegram bot initialize process")

	inMemoryStorage := storage.NewInmemory()
	inMemoryStorage.PopulateCommunicationList()
	ts := telegram.New(botConfig, inMemoryStorage, telegramChan, discordChan)
	go ts.Start()
	go ts.ListenToMessages()

	<-terminateSignal
}
