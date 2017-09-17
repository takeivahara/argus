package controller

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
)

func HelloName(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	name := vars["name"]

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hello:", name)
}