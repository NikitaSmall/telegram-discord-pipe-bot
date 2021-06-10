package discord

import (
	"fmt"
	"log"
	"strings"
	"telegram-discord-pipe-bot/models"

	"github.com/bwmarrin/discordgo"
)

func (ds DiscordSession) handleMessage(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.User.ID {
		return
	}

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

	channel, err := session.Channel(message.ChannelID)
	if err != nil {
		session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("cannot find channel to register %s with err %s", message.ChannelID, err.Error()))
	}

	if err := ds.commStorage.Register(message.ChannelID, channel.Name); err != nil {
		session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("cannot register channel %s with err %s", message.ChannelID, err.Error()))
	}

	session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("channel %s registered successfully for proxying", channel.Name))
}

func (ds DiscordSession) handleUnregister(session *discordgo.Session, message *discordgo.MessageCreate) {
	log.Printf("proxy unregister %s", message.ChannelID)

	if err := ds.commStorage.Unregister(message.ChannelID); err != nil {
		session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("cannot unregister channel %s with err %s", message.ChannelID, err.Error()))
	}

	session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("channel %s unregistered successfully for proxying", message.ChannelID))
}

func (ds DiscordSession) handleRegularMessage(session *discordgo.Session, message *discordgo.MessageCreate) {
	log.Printf("proxy regular message %s", message.ChannelID)

	isRegistered, err := ds.commStorage.IsRegistered(message.ChannelID)
	if err != nil {
		log.Printf("cannot extract if the message is from registered channel with err: %e", err)
		return
	}

	if !isRegistered {
		log.Printf("get message from unregistered channel, exiting")
		return
	}

	ds.FromDiscordChan <- models.Message{
		Author:  message.Author.Username,
		Message: message.Content,
	}
}
