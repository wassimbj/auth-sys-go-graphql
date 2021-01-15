package database

import (
	"github.com/go-pg/pg/v10"
)

func DB() *pg.DB {
	// connStr := "postgres://rootpg:123456@localhost:5432/mediumclone"
	// use opt as an option in the connect fn
	// opt, err := pg.ParseURL(connStr)
	// fmt.Println(opt)
	// if err != nil {
	// 	log.Fatal("Parse DB URL Error: " + err.Error())
	// }

	db := pg.Connect(&pg.Options{
		User:     "rootpg",
		Database: "mediumclone",
		Password: "123456",
		Addr:     "goappdb:5432", // "goappdb" is the postgres database container name
		// Network:  "tcp", // default is tcp
	})

	return db
}
