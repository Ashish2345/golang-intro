package routes

import (
	"github.com/Ashish2345/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStore = func(router *mux.Router) {
	// router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books/", controllers.GetBook).Methods("GET")
	// router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	// router.HandleFunc("/book/update/{bookId}", controllers.UpdateById).Methods("PUT")
	// router.HandleFunc("/book/delete/{bookId}", controllers.DelteById).Methods("DELETE")

}
