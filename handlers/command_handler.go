package handlers

import (
	"log"

	"github.com/LER0ever/AOSC-Bot/command"
	"gopkg.in/telegram-bot-api.v4"
)

// CommandHandler handles the command
func CommandHandler(update tgbotapi.Update) {
	log.Printf("Command Handler")
	cmd := update.Message.Command()
	log.Println(cmd)

	switch cmd {
	case "help":
		command.Help(update)
	case "pakreq":
		command.Pakreq(update)
	case "board":
		command.Board(update)
	case "version":
		command.Version(update)
	case "sunset":
		command.Sunset(update)
	case "atall":
		command.AtAll(update)
	}
}
