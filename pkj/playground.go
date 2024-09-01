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
//////////////////////////////////////////////////////////////////////////////////////////
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
			internal.UserManager(userID)
//////////////////////////////////////////////////////////////////////////////////////////
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


/*
package pkj
import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("1.com", "http://1.com"),
		tgbotapi.NewInlineKeyboardButtonData("2", "2"),
		tgbotapi.NewInlineKeyboardButtonData("3", "3"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("4", "4"),
		tgbotapi.NewInlineKeyboardButtonData("5", "5"),
		tgbotapi.NewInlineKeyboardButtonData("6", "6"),
	),
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	// Loop through each update.
	for update := range updates {
		// Check if we've gotten a message.
		if update.Message != nil {
			// Construct a new message from the given chat ID and containing
			// the text that we received.
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

			// If the message was open, add a copy of our numeric keyboard.
			switch update.Message.Text {
			case "open":
				msg.ReplyMarkup = numericKeyboard
			}

			// Send the message.
			if _, err = bot.Send(msg); err != nil {
				panic(err)
			}
		} else if update.CallbackQuery != nil {
			// Respond to the callback query, telling Telegram to show the user
			// a message with the data received.
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := bot.Request(callback); err != nil {
				panic(err)
			}

			// And finally, send a message containing the data received.
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
		}
	}
}







package main

import (
	"log"
	"os"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("1.com", "http://1.com"),
		tgbotapi.NewInlineKeyboardButtonData("2", "2"),
		tgbotapi.NewInlineKeyboardButtonData("3", "3"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("4", "4"),
		tgbotapi.NewInlineKeyboardButtonData("5", "5"),
		tgbotapi.NewInlineKeyboardButtonData("6", "6"),
	),
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	// Loop through each update.
	for update := range updates {
		// Check if we've gotten a message update.
		if update.Message != nil {
			// Construct a new message from the given chat ID and containing
			// the text that we received.
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

			// If the message was open, add a copy of our numeric keyboard.
			switch update.Message.Text {
			case "open":
				msg.ReplyMarkup = numericKeyboard
			}

			// Send the message.
			if _, err = bot.Send(msg); err != nil {
				panic(err)
			}
		} else if update.CallbackQuery != nil {
			// Respond to the callback query, telling Telegram to show the user
			// a message with the data received.
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)

			if _, err := bot.Request(callback); err != nil {
				panic(err)
			}

			// And finally, send a message containing the data received.
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)

			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
		}
	}
}