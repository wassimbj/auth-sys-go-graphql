package services

import (
	"auth-sys/database"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type NewUser struct {
	tableName struct{} `pg:"users"` // https://pg.uptrace.dev/models/
	Fullname  string
	Email     string
	Password  string
}

func CreateAccount(fullname, email, password string) bool {

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.Fatal("Can't create account, hash password went wrong " + err.Error())
		return false
	}
	res := &NewUser{struct{}{}, fullname, email, string(hashedPass)}
	_, insertErr := database.DB().Model(res).Insert()
	if insertErr != nil {
		log.Fatal("Can't create account, something went wrong " + insertErr.Error())
		return false
	}

	return true
}
