package telegram

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (ts TelegramSession) ListenToMessages() {
	for {
		select {
		case message := <-ts.ToTelegramChan:
			channels, err := ts.commStorage.GetCommunicationList()
			if err != nil {
				log.Printf("cannot find registered channel for telegram, err: %e", err)
				continue
			}

			for _, channel := range channels {
				chanID, err := strconv.ParseInt(channel.ChannelID, 10, 64)
				if err != nil {
					log.Printf("cannot parse channel id for telegram, err: %e", err)
					continue
				}
				ts.Bot.Send(tgbotapi.NewMessage(chanID, fmt.Sprintf("%s: %s", message.Author, message.Message)))
			}
		}
	}
}
