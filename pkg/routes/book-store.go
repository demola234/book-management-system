package routes

import (
	"github.com/demola234/go-book-management/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRouter = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("GET")
	router.HandleFunc("/book/", controllers.GetBooks).Methods("POST")
	router.HandleFunc("/book/{id}", controllers.GetBook).Methods("GET")
	router.HandleFunc("book/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("book/{id}", controllers.DeleteBook).Methods("DELETE")

}
