package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	MONGO_URI = "mongodb://%s:%s@mongo:%v/golangmongo?authSource=admin"
)

type Item struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var collection *mongo.Collection

func main() {
	r := gin.Default()
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
	collection = db.Collection("items")

	r.GET("/items", getItems)
	r.POST("/items", createItem)

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8800"
	}
	err = r.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}

func getItems(c *gin.Context) {
	var items []Item
	cursor, err := collection.Find(context.Background(), nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var item Item
		if err := cursor.Decode(&item); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		items = append(items, item)
	}

	c.JSON(http.StatusOK, items)
}

func createItem(c *gin.Context) {
	var item Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := collection.InsertOne(context.Background(), item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, item)
}
