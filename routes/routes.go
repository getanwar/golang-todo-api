package routes

import (
	"fmt"
	"net/http"
	"todoApp/modules/todo"
)

func Init() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Server is working!")
	})
	http.HandleFunc("/todos", todo.GetTodos)
	http.HandleFunc("/todo", todo.HandleTodo)
}
