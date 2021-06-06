package cmd

import (
	"log"
	"os"
)

func DiscordStart(terminateSignal chan os.Signal) {
	log.Println("start discord bot initialize process")
	<-terminateSignal
}
