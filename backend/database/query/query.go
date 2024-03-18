package query

import (
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

// UpdateQuery génère une requête préparée pour mettre à jour des lignes dans une table avec des valeurs spécifiées et des conditions WHERE.
func UpdateQuery(table string, object interface{}, where WhereOption) (string, []interface{}, error) {
	toJson, err := json.Marshal(object)
	if err != nil {
		fmt.Println(err)
		return "", nil, err
	}
	toMap := make(map[string]interface{})
	json.Unmarshal(toJson, &toMap)

	setString, setValues := GetMapStringWithPlaceholders(toMap)
	whToString, whereValues := GetWhereOptionsString(where)

	values := append(setValues, whereValues...)
	query := fmt.Sprintf("UPDATE %s SET %s WHERE %s;", table, setString, whToString)

	return query, values, nil
}

// DeleteQuery génère une requête préparée pour supprimer des lignes d'une table avec des conditions WHERE.
func DeleteQuery(table string, where WhereOption) (string, []interface{}) {
	whToString, values := GetWhereOptionsString(where)

	query := fmt.Sprintf("DELETE FROM %v WHERE %v;", table, whToString)

	return query, values
}

// SelectOneFrom génère une requête préparée pour sélectionner une ligne d'une table avec des conditions WHERE.
func SelectOneFrom(table string, where WhereOption) (string, []interface{}) {
	whToString, values := GetWhereOptionsString(where)

	query := fmt.Sprintf("SELECT * FROM %v WHERE %v;", table, whToString)

	return query, values
}

// SelectAllFrom génère une requête préparée pour sélectionner toutes les lignes d'une table avec un tri et une limitation.
func SelectAllFrom(table string, orderby string, limit []int) (string, []interface{}) {
	var order string
	if orderby != "" {
		order = fmt.Sprintf("ORDER BY %s", orderby)
	}

	var query string
	var values []interface{}
	if limit != nil {
		query = fmt.Sprintf("SELECT * FROM %v %s LIMIT ?, ?;", table, order)
		values = append(values, limit[0], limit[1])
	} else {
		query = fmt.Sprintf("SELECT * FROM %v %s;", table, order)
	}

	return query, values
}

// SelectAllWhere génère une requête préparée pour sélectionner toutes les lignes d'une table avec des conditions WHERE, un tri et une limitation.
func SelectAllWhere(table string, where WhereOption, orderby string, limit []int) (string, []interface{}) {
	whToString, values := GetWhereOptionsString(where)

	var order string
	if orderby != "" {
		order = fmt.Sprintf("ORDER BY %s", orderby)
	}

	var query string
	if limit != nil {
		query = fmt.Sprintf("SELECT * FROM %v WHERE %v %s LIMIT ?, ?;", table, whToString, order)
		values = append(values, limit[0], limit[1])
	} else {
		query = fmt.Sprintf("SELECT * FROM %v WHERE %v %s;", table, whToString, order)
	}

	return query, values
}

// SelectWithJoinQuery génère une requête préparée pour sélectionner des données avec jointure, conditions WHERE, tri et limitation.
func SelectWithJoinQuery(primaryTable string, joinConditions []JoinCondition, where WhereOption, orderby string, limit []int) (string, []interface{}) {
	joinClauses := []string{}
	var values []interface{}

	for _, join := range joinConditions {
		joinClause := fmt.Sprintf("LEFT JOIN %s ON %s = %s", join.Table, join.ForeignKey, join.Reference)
		joinClauses = append(joinClauses, joinClause)
	}

	joinClausesString := strings.Join(joinClauses, " ")

	whToString, whereValues := GetWhereOptionsString(where)
	values = append(values, whereValues...)

	var order string
	if orderby != "" {
		order = fmt.Sprintf("ORDER BY %s", orderby)
	}

	var query string
	if limit != nil {
		query = fmt.Sprintf("SELECT %v.* FROM %s %s WHERE %s %s LIMIT ?, ?;", primaryTable, primaryTable, joinClausesString, whToString, order)
		values = append(values, limit[0], limit[1])
	} else {
		query = fmt.Sprintf("SELECT %v.* FROM %s %s WHERE %s %s;", primaryTable, primaryTable, joinClausesString, whToString, order)
	}

	return query, values
}

// GetCountQuery génère une requête préparée pour compter les lignes dans une table qui correspondent aux conditions spécifiées.
func GetCountQuery(table string, w WhereOption) (string, []interface{}) {
	whToString, values := GetWhereOptionsString(w)
	query := fmt.Sprintf("SELECT COUNT(*) FROM %v WHERE %v;", table, whToString)
	return query, values
}

// getWhereOptionsString génère une clause WHERE dans une requête préparée.
func GetWhereOptionsString(w WhereOption) (string, []interface{}) {
	var res string
	var values []interface{}

	for key, value := range w {
		if res != "" {
			res += " AND "
		}
		res += fmt.Sprintf("(%s = ?) ", key)
		values = append(values, value)
	}
	return res, values
}

func GetColumnsValues(data map[string]interface{}) (string, []interface{}) {
	var columns []string
	var values []interface{}

	for column, value := range data {
		columns = append(columns, strings.ToLower(column))
		values = append(values, value)
	}

	// Construit une chaîne de colonnes séparées par des virgules
	columnsStr := strings.Join(columns, ", ")
	return columnsStr, values
}

// getMapStringWithPlaceholders génère une chaîne de mise à jour avec des marqueurs de position pour les valeurs.
func GetMapStringWithPlaceholders(data map[string]interface{}) (string, []interface{}) {
	var setStrings []string
	var values []interface{}

	for key, value := range data {
		setStrings = append(setStrings, fmt.Sprintf("%s = ?", key))
		values = append(values, value)
	}

	setString := strings.Join(setStrings, ", ")
	return setString, values
}

// placeholders génère une chaîne contenant des marqueurs de position pour les paramètres d'une requête préparée.
func Placeholders(count int) string {
	if count < 1 {
		return ""
	}
	// Crée une chaîne de caractères contenant count marqueurs de position (?)
	return strings.Repeat("?, ", count-1) + "?"
}

// InsertQuery prépare une requête d'insertion avec des marqueurs de position.
func InsertQuery(table string, object interface{}) (string, []interface{}, error) {
	toJson, err := json.Marshal(object)
	if err != nil {
		fmt.Println(err)
		return "", nil, err
	}
	toMap := make(map[string]interface{})
	json.Unmarshal(toJson, &toMap)
	columns, values := GetColumnsValues(toMap)
	fmt.Println(columns)
	fmt.Println(values)
	query := fmt.Sprintf(`INSERT INTO %v (%v) VALUES (%v);`, table, columns, Placeholders(len(values)))
	return query, values, nil
}
