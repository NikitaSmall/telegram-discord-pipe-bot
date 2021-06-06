package discord

import (
	"log"
	"telegram-discord-pipe-bot/interfaces"

	"github.com/bwmarrin/discordgo"
)

func Start(config interfaces.BotConfiger) *discordgo.Session {
	dg, err := discordgo.New("Bot " + config.GetToken())
	if err != nil {
		log.Fatalf("error initializing discord session with err: %e", err)
	}

	dg.AddHandler(handleMessage)
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		log.Fatalf("error opening discord session connection: %e", err)
	}

	log.Println("discord bot is ready to handle messages")

	return dg
}
