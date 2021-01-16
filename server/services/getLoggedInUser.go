package services

import (
	"auth-sys/database"
	"log"
)

type LoggedInUser struct {
	tableName struct{} `pg:"users,select:users"` // https://pg.uptrace.dev/models/
	Id        int
	Fullname  string
	Email     string
	Createdat string
}

func GetLoggedInUser(userId interface{}) LoggedInUser {
	// get user by his email
	var user LoggedInUser
	err := database.DB().Model(&user).Where("id = ?", userId).Limit(1).Select()
	if err != nil {
		log.Fatal("ERROR: Can't get user by his id" + err.Error())
		return LoggedInUser{}
	}
	return user

}
