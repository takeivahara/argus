package controller

import (
	"net/http"
	"fmt"
)

func Contato(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Accept, X-Requested-With, remember-me")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "contato OK:")
}