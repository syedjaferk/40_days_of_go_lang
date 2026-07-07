package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/todos", TodosHandler)

	log.Println("Todo API Server Started")
	log.Println("Listening on :8081")

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err)
	}
}
