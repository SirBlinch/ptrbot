package build

import (
	"database/sql"
)

type _DbBuilder interface {
	buildDB()
}

type DbBuilder struct{}

func (d DbBuilder) buildDB() {
	db, err := sql.Open("sqlite3", "PTR_DB")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
