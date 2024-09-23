package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MingmarGyalzenSherpa44/TodoApp/config"
	controller "github.com/MingmarGyalzenSherpa44/TodoApp/controllers"
)

func main() {

	fmt.Println(config.Config.DB_URI)

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Ming")
	})

	http.HandleFunc("/todos", controller.HandleTodo)

	fmt.Printf("Hello world")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
