// Метод создания юнита бота

package build

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type _BotUnit interface {
	unit(token string) *tgbotapi.BotAPI
	mok()
}

type BotUnit struct{}

func (b BotUnit) unit(token string) *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}
	return bot
}
func (b BotUnit) mok() {
	fmt.Println("Bot unit mok")
}
