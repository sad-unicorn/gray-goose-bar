package router

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
	"strconv"
	"strings"
)


var barChatId int64 = -350912043

func init() {
	replacementBarChatId, _ := strconv.ParseInt(os.Getenv("BAR_CHAT_ID"), 10, 64)
	if replacementBarChatId != 0 {
		barChatId = replacementBarChatId
	}
}

func Dispatch(userId int, message string) {
	log.Println(userId, ":", message)
	switch {
	case strings.HasPrefix(message, "/hi"):
		messageToBar("Hello!")
	}
}

func messageToBar(text string) {
	message := tgbotapi.NewMessage(barChatId, text)
	message.ParseMode = "markdown"
	_,_ = bot.Send(message)
}