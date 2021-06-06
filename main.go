package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"telegram-discord-pipe-bot/cmd"
	"telegram-discord-pipe-bot/config"
)

func main() {
	log.Println("start pipe bot")
	discordTerminate, _ := terminateSignals()

	cmd.DiscordStart(discordTerminate, config.GetDiscordCondig())
}

// TODO: add telegram signal too
func terminateSignals() (chan os.Signal, chan os.Signal) {
	discordTerminate := make(chan os.Signal)
	signal.Notify(discordTerminate, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	return discordTerminate, nil
}
