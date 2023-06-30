package router

import (
	"backend-api/controllers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/movies", controllers.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movies", controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movies/{id}", controllers.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/api/movies", controllers.DeleteAllMovies).Methods("DELETE")
	router.HandleFunc("/api/movies/{id}", controllers.SetMovieWatched).Methods("PUT")

	return router
}
