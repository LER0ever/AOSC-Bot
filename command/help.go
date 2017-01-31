package command

import (
	"github.com/LER0ever/AOSC-Bot/bot"
	"gopkg.in/telegram-bot-api.v4"
)

const helpStr = "*AOSC Bot Help*\n*=============*\n" +
	"/help to show this help message\n" +
	"/board to show current open tasks\n" +
	"/version to show version and status\n" +
	"\n*Pakreq Help*\n*===========*\n" +
	"Your Pakreq should be like the following:\n" +
	"```\n#pakreq PKGNAME @ UPSTREAM-URL\n" +
	"Description starts from the second line and is optional```\n" +
	"\n*Update Help*\n*===========*\n" +
	"If you want to send a package update request, " +
	"please make sure it goes like this:\n" +
	"```\n#update PKGNAME to VERSION```\nor just\n" +
	"```\n#update PKGNAME```\n"

// Help replies with the help message
func Help(update tgbotapi.Update) {
	bot.ReplyMessage(update.Message.Chat.ID, helpStr, update.Message.MessageID)
}
