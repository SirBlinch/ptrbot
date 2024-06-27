package build

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type _DbConnector interface {
	connectToDBb() *sql.DB
}

type sqlite3DB struct{}

func (s sqlite3DB) connectToDBb() *sql.DB {
	db, err := sql.Open("sqlite3", "D:\\FromFlashCard\\FromLinux\\GO\\PTR_Bot\\DB_1_0.db")
	if err != nil {
		panic(err)
	}
	fmt.Println("Connect to DB		OK")
	defer db.Close()
	return db

}
