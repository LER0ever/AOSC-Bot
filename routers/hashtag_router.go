package routers

import (
	"log"

	"github.com/LER0ever/AOSC-Bot/handlers"
	"github.com/LER0ever/AOSC-Bot/models"
	"gopkg.in/telegram-bot-api.v4"
)

// HashtagRouter handles every message that contains a #
func HashtagRouter(update tgbotapi.Update) {
	log.Printf("Hashtag Router [%s] %s %t\n", update.Message.From.UserName, update.Message.Text, models.IsPakreq(update.Message.Text))
	if models.IsPakreq(update.Message.Text) {
		handlers.PakreqHandler(update)
	}
}
