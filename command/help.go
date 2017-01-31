package command

import (
	"github.com/LER0ever/AOSC-Bot/bot"
	"gopkg.in/telegram-bot-api.v4"
)

const HelpStr = "*AOSC Bot Help*\n*=============*\n" +
	"/help to show this help message\n" +
	"/board to show current open tasks\n" +
	"\n*Pakreq Help*\n*===========*\n" +
	"Your Pakreq should be like the following, (replace & with #):\n" +
	"```\n&pakreq PKGNAME @ UPSTREAM-URL\n" +
	"Description starts from the second line and is optional```\n" +
	"\n*Update Help*\n*===========*\n" +
	"If you want to send a package update request, " +
	"please make sure it goes like this: (replace & with #)\n" +
	"```\n&update PKGNAME to VERSION```\nor just\n" +
	"```\n&update PKGNAME```\n"

func Help(update tgbotapi.Update) {
	bot.ReplyMessage(update.Message.Chat.ID, HelpStr, update.Message.MessageID)
}
