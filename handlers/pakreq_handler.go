package handlers

import (
	"log"
	"strings"
	"time"

	"github.com/LER0ever/AOSC-Bot/bot"
	"github.com/LER0ever/AOSC-Bot/models"
	"gopkg.in/telegram-bot-api.v4"
)

// PakreqHandler comment
func PakreqHandler(update tgbotapi.Update) {
	log.Printf("PakreqHandler")

	Desc := ""
	Msg := update.Message.Text
	if strings.Contains(Msg, "\n") {
		Desc = strings.Join(strings.Split(Msg, "\n")[1:], "\n")
		Msg = strings.Split(Msg, "\n")[0]
	}

	var replymsg string
	arrpakreq := models.PakreqToArray(Msg)

	var pakreq models.Pakreq
	pakreq.PkgName = arrpakreq[0]
	pakreq.PkgURL = arrpakreq[2]
	pakreq.PkgDesc = Desc
	pakreq.Author = update.Message.From.UserName
	pakreq.OpenTime = update.Message.Time()

	succ := models.AddPakreq(pakreq)

	if succ == true {
		replymsg = "Pakreq created for \nPackage: " + pakreq.PkgName + "\nUpstream: " + pakreq.PkgURL +
			"\nAuthor: " + pakreq.Author + "\nStartTime: " + pakreq.OpenTime.Format(time.RFC1123) +
			"\nDescription:\n" + pakreq.PkgDesc
	} else {
		replymsg = "Pakreq failed to create. Maybe there is already a same one opened."
	}

	bot.ReplyMessage(update.Message.Chat.ID, replymsg, update.Message.MessageID)
}
