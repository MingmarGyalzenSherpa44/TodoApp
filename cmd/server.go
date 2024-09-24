package main

import (
	"fmt"
	"log"
	"net/http"

	controller "github.com/MingmarGyalzenSherpa44/TodoApp/controllers"
)

func main() {

	http.HandleFunc("/", controller.RenderHomePage)

	http.HandleFunc("/todos", controller.HandleTodo)

	http.HandleFunc("/delete", controller.DeleteTodo)

	http.HandleFunc("/toggle", controller.ToggleTodo)

	fmt.Printf("Server listening on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
