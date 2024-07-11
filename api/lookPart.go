package api

import (
	"database/sql"

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
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Смотрим список деталей!")
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

					columnsTitles, err := db.Query("PRAGMA TABLE_INFO(Parts)")
					if err != nil {
						panic(err)
					}
					defer columnsTitles.Close()

					for columnsTitles.Next() {
						a := 1
						var columnTitle string
						var partData string
						err := columnsTitles.Scan(&columnTitle)
						if err != nil {
							panic(err)
						}
						err := partDataRow[a].Scan(&partData)
						if err != nil {
							panic(err)
						}
					}

					for rows.Next() {
						var dbValue string
						err := rows.Scan(&dbValue)
						if err != nil {
							panic(err)
						}

					}
				}
			}
		}
	}

}
