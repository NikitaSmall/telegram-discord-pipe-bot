package discord

import (
	"log"
	"telegram-discord-pipe-bot/interfaces"

	"github.com/bwmarrin/discordgo"
)

type DiscordSession struct {
	DG     *discordgo.Session
	config interfaces.BotConfiger
}

func Start(config interfaces.BotConfiger) *DiscordSession {
	dg, err := discordgo.New("Bot " + config.GetToken())
	if err != nil {
		log.Fatalf("error initializing discord session with err: %e", err)
	}

	discordSession := &DiscordSession{
		DG:     dg,
		config: config,
	}

	dg.AddHandler(discordSession.handleMessage)
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		log.Fatalf("error opening discord session connection: %e", err)
	}

	log.Println("discord bot is ready to handle messages")

	return discordSession
}
