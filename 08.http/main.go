package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World Vignesh!")
}

func handler(w http.ResponseWriter, r *http.Request) {

	var user User

	switch r.Method {
	case http.MethodGet:
		fmt.Fprintln(w, "Get Request Received")

	case http.MethodPost:
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		fmt.Fprintln(w, "Post Request Received", user.Name, user.Age)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println("Error : ", err)
	}
}
