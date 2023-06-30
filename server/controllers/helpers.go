package controllers

import (
	"backend-api/exceptions"
	"backend-api/models"
	"fmt"
)

func insertOneMovie(movie models.Movie) {
	i, err := collection.InsertOne(ctxBg, movie)
	if err != nil {
		exceptions.Handle("Error inserting one movie")
	}

	id := i.InsertedID
	fmt.Println("Inserted: ",id)
}

