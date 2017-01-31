package bot

import (
	"github.com/LER0ever/AOSC-Bot/models"
	"gopkg.in/telegram-bot-api.v4"
)

// Bot global variable
var Bot *tgbotapi.BotAPI

func init() {
	Bot, _ = tgbotapi.NewBotAPI(models.CurConfig.TgBotKey)
}

// SendMessage sends the message
func SendMessage(ChatID int64, Msg string) {
	msg := tgbotapi.NewMessage(ChatID, Msg)
	msg.ParseMode = tgbotapi.ModeMarkdown
	Bot.Send(msg)
}

// ReplyMessage replies the message
func ReplyMessage(ChatID int64, Msg string, ReplyMsgID int) {
	msg := tgbotapi.NewMessage(ChatID, Msg)
	msg.ParseMode = tgbotapi.ModeMarkdown
	msg.ReplyToMessageID = ReplyMsgID
	Bot.Send(msg)
}
