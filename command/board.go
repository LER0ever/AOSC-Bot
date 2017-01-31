package command

import (
	"fmt"

	"github.com/LER0ever/AOSC-Bot/bot"
	"github.com/LER0ever/AOSC-Bot/models"
	"gopkg.in/telegram-bot-api.v4"
)

const boardPreStr = "*Todo Board*\n*==========*\n"
const boardPakreqTitle = "_Pakreqs_\n--------------\n"
const boardUpdateTitle = "_Updates_\n--------------\n"
const boardSecurityTitle = "_Securities_\n--------------------\n"

// Board replies the current todo board
func Board(update tgbotapi.Update) {
	replymsg := boardPreStr
	if len(models.CurConfig.CurPakreqs) != 0 {
		replymsg += boardPakreqTitle
		for i, pkg := range models.CurConfig.CurPakreqs {
			replymsg += fmt.Sprintf("%d", i+1) + ">" + pkg.PkgName + ": [Upstream](" + pkg.PkgURL + ") | [Issue](" + pkg.IssueURL + ")\n"
		}
	}
	if len(models.CurConfig.CurUpdates) != 0 {
		replymsg += boardUpdateTitle
		for i, pkg := range models.CurConfig.CurUpdates {
			replymsg += fmt.Sprintf("%d", i+1) + ">" + pkg.PkgName + ": [ToVersion](" + pkg.ToVersion + ")\n"
		}
	}
	bot.ReplyMessage(update.Message.Chat.ID, replymsg, update.Message.MessageID)
}
