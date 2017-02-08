package command

import (
	"github.com/LER0ever/AOSC-Bot/bot"
	"github.com/LER0ever/AOSC-Bot/models"
	"gopkg.in/telegram-bot-api.v4"
)

var AtAllStr = "Listen you all!\n"

func AtAll(update tgbotapi.Update) {
	for _, m := range models.CurConfig.MembersID {
		AtAllStr += "@" + m + " "
	}
	bot.ReplyMessage(update.Message.Chat.ID, AtAllStr, update.Message.MessageID)
}
