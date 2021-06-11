package telegram

import (
	"fmt"
	"log"
	"strings"
	"telegram-discord-pipe-bot/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (ts TelegramSession) handleMessage(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		switch {
		case strings.HasPrefix(update.Message.Text, "/proxy-register"):
			ts.handleRegister(update)
		case strings.HasPrefix(update.Message.Text, "/proxy-unregister"):
			ts.handleUnregister(update)
		default:
			ts.handleDefaultMessage(update)
		}
	}
}

func (ts TelegramSession) handleRegister(update tgbotapi.Update) {
	log.Printf("proxy register telegram %s", idToString(update.Message.Chat.ID))

	err := ts.commStorage.Register(idToString(update.Message.Chat.ID), update.Message.Chat.Title)
	if err != nil {
		ts.Bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("unable to register telegram channel with err: %e", err)))
	}

	ts.Bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("channel %s  is registed successfully", update.Message.Chat.Title)))
}

func (ts TelegramSession) handleUnregister(update tgbotapi.Update) {
	log.Printf("proxy unregister telegram %s", idToString(update.Message.Chat.ID))

	err := ts.commStorage.Unregister(idToString(update.Message.Chat.ID))
	if err != nil {
		ts.Bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("unable to unregister telegram channel with err: %e", err)))
	}

	ts.Bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("channel %s  is unregisted successfully", update.Message.Chat.Title)))
}

func (ts TelegramSession) handleDefaultMessage(update tgbotapi.Update) {
	isRegistered, err := ts.commStorage.IsRegistered(idToString(update.Message.Chat.ID))
	if err != nil {
		log.Printf("cannot extract if the message is from telegram registered channel with err: %e", err)
		return
	}

	if !isRegistered {
		log.Printf("get message from unregistered telegram channel, exiting")
		return
	}

	ts.FromTelegramChan <- models.Message{
		Author:  update.Message.From.UserName,
		Message: update.Message.Text,
	}
}

func idToString(id int64) string {
	return fmt.Sprintf("%d", id)
}
