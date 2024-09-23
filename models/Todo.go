package models

import (
	"context"
	"time"

	"github.com/MingmarGyalzenSherpa44/TodoApp/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Todo struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Title     string             `bson:"title"`
	Completed bool               `bson:"completed"`
	CreatedAt time.Time          `bson:"created_at"`
}

var todoCollection *mongo.Collection

func init() {
	todoCollection = db.Database.Collection("todos")
}

func Create(t *Todo) (*mongo.InsertOneResult, error) {

	result, err := db.Database.Collection("todos").InsertOne(context.TODO(), t)

	return result, err
}

func Get() (*mongo.Cursor, error) {

	result, err := todoCollection.Find(context.TODO(), bson.M{})

	return result, err
}
