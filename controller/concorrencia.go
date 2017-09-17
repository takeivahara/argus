package controller

import (
	"net/http"
	"fmt"


	"strconv"
	"time"
)

func Concorrencia(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Accept, X-Requested-With, remember-me")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "contato:OK")
	escreverNumeroEmMinuto(w)

}

func escreverNumeroEmMinuto(w http.ResponseWriter) {
	//timer := time.NewTimer(time.Second * time.Duration(tempoEmSegundos))
	var conteudo []string

	start := time.Now()
	for i := 0; i < 999999999; i++ {

		conteudo = append(conteudo, strconv.Itoa(i))
		duration := time.Now().Sub(start)
		if duration > time.Second*10 {
			break
		}



	}

	fmt.Fprint(w, strconv.Itoa(len(conteudo)))
	//arquivo, err := os.Create("caminho.txt")

	//if err != nil {
	//	panic(err)
	//}

	//defer arquivo.Close()

	// Cria um escritor responsavel por escrever cada linha do slice no arquivo de texto
	//escritor := bufio.NewWriter(arquivo)

	//fmt.Fprintln(escritor, len(conteudo))

	//escritor.Flush()
	//<- timer.C
}

