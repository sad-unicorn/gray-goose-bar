package main

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sad-unicorn/gray-goose-bar/db"
	"log"
	"os"
)

func main() {
	initDatabase()
	startBot()
}

func initDatabase() {
	dbUser := requireEnv("DATABASE_USER")
	dbPass := requireEnv("DATABASE_PASS")
	dbName := requireEnv("DATABASE_NAME")
	dbHost := requireEnv("DATABASE_HOST")
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPass, dbHost, dbName)
	db.InitDB(dbUrl)

	res, err := db.QueryForInt("SELECT 1")
	if err != nil {
		panic("Database is failed to `select 1`: " + err.Error())
	}

	fmt.Println("Db query: ", res)
}

func requireEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		panic("No " + key + " environment variable is set")
	}
	return val
}

func startBot() {
	token := requireEnv("BOT_TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
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
        }
	}
}
