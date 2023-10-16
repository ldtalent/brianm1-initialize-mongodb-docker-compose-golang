package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	MONGO_URI = "mongodb://%s:%s@mongo:%v/golangmongo?authSource=admin"
)

func main() {
	conn := fmt.Sprintf(MONGO_URI, "root", "rootpassword", 8081)

	// Access the client
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(conn))
	// Capture the errors
	if err != nil {
		log.Fatalf("Could not connect to MongoDB: %v", err)
	}

	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			log.Fatalf("could not connect to MongoDB: %v", err)
		}
	}()

	db := client.Database("golangmongo")
}
