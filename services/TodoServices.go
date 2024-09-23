package todoServices

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/MingmarGyalzenSherpa44/TodoApp/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var newTodo models.Todo
	err := json.NewDecoder(r.Body).Decode(&newTodo)

	if err != nil {
		http.Error(w, "Invalid Input", http.StatusBadRequest)
		return
	}

	newTodo.CreatedAt = time.Now()
	newTodo.Completed = false
	if result, err := models.Create(&newTodo); err != nil {
		http.Error(w, "Error creating todo", http.StatusInternalServerError)
	} else {
		newTodo.ID = result.InsertedID.(primitive.ObjectID)
	}

	response := struct {
		Message string
		Todo    models.Todo
	}{
		Message: "Todo created successfully!",
		Todo:    newTodo,
	}

	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusCreated)

}

func GetTodos(w http.ResponseWriter, r *http.Request) {

	cursor, err := models.Get()

	if err != nil {
		http.Error(w, "Error getting todos", http.StatusInternalServerError)
	}

	var todos []models.Todo

	if err = cursor.All(context.TODO(), &todos); err != nil {
		log.Fatal(err)
	}

	response := struct {
		Message string
		Todos   []models.Todo
	}{
		Message: "Todos fetched successfully!",
		Todos:   todos,
	}

	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusAccepted)
}
