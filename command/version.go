package command

import (
	"runtime"

	"github.com/LER0ever/AOSC-Bot/bot"
	"github.com/LER0ever/AOSC-Bot/models"
	"gopkg.in/telegram-bot-api.v4"
)

var versionStr = "*AOSC Bot*\n========\n" +
	"Version: " + models.CurConfig.Version +
	"\nGo Runtime: " + runtime.Version() +
	"\nArch: " + runtime.GOARCH +
	"\nOS: " + runtime.GOOS

// Version replies the status info of the bot
func Version(update tgbotapi.Update) {
	bot.ReplyMessage(update.Message.Chat.ID, versionStr, update.Message.MessageID)
}
