package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/MingmarGyalzenSherpa44/TodoApp/models"
	todoServices "github.com/MingmarGyalzenSherpa44/TodoApp/services"
)

func HandleTodo(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("%v", r.Method)
	switch r.Method {
	case "GET":
		todoServices.GetTodos(w, r)

	case "POST":
		todoServices.CreateTodo(w, r)

	}

}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	todoServices.UpdateTodo(w, r)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {

	todoServices.DeleteTodo(w, r)

}

func RenderHomePage(w http.ResponseWriter, r *http.Request) {
	cursor, err := models.Get()
	if err != nil {
		log.Fatal(err)
	}

	var todos []models.Todo
	if err = cursor.All(context.TODO(), &todos); err != nil {
		log.Fatal(err)
	}

	tmpl := template.Must(template.ParseFiles("template/index.html"))

	tmpl.Execute(w, struct {
		Todos []models.Todo
	}{
		Todos: todos,
	})

}
