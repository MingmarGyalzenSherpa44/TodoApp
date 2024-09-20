package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MingmarGyalzenSherpa44/TodoApp/config"
)

func main() {

	fmt.Println(config.Config.DB_URI)

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Ming")
	})
	fmt.Printf("Hello world")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
