package build

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectToDBb() *sql.DB {
	db, err := sql.Open("sqlite3", "D:\\FromFlashCard\\FromLinux\\GO\\PTR_Bot\\DB_1_0.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	return db

}
