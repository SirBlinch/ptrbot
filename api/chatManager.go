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
			continue
		}
		var userInput = update.Message.Text
		svitch(userInput){
			case "/start","/Start": ToDo()
			case "/Новая деталь" :
			case "/Новый инструмент" :
			case "" :
			case "" :
			case "" :
			case "" :
		}	

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, " Здрасте, я работаю!!")
		bot.Send(msg)
	}

}
