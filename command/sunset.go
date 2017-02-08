package command

import (
	"os"

	"github.com/LER0ever/AOSC-Bot/bot"
	"gopkg.in/telegram-bot-api.v4"
)

const SunsetStr = "Sir, that's my shutdown command. "

func Sunset(update tgbotapi.Update) {
	replymsg := SunsetStr
	if update.Message.From.UserName == "LER0ever" {
		replymsg += "Shutting down..."
		bot.ReplyMessage(update.Message.Chat.ID, replymsg, update.Message.MessageID)
		os.Exit(0)
	} else {
		replymsg += "Warning: you don't have that clearance. @LER0ever"
		bot.ReplyMessage(update.Message.Chat.ID, replymsg, update.Message.MessageID)
	}
}
