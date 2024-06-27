// Запускаем бота
package build

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func Start() *sql.DB {
	token, myBot, ptrDB := Init()
	fmt.Println("prt_token is " + token.token())
	myBot.mok()
	//myBot.unit(token.token())
	var db = ptrDB.connectToDBb()
	fmt.Println("Я стартанул!")
	return db
}
