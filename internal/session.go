package internal

import (
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Session(user User, bot *tgbotapi.BotAPI) {
	timer := time.NewTimer(10 * time.Minute)
	defer timer.Stop()
	update := <-user.userChanal
	if update.Message != nil {

		userInput := update.Message.Text

		switch userInput {
		case "/start":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Рады вас приветствовать в главном меню бота PTR33!\nВыберите необходимую вам функцию.")
			msg.ReplyMarkup = Greeting()
			bot.Send(msg)
		}
	}
	// Проверяем на нажатие кнопок
	if update.CallbackQuery != nil {

		switch update.CallbackQuery.Data {
		case "lookPart":
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Перехожу в меню просмотра деталей")
			msg.ReplyMarkup = PartLookButtons()
			bot.Send(msg)
			PartLook(bot, user)
		case "addPart":
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Перехожу в меню добавления детали.\nВведите название/номер детали.")
			bot.Send(msg)
			//update.CallbackQuery.Data = addPart(bot, updates)
		case "replacePart":
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "К сожалению эта функция пока не реализована.")
			bot.Send(msg)
			ReplacePart(bot, user)
		case "lookTool":
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "К сожалению эта функция пока не реализована.")
			bot.Send(msg)
		case "addTool":
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "К сожалению эта функция пока не реализована.")
			bot.Send(msg)
		case "replaceTool":
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "К сожалению эта функция пока не реализована.")
			bot.Send(msg)
		}

	}
}
