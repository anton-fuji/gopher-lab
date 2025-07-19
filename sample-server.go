package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Todo struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
}

var todos []Todo
var nextID = 1

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/todos", todosHnadler)
	http.HandleFunc("/todos/", todoByIdHnadler)

	fmt.Println("Server running at: 8000")
	http.ListenAndServe(":8000", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"message": "Hello, World!"})
}

func todosHnadler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		json.NewEncoder(w).Encode(todos)
		return
	}
	if r.Method == "POST" {
		var t Todo
		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			http.Error(w, "Invalid", http.StatusBadRequest)
			return
		}

		t.ID = nextID
		nextID++
		todos = append(todos, t)
		json.NewEncoder(w).Encode(t)
		return
	}
	http.Error(w, "Method Not Allow", http.StatusMethodNotAllowed)
}

func todoByIdHnadler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/todos/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	for _, t := range todos {
		if t.ID == id {
			json.NewEncoder(w).Encode(t)
			return
		}
	}
	http.Error(w, "Not Found", http.StatusBadRequest)
}
