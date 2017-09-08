package repository

import (
	"projetoPOC/model"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func SavePessoa(pessoa model.Pessoa) {

}
func BuscarTodasPessoa() []model.Pessoa {


	db, err := sql.Open("mysql", "root:Admin@123@/teste_go")
	if err != nil {
		panic(err.Error())

	}

	rows, err := db.Query("SELECT * FROM pessoa")
	if err != nil {
		panic(err.Error())
	}

	listPessoa:=[]model.Pessoa{}
	for rows.Next() {
		var pessoa model.Pessoa
		rows.Scan(&pessoa.Id, &pessoa.Nome, &pessoa.Matricula, &pessoa.DataNascimento)
		listPessoa = append(listPessoa, pessoa)
	}
	/*columns, err := rows.Columns()

	if err != nil {
		panic(err.Error())
	}

	tableData := make([]map[string]interface{}, 0)

	count := len(columns)

	values := make([]interface{}, count)
	scanArgs := make([]interface{}, count)
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err := rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		entry := make(map[string]interface{})
		for i, col := range columns {
			v := values[i]

			b, ok := v.([]byte)
			if (ok) {
				entry[col] = string(b)
			} else {
				entry[col] = v
			}
		}

		tableData = append(tableData, entry)

	}
*/
	defer rows.Close()
	defer db.Close()

	return listPessoa
}

