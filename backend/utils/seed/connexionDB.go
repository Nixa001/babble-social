package seed

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func CreateDB() *sql.DB {
	db, err := sql.Open("sqlite3", "../../database/social_network.db")
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()
	return db

}

var DB = CreateDB()
