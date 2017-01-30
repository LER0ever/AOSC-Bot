package handlers

import (
	"log"

	"strings"

	"github.com/LER0ever/AOSC-Bot/bot"
	"github.com/LER0ever/AOSC-Bot/models"
	"gopkg.in/telegram-bot-api.v4"
)

// PakreqHandler comment
func PakreqHandler(update tgbotapi.Update) {
	log.Printf("PakreqHandler")
	var replymsg string
	replymsg = "Pakreq created for " + strings.Join(models.PakreqToPkgString(update.Message.Text), " ")
	bot.ReplyMessage(update.Message.Chat.ID, replymsg, update.Message.MessageID)
}
