package services

import (
	"auth-sys/database"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	tableName struct{} `pg:"users"` // https://pg.uptrace.dev/models/
	Id        int
	Email     string
	Password  string
}

func LoginUser(email, password string) (bool, int) {

	// get user by his email
	var foundedUser []User
	err := database.DB().Model(&foundedUser).Where("email = ?", email).Limit(1).Select()
	if err != nil {
		log.Fatal("ERROR: Can't get user by his email " + err.Error())
		return false, 0
	}

	if len(foundedUser) == 1 {
		err := bcrypt.CompareHashAndPassword([]byte(foundedUser[0].Password), []byte(password))

		if err != nil {
			log.Println("Password is incorrect")
			return false, 0
		}

		fmt.Println("User is found !")

		return true, foundedUser[0].Id

	} else {
		log.Println("Email and Password are incorrect")
		return false, 0
	}
}
