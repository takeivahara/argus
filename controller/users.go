package controller

import (
	"net/http"
	"encoding/json"
	"projetoPOC/model"
)

func Users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Accept, X-Requested-With, remember-me")

	user := model.User{"admin", "123"}

	js, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	w.WriteHeader(http.StatusOK)

}