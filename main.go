package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/mymmrac/telego"

	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

// Структура для розкодування JSON-відповіді
type WeatherData struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

func main() {
	botToken := "token"

	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	updates, _ := bot.UpdatesViaLongPolling(nil)

	bh, _ := th.NewBotHandler(bot, updates)

	defer bh.Stop()
	defer bot.StopLongPolling()

	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)

		message := tu.Message(
			chatID,
			"Привіт! В цьому боті ви можете дізнаватись температуру повітр'я в різних містах України! Щоб почати введіть /та назву свого міста. Успіхів!",
		)

		bot.SendMessage(message)

	}, th.CommandEqual("start"))