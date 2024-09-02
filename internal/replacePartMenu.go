package internal

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	//"database/sql"
)

func ReplacePart(bot *tgbotapi.BotAPI, user User) {
	update := <-user.userChanal
	if update.CallbackQuery != nil {
		switch update.CallbackQuery.Data {
		case "replaceName":
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Оно работает!")
			bot.Send(msg)
		}
	}
}

func replacePartButtons() tgbotapi.InlineKeyboardMarkup {
	var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Редактировать название", "lookPart"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Редактировать материал изделия", "replacePart"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Редактировать цену", "lookTool"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Редактировать время изготовления", "addTool"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Редактировать номер программы", "replaceTool"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Редактировать список инструментов", "replaceTool"),
		),
	)
	return numericKeyboard
}
