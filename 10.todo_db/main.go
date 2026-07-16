package main

import (
	"log"
	"net/http"
)

func main() {

	ConnectDB()

	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/todos", TodosHandler)

	log.Println("TODO API Server. Listening on 8081")

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err)
	}

}
