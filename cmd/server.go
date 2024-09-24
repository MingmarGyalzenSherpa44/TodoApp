package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"text/template"

	controller "github.com/MingmarGyalzenSherpa44/TodoApp/controllers"
	"github.com/MingmarGyalzenSherpa44/TodoApp/models"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
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

	})

	http.HandleFunc("/todos", controller.HandleTodo)

	fmt.Printf("Hello world")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
