// Запускаем бота
package build

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Start() *tgbotapi.BotAPI {
	token, myBot := Init()
	fmt.Println("prt_token is " + token.token())
	//myBot.mok()
	var bot = myBot.unit(token.token())
	bot.Debug = true
	fmt.Println("Я стартанул!")
	return bot
}
