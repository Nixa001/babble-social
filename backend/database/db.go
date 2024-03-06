package database

import (
	q "backend/database/query"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	*sql.DB
}

// Insert insère des données dans la table spécifiée en utilisant une requête préparée.
// function avec value receveur de type Database
func (d *Database) Insert(table string, data interface{}) error {
	query, values, err := q.InsertQuery(table, data)
	if err != nil {
		fmt.Println(err)
		return err
	}
	prep, err := d.Prepare(query)

	if err != nil {
		fmt.Println(err)
		return err
	}
	defer prep.Close()

	_, err = prep.Exec(values...)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
