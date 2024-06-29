package main

import (
	"fmt"
	"log"
	"ptrbot/build"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println("Get started!!")
	var db, bot = build.Start()

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

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
