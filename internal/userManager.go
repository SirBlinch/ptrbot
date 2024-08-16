package internal

import (
	"database/sql"
)

type user struct {
	id          int64
	state       int
	permissions []string
}

var onlineUsers []user

func usersManager(userID int64) {
	//проверяем список пользователей в сети
	if checkOnline(userID) {
		return
	}
	//если нет в списке то добавляем
	toOnline(userID)
}

func checkOnline(userID int64) bool {
	for _, _user := range onlineUsers {
		if _user.id == userID {
			return true
		}
	}
	return false

}

func toOnline(userID int64) {
	db, err := sql.Open("sqlite3", "D:\\FromFlashCard\\FromLinux\\GO\\PTR_Bot\\DB_1_0.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//проверяем наличие пользователя в базе
	var inDB int
	err = db.QueryRow("SELECT COUNT FROM Users WHERE id = ?", userID).Scan(&inDB)
	if err != nil {
		panic(err)
	}
	//если есть - добавляем в список пользователей в сети
	if inDB != 0 {
		var userPermissions []string
		_permissions, err := db.Query("SELECT * FROM Users_Permissions WHERE userID = ?", userID)
		if err != nil {
			panic(err)
		}
		for _permissions.Next() {
			var permission string
			err = _permissions.Scan(&permission)
			if err != nil {
				panic(err)
			}
			userPermissions = append(userPermissions, permission)
		}
		newUser := user{id: userID, permissions: userPermissions}
		onlineUsers = append(onlineUsers, newUser)

	} else { //Если нет - добавляем пользователя в базу и присваиваем базовый уровень доступа, оповещаем администратора
		addUser(userID)
	}

}

func addUser(userID int64) {
	db, err := sql.Open("sqlite3", "D:\\FromFlashCard\\FromLinux\\GO\\PTR_Bot\\DB_1_0.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO Users (id,) VALUES ($1,)", userID)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("INSERT INTO Users_Permissions (userID, PermissionID) VALUES ($1,0)", userID)
	if err != nil {
		panic(err)
	}
}
