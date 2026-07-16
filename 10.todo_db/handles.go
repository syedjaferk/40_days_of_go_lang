package main

import (
	"encoding/json"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	response := map[string]string{
		"message": "Welcome to TODO api level 1",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func GetTodosHandler(w http.ResponseWriter, r *http.Request) {

	rows, err := DB.Query(`SELECT id, title, completed FROM todos ORDER by id`)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	defer rows.Close()

	var todos []Todo

	for rows.Next() {
		var todo Todo
		err := rows.Scan(
			&todo.ID, &todo.Title, &todo.Completed,
		)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		todos = append(todos, todo)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var req CreateTodoRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	var todo Todo

	query := `
		INSERT INTO todos(title) VALUES ($1) RETURNING id, title, completed
	`
	err = DB.QueryRow(query, req.Title).Scan(&todo.ID, &todo.Title, &todo.Completed)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

func TodosHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetTodosHandler(w, r)
	case http.MethodPost:
		CreateTodoHandler(w, r)
	}

}
