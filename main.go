package main

import (
	"net/http"
	"log"
	"argus/route"
)

func main() {

	router := route.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}

