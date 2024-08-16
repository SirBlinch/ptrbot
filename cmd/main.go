package main

import (
	"fmt"
	"ptrbot/api"
	"ptrbot/build"

	//"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println("Get started!!")
	var bot = build.Start()

	api.ChatManager(bot)

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

/*
	Добавить деталь{
		Название (уникальное)
		while заполнение != закончить{
			svitch (что добавить)
			добавить цену
			добавить время
			добавить материал
			добавить инструмент{
				svitch (выбрать инструмент)
				инструмент 1
				инструмент 2
				инструмент 3
				.
				.
				инструмент n
				добавить инструмент (меню добавления инструмента)

			}
		}
	}



	Добавить инструмент{
		Название (уникальное)
		while заполнение != закончить{
			svitch (что указать)
			указать количество
			указать под какой станок/выбрать из вариантов
		}
	}


*/
