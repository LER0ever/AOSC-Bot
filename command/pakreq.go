package command

import (
	"github.com/LER0ever/AOSC-Bot/bot"
	"gopkg.in/telegram-bot-api.v4"
)

const PakreqStr = "Sorry I don't take /pakreq, please use #pakreq instead or consult @LER0ever"

func Pakreq(update tgbotapi.Update) {
	bot.ReplyMessage(update.Message.Chat.ID, PakreqStr, update.Message.MessageID)
}
