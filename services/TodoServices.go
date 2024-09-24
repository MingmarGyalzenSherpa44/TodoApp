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
	err := r.ParseForm()

	if err != nil {
		http.Error(w, "Invalid Input", http.StatusBadRequest)
		return
	}

	newTodo.Title = r.FormValue("title")
	newTodo.CreatedAt = time.Now()
	newTodo.Completed = false
	if result, err := models.Create(&newTodo); err != nil {
		http.Error(w, "Error creating todo", http.StatusInternalServerError)
	} else {
		newTodo.ID = result.InsertedID.(primitive.ObjectID)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)

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

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Fatal(err)
	}

	if _, err := models.Delete(objectId); err != nil {

		log.Fatal(err)

		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func ToggleTodo(w http.ResponseWriter, r *http.Request) {

	reqStruct := struct {
		Completed bool
	}{}
	json.NewDecoder(r.Body).Decode(&reqStruct)

	id := r.URL.Query().Get("id")

	if id == "" {
		if _, err := models.UpdateAll(true); err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(struct {
			Message   string
			Completed bool
		}{
			Message: "Toggle successful",
		})

		return
	}

	response := struct {
		Message string
	}{
		Message: "",
	}

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		response.Message = "Not a valid id!"
		json.NewEncoder(w).Encode(response)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err := models.Update(objectId, reqStruct.Completed); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(struct {
		Message   string
		Completed bool
	}{
		Message: "Toggle successful",
	})

}
