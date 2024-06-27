// Инициализация объектов бота //

// берем токен и создаём юнит бота

package build

func Init() (Token, _BotUnit, _DbConnector) {
	var tgToken Token = envToken{}
	var myBot _BotUnit = BotUnit{}
	var ptrDB _DbConnector = sqlite3DB{}
	return tgToken, myBot, ptrDB
}
