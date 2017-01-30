package routers

import (
	"log"

	"gopkg.in/telegram-bot-api.v4"
)

// CommandRouter handles every message that starts with /
func CommandRouter(update tgbotapi.Update) {
	log.Printf("Command Router [%s] %s\n", update.Message.From.UserName, update.Message.Text)
}
