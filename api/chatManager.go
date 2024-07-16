package api

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ChatManager(bot *tgbotapi.BotAPI) {

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	//мониторим чат на входящие сообщения, если что-то пишут то в ответ отправляем своё сообщение
	for update := range updates {
		if update.Message == nil {
			//мониторим на нажатия кнопок
			if update.CallbackQuery == nil {
				continue
			}

			//callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			//bot.Request(callback)

			switch update.CallbackQuery.Data {
			case "lookPart":
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Перехожу в меню просмотра деталей")
				msg.ReplyMarkup = partLookButtons()
				bot.Send(msg)
				partLook(bot, updates)
			case "addPart":
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Перехожу в меню добавления детали.\nВведите название/номер детали.")
				bot.Send(msg)
				update.CallbackQuery.Data = addPart(bot, updates)
			case "replacePart":
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "К сожалению эта функция пока не реализована.")
				bot.Send(msg)
				replacePart(bot, updates)
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

			continue
		}

		userInput := update.Message.Text

		switch userInput {
		case "/start":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Рады вас приветствовать в главном меню бота PTR33!\nВыберите необходимую вам функцию.")
			msg.ReplyMarkup = greeting()
			bot.Send(msg)

		}

	}

}
