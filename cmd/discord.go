package cmd

import (
	"log"
	"os"
	"telegram-discord-pipe-bot/discord"
	"telegram-discord-pipe-bot/interfaces"
	"telegram-discord-pipe-bot/models"
	"telegram-discord-pipe-bot/storage"
)

func DiscordStart(terminateSignal chan os.Signal, botConfig interfaces.BotConfiger, discordChan, telegramChan chan models.Message) {
	log.Println("start discord bot initialize process")

	storage := storage.NewFirestore("discord")
	storage.PopulateCommunicationList()
	ds := discord.New(botConfig, storage, discordChan, telegramChan)

	ds.Start()
	defer func() {
		if err := ds.DG.Close(); err != nil {
			log.Printf("error closing discord bot session: %e", err)
		}
	}()

	go ds.ListenToMessages()

	<-terminateSignal
}
