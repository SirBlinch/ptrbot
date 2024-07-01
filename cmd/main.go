package main

import (
	"fmt"
	"ptrbot/build"

	//"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println("Get started!!")
	var _, bot = build.Start()

	api.chatManager(bot)

	/*
		db, err := sql.Open("sqlite3", "D:\\FromFlashCard\\FromLinux\\GO\\PTR_Bot\\DB_1_0.db")
		if err != nil {
		panic(err)
		}
		defer db.Close()

		// Вставка записи в таблицу

		_, err = db.Exec("INSERT INTO Tools (name, quantity, machine) VALUES ('Фреза D10', '10', 'Фрезер')")
		if err != nil {
			panic(err)
		}
	*/

}

/*
БД:
Поля:
-№ детали (ключ)
-Стоимость
-Время изготовления (возможно лучше разделить на по токарным/фрезерным)
-Инструмент (с разделением по станкам)
-Материал
-№ програмы на станках
*/

//Git TEST
