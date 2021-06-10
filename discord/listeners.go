package discord

import (
	"fmt"
	"log"
)

func (ds DiscordSession) ListenToMessages() {
	for {
		select {
		case message := <-ds.ToDiscordChan:
			channels, err := ds.commStorage.GetCommunicationList()
			if err != nil {
				log.Printf("cannot find registered channel for discord, err: %e", err)
				continue
			}

			for _, channel := range channels {
				ds.DG.ChannelMessageSend(channel.ChannelID, fmt.Sprintf("%s: %s", message.Author, message.Message))
			}
		}
	}
}
