package api

import (
	"database/sql"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func addPart(bot *tgbotapi.BotAPI, updates tgbotapi.UpdatesChannel) string {
	for update := range updates {
		if update.Message != nil {
			userInput := update.Message.Text
			db, err := sql.Open("sqlite3", "D:\\FromFlashCard\\FromLinux\\GO\\PTR_Bot\\DB_1_0.db")
			if err != nil {
				panic(err)
			}
			defer db.Close()

			_, err = db.Exec("INSERT INTO Parts (name) VALUES ($1)", userInput)
			if err != nil {
				if strings.Contains(strings.ToLower(err.Error()), "unique constraint") {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Такое название/номер детали уже существует!\nПожалуйста введите уникальное название/номер детали или внесите изменения в существующую деталь.")
					bot.Send(msg)
					continue
				} else {
					panic(err)
				}
			}
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Деталь успешно добавлена!\n Теперь вы можете указать информацию о ней:")
			msg.ReplyMarkup = replacePartButtons()
			bot.Send(msg)
			return "replacePart"

		}
	}

}
