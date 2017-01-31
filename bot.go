package main

import (
	"log"
	"strings"

	"github.com/LER0ever/AOSC-Bot/bot"
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
