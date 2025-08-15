package router

import (
	"github.com/gorilla/mux"
	"github.com/pran2805/go-lang-learning/controllers"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controllers.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movie/create", controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/update/{id}", controllers.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/delete", controllers.DeleteAll).Methods("DELETE")
	router.HandleFunc("/api/movie/delete/{id}", controllers.DeleteMovieById).Methods("DELETE")

	return router
}
