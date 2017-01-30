package bot

import (
	"os"

	"gopkg.in/telegram-bot-api.v4"
)

// Bot global variable
var Bot, _ = tgbotapi.NewBotAPI(os.Getenv("AOSC_BOT_KEY"))

// SendMessage sends the message
func SendMessage(ChatID int64, Msg string) {
	msg := tgbotapi.NewMessage(ChatID, Msg)

	Bot.Send(msg)
}

// ReplyMessage replies the message
func ReplyMessage(ChatID int64, Msg string, ReplyMsgID int) {
	msg := tgbotapi.NewMessage(ChatID, Msg)
	msg.ReplyToMessageID = ReplyMsgID
	Bot.Send(msg)
}
