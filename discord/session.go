package discord

import (
	"log"
	"telegram-discord-pipe-bot/interfaces"
	"telegram-discord-pipe-bot/models"

	"github.com/bwmarrin/discordgo"
)

type DiscordSession struct {
	DG *discordgo.Session

	config      interfaces.BotConfiger
	commStorage interfaces.CommunicationStorager

	FromDiscordChan chan models.Message
	ToDiscordChan   chan models.Message
}

func New(config interfaces.BotConfiger, commStorage interfaces.CommunicationStorager, fromDiscord, toDiscord chan models.Message) *DiscordSession {
	dg, err := discordgo.New("Bot " + config.GetToken())
	if err != nil {
		log.Fatalf("error initializing discord session with err: %e", err)
	}

	discordSession := &DiscordSession{
		DG:              dg,
		config:          config,
		commStorage:     commStorage,
		FromDiscordChan: fromDiscord,
		ToDiscordChan:   toDiscord,
	}

	dg.AddHandler(discordSession.handleMessage)
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	return discordSession
}

func (ds DiscordSession) Start() {
	if err := ds.DG.Open(); err != nil {
		log.Fatalf("error opening discord session connection: %e", err)
	}

	log.Println("discord bot is ready to handle messages")
}
