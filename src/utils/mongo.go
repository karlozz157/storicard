package utils

import (
	"context"
	"log"
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	db   *mongo.Database
	once sync.Once
)

func InitMongoDb() *mongo.Database {
	once.Do(func() {
		clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
		client, err := mongo.Connect(context.Background(), clientOptions)
		if err != nil {
			log.Fatal(err)
		}
		secondary := readpref.Secondary()
		dbOpts := options.Database().SetReadPreference(secondary)

		db = client.Database(os.Getenv("MONGO_DATABASE"), dbOpts)
	})

	return db
}
