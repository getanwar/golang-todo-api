package todo

import (
	"encoding/json"
	"io"
	"net/http"
)

type Todo struct {
	ID    int
	Title string
	Done  bool
}

var todos = []Todo{
	{ID: 1, Title: "Do homework", Done: false},
	{ID: 2, Title: "Clean dishes", Done: false},
}

func findIndexById(id int) int {
	for index, item := range todos {
		if item.ID == id {
			return index
		}
	}
	return -1
}

func parseBody(r *http.Request) Todo {
	var input Todo
	reqBody, _ := io.ReadAll(r.Body)
	json.Unmarshal(reqBody, &input)

	return input
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(todos)
}

func HandleTodo(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		addTodo(w, r)
	case "PUT":
		updateTodo(w, r)
	case "DELETE":
		deleteTodo(w, r)
	case "GET":
		getTodo(w, r)
	}
}

func addTodo(w http.ResponseWriter, r *http.Request) {
	input := parseBody(r)
	li := todos[len(todos)-1]
	input.ID = li.ID + 1

	todos = append(todos, input)

	json.NewEncoder(w).Encode(todos)
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	input := parseBody(r)
	index := findIndexById(input.ID)

	if index != -1 {
		todos[index] = input
		json.NewEncoder(w).Encode(todos[index])
	}

}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	input := parseBody(r)
	index := findIndexById(input.ID)

	if index != -1 {
		item := todos[index]
		todos = append(todos[:index], todos[index+1:]...)
		json.NewEncoder(w).Encode(item)
	}
}

func getTodo(w http.ResponseWriter, r *http.Request) {
	input := parseBody(r)
	index := findIndexById(input.ID)

	if index != -1 {
		item := todos[index]
		json.NewEncoder(w).Encode(item)
	}

}
