package main

import (
	"encoding/json"
	"net/http"
)

// HomeHandler handles the root endpoint.
func HomeHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	response := map[string]string{
		"message": "Welcome to Todo API",
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}

// TodosHandler dispatches requests based on HTTP method.
func TodosHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		GetTodosHandler(w, r)

	case http.MethodPost:
		CreateTodoHandler(w, r)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

// GET /todos
func GetTodosHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(todos)
}

// POST /todos
func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {

	var req CreateTodoRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.Title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	todo := Todo{
		ID:        nextID,
		Title:     req.Title,
		Completed: false,
	}

	nextID++

	todos = append(todos, todo)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(todo)
}
