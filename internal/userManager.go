package internal

import (
	"database/sql"
)

type user struct {
	id          int
	state       int
	permissions string
}

var onlineUsers []user

func addUser(userID int) {
	db, err := sql.Open("sqlite3", "D:\\FromFlashCard\\FromLinux\\GO\\PTR_Bot\\DB_1_0.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO Users (id,) VAIUES ($1,)", userID)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("INSERT INTO Users_Permissions (userID, PermissionID) VALUES ($1,0)", userID)
	if err != nil {
		panic(err)
	}
}

func toOnline(userID int) {
	db, err := sql.Open("sqlite3", "D:\\FromFlashCard\\FromLinux\\GO\\PTR_Bot\\DB_1_0.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	var value int
	err = db.QueryRow("SELECT COUNT(*) FROM Users WHERE id = ?", userID).Scan(&value)
	if err != nil {
		panic(err)
	}
	if value == 0 {
		addUser(userID)
	} else {

	}
}

func toOffline() {}
