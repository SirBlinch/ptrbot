// Инициализация объектов бота //

// берем токен и создаём юнит бота

package build

func Init() (Token, _BotUnit) {
	var tgToken Token = envToken{}
	var myBot _BotUnit = BotUnit{}
	return tgToken, myBot
}
