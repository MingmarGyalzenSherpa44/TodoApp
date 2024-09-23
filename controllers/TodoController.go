package controller

import (
	"net/http"

	todoServices "github.com/MingmarGyalzenSherpa44/TodoApp/services"
)

func HandleTodo(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":

	case "POST":
		todoServices.CreateTodo(w, r)

	}

}
