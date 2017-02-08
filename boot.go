package main

import (
	"log"
	"strings"

	"github.com/LER0ever/AOSC-Bot/bot"
	"github.com/LER0ever/AOSC-Bot/github"
	"github.com/LER0ever/AOSC-Bot/models"
	"github.com/LER0ever/AOSC-Bot/routers"
	"gopkg.in/telegram-bot-api.v4"
)

// BootBot Boots the Robot.
func BootBot() {

	bot.Bot.Debug = true

	log.Printf("Authorized on account %s", bot.Bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, _ := bot.Bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		addToMemberList(update.Message.From.UserName)
		if update.Message.ReplyToMessage != nil {
			routers.ReplyRouter(update)
		} else {
			if strings.Contains(update.Message.Text, `#`) {
				routers.HashtagRouter(update)
			}
			if update.Message.IsCommand() {
				routers.CommandRouter(update)
			}
		}
	}
}

// InitGHClient inits the git client
func InitGHClient() {
	github.InitClient()
}

func addToMemberList(userID string) {
	for _, m := range models.CurConfig.MembersID {
		if m == userID {
			return
		}
	}
	models.CurConfig.MembersID = append(models.CurConfig.MembersID, userID)
	log.Printf("addToMemberList: %s", userID)
	models.ExportToFile()
}
