package controllers

import (
	"backend-api/exceptions"
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection
var ctxTODO context.Context = context.TODO()
var ctxBg context.Context = context.Background()

func init() {

	if err := godotenv.Load(); err != nil {
		exceptions.Handle("Error loading .env file")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		exceptions.Handle("MONGO_URI not found")
	}

	client, err := mongo.Connect(ctxTODO, options.Client().ApplyURI(uri))
	if err != nil {
		exceptions.Handle("Error connecting to MongoDB")
	}

	collection = client.Database("fluttergo-fullstack").Collection("movies")

	fmt.Println("Connected to MongoDB")
}
