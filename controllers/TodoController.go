package controller

import (
	"fmt"
	"net/http"

	todoServices "github.com/MingmarGyalzenSherpa44/TodoApp/services"
)

func HandleTodo(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("%v", r.Method)
	switch r.Method {
	case "GET":
		todoServices.GetTodos(w, r)
	case "POST":
		todoServices.CreateTodo(w, r)

	case "DELETE":
		todoServices.DeleteTodo(w, r)
	}

}
