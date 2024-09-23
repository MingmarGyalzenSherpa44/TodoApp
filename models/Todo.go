package models

import (
	"context"
	"log"
	"time"

	"github.com/MingmarGyalzenSherpa44/TodoApp/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	title     string             `bson:"title"`
	completed bool               `bson:"completed"`
	createdAt time.Time          `bson:"created_at"`
}

func Create(t *Todo) {

	_, err := db.Database.Collection("todos").InsertOne(context.TODO(), t)

	if err != nil {
		log.Fatal(err)
	}

}
