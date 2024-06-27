package main

import (
	"database/sql"
	"fmt"
	"ptrbot/build"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println("Get started!!")
	var db = build.Start()

	db, err := sql.Open("sqlite3", "DB_1_0.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	record := map[string]interface{}{
		"name":     "Инструмент 1",
		"quantity": 10,
		"machine":  "Станок 1",
	}

	// Вставка записи в таблицу
	_, err = db.Exec("INSERT INTO Tools (name, quantity, machine) VALUES ($1, $2, $3)", record)
	if err != nil {
		panic(err)
	}

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
