package handlers

import (
	"log"
	"time"

	"github.com/LER0ever/AOSC-Bot/bot"
	"github.com/LER0ever/AOSC-Bot/models"
	"gopkg.in/telegram-bot-api.v4"
)

// PakreqHandler comment
func PakreqHandler(update tgbotapi.Update) {
	log.Printf("PakreqHandler")
	var replymsg string
	arrpakreq := models.PakreqToArray(update.Message.Text)
	replymsg = "Pakreq created for \nPackage: " + arrpakreq[0] + "\nUpstream: " + arrpakreq[2] +
		"\nAuthor: " + update.Message.From.UserName + "\nStartTime: " + update.Message.Time().Format(time.RFC1123)
	bot.ReplyMessage(update.Message.Chat.ID, replymsg, update.Message.MessageID)
}
