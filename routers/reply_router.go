package routers

import (
	"log"

	"gopkg.in/telegram-bot-api.v4"
)

// ReplyRouter handles every reply message
func ReplyRouter(update tgbotapi.Update) {
	log.Printf("[%s] %s :: %s\n", update.Message.From.UserName, update.Message.Text, update.Message.ReplyToMessage.Text)
}
