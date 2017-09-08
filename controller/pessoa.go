package controller

import (
	"net/http"
	"projetoPOC/repository"
	"fmt"
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

	pessoa := repository.BuscarTodasPessoa()

	fmt.Fprint(w, "Pessoa de Matricula: ")
	fmt.Fprintln(w, pessoa[0].Matricula)

	fmt.Fprint(w, "Pessoa de Nome: ")
	fmt.Fprintln(w, pessoa[0].Nome)

	fmt.Fprint(w, "Pessoa de Nascimento: ")
	fmt.Fprintln(w, pessoa[0].DataNascimento)

	fmt.Fprint(w, "Pessoa de Matricula: ")
	fmt.Fprintln(w, pessoa[1].Matricula)

	fmt.Fprint(w, "Pessoa de Nome: ")
	fmt.Fprintln(w, pessoa[1].Nome)

	fmt.Fprint(w, "Pessoa de Nascimento: ")
	fmt.Fprintln(w, pessoa[1].DataNascimento.Format("dd/MM/yyyy"))
}