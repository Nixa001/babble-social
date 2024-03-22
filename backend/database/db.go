package database

import (
	q "backend/database/query"
	"backend/utils/seed"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	*sql.DB
}

var DB *Database

func init() {
	db, err := sql.Open("sqlite3", "../backend/database/social_network.db")
	if err != nil {
		log.Println("Error opening database")
		fmt.Println(err)
		os.Exit(1)
	}
	seed.CreateTable(db)
	log.Println("Database opened")
	DB = &Database{db}
}

func (d *Database) Insert(table string, data any) error {
	query, err := q.InsertQuery(table, data)
	fmt.Println(query)
	if err != nil {
		return fmt.Errorf("error creating insert query: %v", err)
	}
	prep, err := d.Prepare(query)
	if err != nil {
		return fmt.Errorf("error preparing insert query: %v", err)
	}
	_, err = prep.Exec()

	return err
}

func (d *Database) Delete(table string, where q.WhereOption) error {

	query := q.DeleteQuery(table, where)
	stmt, err := d.Prepare(query)
	if err != nil {
		return fmt.Errorf("error preparing delete query: %v", err)
	}
	_, err = stmt.Exec()
	return err
}

func (d *Database) Update(table string, object any, where q.WhereOption) error {
	var err error
	query, err := q.UpdateQuery(table, object, where)
	if err != nil {
		return fmt.Errorf("error creating update query: %v", err)
	}
	stmt, err := d.Prepare(query)

	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	return err
}

func (d *Database) GetOneFrom(table string, where q.WhereOption) (*sql.Row, error) {

	query := q.SelectOneFrom(table, where)
	stmt, err := d.Prepare(query)
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRow()
	return row, nil
}

func (d *Database) GetAllFrom(table string, where q.WhereOption, orderby string, limit []int) (*sql.Rows, error) {
	var query string
	if where == nil {
		query = q.SelectAllFrom(table, orderby, limit)
	} else {
		query = q.SelectAllWhere(table, where, orderby, limit)
	}

	stmt, err := d.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("error preparing select query: %v", err)
	}
	rows, err := stmt.Query()

	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (d *Database) GetAllAndJoin(table string, j []q.JoinCondition, where q.WhereOption, orderby string, limit []int) (*sql.Rows, error) {
	var query = q.SelectWithJoinQuery(table, j, where, orderby, limit)

	stmt, err := d.Prepare(query)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query()

	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (d *Database) GetCount(table string, where q.WhereOption) (*sql.Row, error) {
	var query = q.GetCountQuery(table, where)

	stmt, err := d.Prepare(query)
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRow()

	return row, nil
}
