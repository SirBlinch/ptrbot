package api

import (
	"log"
	"ptrbot/internal"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

/*
Надо:
- Список пользователей онлайн
- канал для сообщений от пользователей
- горутина для работы с пользователями
*/

func ChatManager(bot *tgbotapi.BotAPI) {

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	var userID int64
	//мониторим чат на входящие сообщения
	for update := range updates {
		//Проверяем на сообщения
		if update.Message != nil {

			userID = update.Message.From.ID
			internal.UserManager(userID)

			userInput := update.Message.Text

			switch userInput {
			case "/start":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Рады вас приветствовать в главном меню бота PTR33!\nВыберите необходимую вам функцию.")
				msg.ReplyMarkup = greeting()
				bot.Send(msg)
			}
		}
		// Проверяем на нажатие кнопок
		if update.CallbackQuery != nil {

			userID = update.CallbackQuery.From.ID

			switch update.CallbackQuery.Data {
			case "lookPart":
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Перехожу в меню просмотра деталей")
				msg.ReplyMarkup = partLookButtons()
				bot.Send(msg)
				partLook(bot, updates)
			case "addPart":
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Перехожу в меню добавления детали.\nВведите название/номер детали.")
				bot.Send(msg)
				//update.CallbackQuery.Data = addPart(bot, updates)
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

		}

	}
}
