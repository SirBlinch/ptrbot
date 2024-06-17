// Метод создания юнита бота

package build

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type BotUnit struct{}

func (b BotUnit) botUnit(token string) *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}
	return bot
}
