package repository

import (
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
	"database/sql"
)

func abrirConexaoBancoDados() *sql.DB {

	db, err := sql.Open("mysql", "root:Admin@123@/teste_go")
	if err != nil {
		panic(err.Error())

	}
	return db
}

func retornarJsonGenerico(query string) []byte{

	db := abrirConexaoBancoDados()
	rows, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}

	names, err := rows.Columns()
	if err != nil {
		panic(err)
	}
	// vals stores the values from one row
	vals := make([]interface{}, 0, len(names))
	for _, _ = range names {
		vals = append(vals, new(string))
	}
	// rowMaps stores all rows
	rowMaps := make([]map[string]string, 0)
	for rows.Next() {
		rows.Scan(vals...)
		// now make value list into name=>value map
		currRow := make(map[string]string)
		for i, name := range names {
			currRow[name] = *(vals[i].(*string))
		}
		// accumulating rowMaps is the easy way out
		rowMaps = append(rowMaps, currRow)
	}
	json, err := json.Marshal(rowMaps)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	defer db.Close()

	return json
}



