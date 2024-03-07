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

func (d *Database) Update(table string, object interface{}, where q.WhereOption) error {
	query, values, err := q.UpdateQuery(table, object, where)
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

func (d *Database) Delete(table string, where q.WhereOption) error {
	query, values := q.DeleteQuery(table, where)
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

func (d *Database) GetOneForm(table string, where q.WhereOption) (*sql.Row, error) {
	query, values := q.SelectOneFrom(table, where)
	prep, err := d.Prepare(query)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer prep.Close()

	row := prep.QueryRow(values...)
	return row, nil
}

func (d *Database) GetAllFrom(table string, where q.WhereOption, orderby string, limit []int) (*sql.Rows, error) {
	var query string
	var values []interface{}
	if where == nil {
		query, values = q.SelectAllFrom(table, orderby, limit)
	} else {
		query, values = q.SelectAllWhere(table, where, orderby, limit)
	}

	prep, err := d.Prepare(query)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer prep.Close()

	rows, err := prep.Query(values...)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return rows, nil
}

func (d *Database) GetAllAndJoin(table string, j []q.JoinCondition, where q.WhereOption, orderby string, limit []int) (*sql.Rows, error) {
	query, values := q.SelectWithJoinQuery(table, j, where, orderby, limit)

	prep, err := d.Prepare(query)
	if err != nil {
		return nil, err
	}
	rows, err := prep.Query(values...)

	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (d *Database) GetCount(table string, where q.WhereOption) (*sql.Row, error) {
	query, values := q.GetCountQuery(table, where)

	prep, err := d.Prepare(query)
	if err != nil {
		return nil, err
	}

	row := prep.QueryRow(values...)
	return row, nil

}
