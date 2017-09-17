package repository

import (
	"argus/model"
	"encoding/json"
)


func SavePessoa(pessoa model.Pessoa) {

}
func BuscarTodasPessoa() model.PessoaBD {


	jsonPessoa := retornarJsonGenerico("SELECT * FROM pessoa")

	//listPessoa:=[]model.Pessoa{}
	var jsontype model.PessoaBD
	json.Unmarshal(jsonPessoa, &jsontype)
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


	return jsontype
}

