package database

import (
	"github.com/go-pg/pg/v10"
)

func DB() *pg.DB {
	// connStr := "postgres://rootpg:123456@localhost:5432/authsys"
	// use opt as an option in the connect fn
	// opt, err := pg.ParseURL(connStr)
	// fmt.Println(opt)
	// if err != nil {
	// 	log.Fatal("Parse DB URL Error: " + err.Error())
	// }

	db := pg.Connect(&pg.Options{
		User:     "rootpg",
		Database: "authsys",
		Password: "123456",
		// both of this works, but when i run go test, i get an error that goappdb host doesn't exist.
		// Addr:     "goappdb:5432", // "goappdb" is the postgres database container name
		Addr: "192.168.1.12:5432", // my local IPv4 address, type "ipconfig" if you are in windows and you shall see it at first.
	})

	return db
}
