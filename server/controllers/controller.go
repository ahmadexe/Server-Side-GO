package controllers

import (
	"backend-api/exceptions"
	"backend-api/models"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func insertOneMovie(movie models.Movie) {
	fmt.Println(movie)
	i, err := collection.InsertOne(ctxBg, movie)
	if err != nil {
		exceptions.Handle("Error inserting one movie")
	}

	id := i.InsertedID
	fmt.Println("Inserted: ", id)
}

func setMovieWatched(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	r, err := collection.UpdateOne(ctxBg, filter, update)
	if err != nil {
		exceptions.Handle("Error updating movie")
	}

	fmt.Println("Updated: ", r.ModifiedCount)
}

func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}

	r, err := collection.DeleteOne(ctxBg, filter)
	if err != nil {
		exceptions.Handle("Error deleting movie")
	}

	fmt.Println("Deleted: ", r.DeletedCount)
}

func deleteAllMovies() {
	r, err := collection.DeleteMany(ctxBg, bson.M{})
	if err != nil {
		exceptions.Handle("Error deleting movies")
	}

	fmt.Println("Deleted: ", r.DeletedCount)
}

func getAllMovies() []models.Movie {
	movies := []models.Movie{}
	cur, err := collection.Find(ctxBg, bson.M{})
	if err != nil {
		exceptions.Handle("Error getting all movies")
	}

	for cur.Next(ctxBg) {
		var movie models.Movie
		err := cur.Decode(&movie)
		if err != nil {
			exceptions.Handle("Error decoding movie")
		}

		movies = append(movies, movie)
	}
	defer cur.Close(ctxBg)

	return movies
}

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	movies := getAllMovies()
	enc := json.NewEncoder(w)
	enc.Encode(movies)
	fmt.Println("Movies: ", movies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "POST")
	var movie models.Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	fmt.Println("Movie: ", movie)
	insertOneMovie(movie)
	enc := json.NewEncoder(w)
	enc.Encode(movie)
	fmt.Println("Movie: ", movie)
}

func SetMovieWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "PUT")
	params := mux.Vars(r)
	var id string = params["id"]
	setMovieWatched(id)
	fmt.Println("Movie watched: ", id)
}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "DELETE")
	params := mux.Vars(r)
	var id string = params["id"]
	deleteOneMovie(id)
	fmt.Println("Movie deleted: ", id)
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "DELETE")
	deleteAllMovies()
	fmt.Println("All movies deleted")
}