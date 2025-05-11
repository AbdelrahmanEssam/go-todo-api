package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json: "title"`
	Done  bool   `json:"done"`
}

var todos = []Todo{}

var idCounter = 1

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Todo API is running")
	})

	http.HandleFunc("/todos", todosHandler)

	fmt.Println("Server is listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func todosHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		getTodos(w, r)
	case http.MethodPost:
		createTodo(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	var t Todo
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	t.ID = idCounter
	idCounter++
	todos = append(todos, t)

	w.Header().Set("Cintent-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(t)
}
