package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func GetCollection(collection string) *mongo.Collection {
	return MongoClient.Database(os.Getenv("MONGODB_DATABASE")).Collection(collection)
}

func StartMongoDB() error {
	uri := os.Getenv("MONGODB_URI") 

	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	var err error
	MongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	fmt.Println(MongoClient, "[*] Connected to MongoDB!")

	return nil
}

func CloseMongoDB() {
	if err := MongoClient.Disconnect(context.Background()); err != nil {
		panic(err)
	}
	fmt.Println("Disconnected from MongoDB!")
}