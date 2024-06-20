// Инициализация объектов бота //

// берем токен и создаём юнит бота

package build

func Init() (Token, BotUnit) {
	var tgToken Token = envToken{}
	var myBot Unit = BotUnit{}
	return tgToken, myBot
}
