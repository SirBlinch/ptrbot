// Запускаем бота
package build

import (
	"database/sql"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/mattn/go-sqlite3"
)

func Start() (*sql.DB, *tgbotapi.BotAPI) {
	token, myBot, ptrDB := Init()
	fmt.Println("prt_token is " + token.token())
	//myBot.mok()
	var bot = myBot.unit(token.token())
	bot.Debug = true
	var db = ptrDB.connectToDBb()
	fmt.Println("Я стартанул!")
	return db, bot
}
