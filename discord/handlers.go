package discord

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func (ds DiscordSession) handleMessage(session *discordgo.Session, message *discordgo.MessageCreate) {
	log.Printf("handling discord message %s", message.ChannelID)

	switch {
	case strings.HasPrefix(message.Content, "/proxy-register"):
		ds.handleRegister(session, message)
	case strings.HasPrefix(message.Content, "/proxy-unregister"):
		ds.handleUnregister(session, message)
	default:
		ds.handleRegularMessage(session, message)
	}
}

func (ds DiscordSession) handleRegister(session *discordgo.Session, message *discordgo.MessageCreate) {
	log.Printf("proxy register %s", message.ChannelID)
}

func (ds DiscordSession) handleUnregister(session *discordgo.Session, message *discordgo.MessageCreate) {
	log.Printf("proxy unregister %s", message.ChannelID)
}

func (ds DiscordSession) handleRegularMessage(session *discordgo.Session, message *discordgo.MessageCreate) {
	log.Printf("proxy regular message %s", message.ChannelID)
}
