package main

import (
	"log"
	"net/http"
	"tsis1_golang/pkg/handlers"
)

func main() {
	router := handlers.NewRouter()

	log.Fatal(http.ListenAndServe(":8080",router))
}