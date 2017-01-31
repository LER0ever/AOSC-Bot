package routers

import (
	"log"

	"github.com/LER0ever/AOSC-Bot/handlers"
	"gopkg.in/telegram-bot-api.v4"
)

// CommandRouter handles every message that starts with /
func CommandRouter(update tgbotapi.Update) {
	log.Printf("Command Router [%s] %s\n", update.Message.From.UserName, update.Message.Text)
	handlers.CommandHandler(update)
}
