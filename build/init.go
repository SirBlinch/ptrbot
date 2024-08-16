// Инициализация объектов бота //

// берем токен и создаём юнит бота

package build

func Init() (Token, _BotUnit) {
	var tgToken Token = envToken{}
	var myBot _BotUnit = BotUnit{}
	return tgToken, myBot
}

/*
Инициализируем:
Token
bot construktor
data base
chat manager		// Отправляет и получает сообщения в/из чата
db manager			// реализует запись и чтение из/в базу данных при чтении результат передаёт на вывод в chat manager
*/
