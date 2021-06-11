package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"telegram-discord-pipe-bot/cmd"
	"telegram-discord-pipe-bot/config"
	"telegram-discord-pipe-bot/models"
)

func main() {
	log.Println("start pipe bot")
	discordTerminate, telegramTerminate := terminateSignals()

	discordChan := make(chan models.Message)
	telegramChan := make(chan models.Message)

	go cmd.DiscordStart(discordTerminate, config.GetDiscordCondig(), discordChan, telegramChan)
	cmd.TelegramStart(telegramTerminate, config.GetTelegramConfig(), discordChan, telegramChan)
}

// TODO: add telegram signal too
func terminateSignals() (chan os.Signal, chan os.Signal) {
	discordTerminate := make(chan os.Signal)
	telegramTerminate := make(chan os.Signal)
	signal.Notify(discordTerminate, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	signal.Notify(telegramTerminate, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	return discordTerminate, telegramTerminate
}
