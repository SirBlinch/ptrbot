// Инициализация объектов бота //

// берем токен и создаём юнит бота

package build

func Init() (Token, BotUnit) {
	var tgToken Token = flagToken{}
	var myBot BotUnit
	return tgToken, myBot
}
