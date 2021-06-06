package cmd

import (
	"log"
	"os"
	"telegram-discord-pipe-bot/discord"
	"telegram-discord-pipe-bot/interfaces"
)

func DiscordStart(terminateSignal chan os.Signal, botConfig interfaces.BotConfiger) {
	log.Println("start discord bot initialize process")

	ds := discord.Start(botConfig)
	defer func() {
		if err := ds.Close(); err != nil {
			log.Printf("error closing discord bot session: %e", err)
		}
	}()

	<-terminateSignal
}
