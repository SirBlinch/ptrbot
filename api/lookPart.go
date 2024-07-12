package api

import (
	"database/sql"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func partLookButtons() tgbotapi.InlineKeyboardMarkup {
	var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Просмотр списка деталей", "PartList"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Посмотреть конкретную деталь", "SwitchPart"),
		),
	)
	return numericKeyboard
}

func partLook(bot *tgbotapi.BotAPI, updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.CallbackQuery != nil {
			switch update.CallbackQuery.Data {
			case "PartList":
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Смотрим список деталей!\n Вот такие детали у нас есть:")
				bot.Send(msg)
				db, err := sql.Open("sqlite3", "D:\\FromFlashCard\\FromLinux\\GO\\PTR_Bot\\DB_1_0.db")
				if err != nil {
					panic(err)
				}
				defer db.Close()
				rows, err := db.Query("SELECT name FROM Parts")
				if err != nil {
					panic(err)
				}
				defer rows.Close()
				for rows.Next() {
					var value string
					err := rows.Scan(&value)
					if err != nil {
						panic(err)
					}
					msg = tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, value)
					bot.Send(msg)
				}

			case "SwitchPart":
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Смотрим конкретную деталь! \n Введите название детали.")
				bot.Send(msg)
				for update := range updates {
					if update.Message != nil {
						value := update.Message.Text
						db, err := sql.Open("sqlite3", "D:\\FromFlashCard\\FromLinux\\GO\\PTR_Bot\\DB_1_0.db")
						if err != nil {
							panic(err)
						}
						defer db.Close()

						partDataRow, err := db.Query("SELECT * FROM Parts WHERE name = $1", value)
						if err != nil {
							panic(err)
						}
						defer partDataRow.Close()

						columns, err := partDataRow.Columns()
						if err != nil {
							panic(err)
						}

						values := make([]interface{}, len(columns))
						_values := make([]interface{}, len(columns))
						for i := range columns {
							_values[i] = &values[i]
						}
						for partDataRow.Next() {
							err := partDataRow.Scan(_values...)
							if err != nil {
								panic(err)
							}
						}

						if values[0] == nil {
							msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Такой детали у нас нет, попробуйте еще раз!")
							bot.Send(msg)
							continue
						}

						for i, column := range columns {
							value, noerr := values[i].(string)
							if !noerr {
								value = strconv.FormatInt(values[i].(int64), 10)
							}
							//var output string =  column + ": " + value
							msg := tgbotapi.NewMessage(update.Message.Chat.ID, column+": "+value)
							bot.Send(msg)
						}
						break
					}
				}
			}
		}
	}

}
