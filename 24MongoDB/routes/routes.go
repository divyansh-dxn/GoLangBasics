package routes

import (
	"gomongo/controllers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router{
	router := mux.NewRouter()
	router.HandleFunc("/api",controllers.ServeHome).Methods("GET")
	router.HandleFunc("/api/movie",controllers.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movie",controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}",controllers.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}",controllers.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/api/movie/delete/all",controllers.DeleteAllMovies).Methods("DELETE")
	return router
}