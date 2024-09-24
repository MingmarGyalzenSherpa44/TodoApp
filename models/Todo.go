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

	result, err := todoCollection.InsertOne(context.TODO(), t)

	return result, err
}

func Get() (*mongo.Cursor, error) {

	result, err := todoCollection.Find(context.TODO(), bson.M{})

	return result, err
}

func Delete(_id primitive.ObjectID) (*mongo.DeleteResult, error) {

	result, err := todoCollection.DeleteOne(context.TODO(), bson.M{"_id": _id})
	return result, err
}

func Update(id primitive.ObjectID, completed bool) (*mongo.UpdateResult, error) {

	result, err := todoCollection.UpdateByID(context.TODO(), id, bson.M{"$set": bson.M{"completed": completed}})
	return result, err
}

func UpdateAll(completed bool) (*mongo.UpdateResult, error) {
	result, err := todoCollection.UpdateMany(context.TODO(), bson.D{}, bson.M{"$set": bson.M{"completed": completed}})
	return result, err
}
