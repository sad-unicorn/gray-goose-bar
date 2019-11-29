package router

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sad-unicorn/gray-goose-bar/cwapi"
	"log"
)


var bot *tgbotapi.BotAPI

func StartBot(token string) {
	var err error
	bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}
		if int64(update.Message.From.ID) == update.Message.Chat.ID {
			_, _ = bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "pong"))
			cwapi.Publish()
			continue
		}

		Dispatch(update.Message.From.ID, update.Message.Text)
	}
}
