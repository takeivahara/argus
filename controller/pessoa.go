package controller

import (
	"net/http"
	"argus/repository"
	"fmt"
	"argus/model"
	"strconv"
	"time"
)

func SavePessoa(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Accept, X-Requested-With, remember-me")
	w.WriteHeader(http.StatusOK)

}


func BuscarTodasPessoa(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Accept, X-Requested-With, remember-me")
	w.WriteHeader(http.StatusOK)

	pessoaBD:= repository.BuscarTodasPessoa()
	var listPessoa []model.Pessoa
	for i := 0; i < len(pessoaBD); i++ {

		var pessoa model.Pessoa
		//fmt.Fprint(w, "Pessoa de Matricula: ")
		retMatricula,_ := strconv.Atoi(pessoaBD[i].MATRICULA)
		pessoa.Matricula = retMatricula
		//fmt.Fprintln(w, pessoa.Matricula)

		//fmt.Fprint(w, "Pessoa de Nome: ")
		pessoa.Nome = pessoaBD[i].NOME
		//fmt.Fprintln(w, pessoa.Nome)

		//fmt.Fprint(w, "Pessoa de Nascimento: ")
		layout := "2006-01-02 15:04:05"
		retDataNascimento,_ := time.Parse(layout, pessoaBD[i].DATANASCIMENTO)
		pessoa.DataNascimento = retDataNascimento
		//fmt.Fprintln(w, pessoa.DataNascimento.Format("dd/MM/yyyy"))

		listPessoa = append(listPessoa, pessoa)
	}

	for i := 0; i < len(listPessoa); i++ {
		fmt.Fprint(w, "Pessoa de Matricula: ")
		fmt.Fprintln(w, listPessoa[i].Matricula)
		fmt.Fprint(w, "Pessoa de Nome: ")
		fmt.Fprintln(w, listPessoa[i].Nome)
		fmt.Fprint(w, "Pessoa de Nascimento: ")
		fmt.Fprintln(w, listPessoa[i].DataNascimento.Format("dd/MM/yyyy"))
	}
}