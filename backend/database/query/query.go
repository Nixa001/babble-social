package query

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
)

type UpdateOption map[string]interface{}

type WhereOption map[string]interface{}
type JoinOnOption struct {
	Tables []string
	On     map[string]interface{}
	Where  WhereOption
}

type JoinCondition struct {
	Table      string
	ForeignKey string
	Reference  string
}

func getMapString(opt UpdateOption) string {
	var res string
	for key, value := range opt {
		if value == 0 || value == nil || value == "" || key == "created_at" || key == "user_id" || key == "birth_date" {
			continue
		} else {
			if res != "" {
				res += ", "
			}
			if v, ok := value.(string); ok {
				res += fmt.Sprintf(`%s="%v"`, key, v)
				continue
			}
			res += fmt.Sprintf("%s=%v", key, value)
		}
	}
	return res
}

func UpdateQuery(table string, object any, where WhereOption) (string, error) {
	toJson, err := json.Marshal(object)
	if err != nil {
		return "", fmt.Errorf("error marshalling object: %v", err)
	}
	toMap := make(map[string]interface{})
	json.Unmarshal(toJson, &toMap)

	toString := getMapString(toMap)
	whToString := getWhereOptionsString(where)
	query := fmt.Sprintf("UPDATE %s SET %v WHERE %s;", table, toString, whToString)

	return query, nil
}

func DeleteQuery(table string, where WhereOption) string {

	whToString := getWhereOptionsString(where)
	query := fmt.Sprintf("DELETE FROM %v WHERE %v;", table, whToString)

	return query
}

func SelectOneFrom(table string, where WhereOption) string {

	whToString := getWhereOptionsString(where)
	query := fmt.Sprintf("SELECT * FROM %v WHERE %v;", table, whToString)

	return query
}

func SelectAllFrom(table string, orderby string, limit []int) string {
	var order string
	if orderby != "" {
		order = fmt.Sprintf("ORDER BY %s", orderby)
	}
	query := fmt.Sprintf("SELECT * FROM %v %s;", table, order)
	if limit != nil {
		query = fmt.Sprintf("SELECT * FROM %v %s LIMIT %v, %v;", table, order, limit[0], limit[1])

	}
	return query
}

func SelectAllWhere(table string, where WhereOption, orderby string, limit []int) string {

	whToString := getWhereOptionsString(where)
	var order string
	if orderby != "" {
		order = fmt.Sprintf("ORDER BY %s", orderby)
	}
	query := fmt.Sprintf("SELECT * FROM %v WHERE %v %s;", table, whToString, order)

	if limit != nil {
		query = fmt.Sprintf("SELECT * FROM %v WHERE %v %s LIMIT %v, %v;", table, whToString, order, limit[0], limit[1])

	}

	return query
}

func InsertQuery(table string, object any) (string, error) {
	fmt.Println("initial => ", object)
	toJson, err := json.Marshal(object)
	if err != nil {
		return "", fmt.Errorf("error marshalling object: %v", err)
	}
	toMap := make(map[string]interface{})
	json.Unmarshal(toJson, &toMap)
	columns, values := getColumnsValues(toMap)

	query := fmt.Sprintf(`INSERT INTO %v (%v) VALUES (%v);`, table, columns, values)

	return query, nil
}

func SelectWithJoinQuery(primaryTable string, joinConditions []JoinCondition, where WhereOption, orderby string, limit []int) string {

	joinClauses := []string{}

	for _, join := range joinConditions {
		joinClause := fmt.Sprintf("LEFT JOIN %s ON %s = %s", join.Table, join.ForeignKey, join.Reference)
		joinClauses = append(joinClauses, joinClause)
	}

	joinClausesString := strings.Join(joinClauses, " ")

	whToString := getWhereOptionsString(where)
	var order string
	if orderby != "" {
		order = fmt.Sprintf("ORDER BY %s", orderby)
	}

	query := fmt.Sprintf("SELECT %v.* FROM %s %s WHERE %s %s;", primaryTable, primaryTable, joinClausesString, whToString, order)
	if limit != nil {
		query = fmt.Sprintf("SELECT %v.* FROM %s %s WHERE %s %s LIMIT %v, %v;", primaryTable, primaryTable, joinClausesString, whToString, order, limit[0], limit[1])
	}
	return query
}

func GetCountQuery(table string, w WhereOption) string {
	whToString := getWhereOptionsString(w)
	query := fmt.Sprintf("SELECT COUNT(*) FROM %v WHERE %v;", table, whToString)
	return query

}

func GetRowIndexQuery(table string, w WhereOption) string {
	whToString := getWhereOptionsString(w)
	query := fmt.Sprintf("SELECT COUNT(*) FROM %v WHERE %v;", table, whToString)
	return query

}

func getWhereOptionsString(w WhereOption) string {
	var res string
	for key, value := range w {
		if res != "" {
			res += "AND "
		}
		res += fmt.Sprintf("(%s%v) ", key, value)
	}
	return res
}

func getColumnsValues(toMap map[string]interface{}) (string, string) {
	var columns, values string
	fmt.Println("to map => ", toMap)
	for k, v := range toMap {
		if v == 0 || v == "" || v == nil {
			continue
		}
		if values != "" {
			values += ", "
		}
		if columns != "" {
			columns += ", "
		}
		columns += strings.ToLower(k)
		if v1, ok := v.(string); ok {
			values += fmt.Sprintf("\"%v\"", v1)
		} else {
			values += fmt.Sprintf("%v", v)
		}
	}
	fmt.Println("columns => ", columns)
	fmt.Println("values => ", values)
	return columns, values
}

func AllTablesQuery() string {
	query := "SELECT name FROM sqlite_schema WHERE type='table' AND name NOT LIKE 'sqlite_%';"
	return query
}

func InsertData(db *sql.DB, query string, values ...interface{}) error {
	// Préparer la requête d'insertion
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Exécuter la requête d'insertion avec les valeurs
	_, err = stmt.Exec(values...)
	if err != nil {
		return err
	}

	return nil
}
