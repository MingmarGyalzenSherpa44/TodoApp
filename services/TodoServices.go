package todoServices

import (
	"encoding/json"
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
