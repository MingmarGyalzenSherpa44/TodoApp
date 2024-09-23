package db

import (
	"context"
	"log"

	"github.com/MingmarGyalzenSherpa44/TodoApp/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Database *mongo.Database

func connectDB() *mongo.Database {
	uri := config.Config.DB_URI

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		log.Fatal(err)
	}

	return client.Database("Todo")

}

func init() {
	Database = connectDB()
}
