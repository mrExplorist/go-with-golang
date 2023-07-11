package router

import (
	"github.com/gorilla/mux"
	controller "github.com/mrExplorist/mongoAPI/controllers"
)

func Router() *mux.Router{



	router := mux.NewRouter()

	router.HandleFunc("/" , controller.ServeHome).Methods("GET")
	router.HandleFunc("/api/movies",controller.GetMyAllMovies).Methods("GET")
	router.HandleFunc("/api/movie",controller.CreateMovie).Methods("POST")

	router.HandleFunc("/api/movie/{id}",controller.MarkAsWatched ).Methods("PUT")
	router.HandleFunc("/api/movie/{id}",controller.DeleteAMovie ).Methods("DELETE")
	router.HandleFunc("/api/deleteallmovies",controller.DeleteAllMovies).Methods("DELETE")



	return router
}