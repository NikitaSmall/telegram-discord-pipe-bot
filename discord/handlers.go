package discord

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func handleMessage(session *discordgo.Session, message *discordgo.MessageCreate) {
	log.Printf("handling discord message %s", message.ChannelID)
}
