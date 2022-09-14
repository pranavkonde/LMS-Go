package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pranavkonde/LMS-Go/api"
	"github.com/pranavkonde/LMS-Go/book"
	"github.com/pranavkonde/LMS-Go/transaction"
	"github.com/pranavkonde/LMS-Go/user"
)

const (
// versionHeader = "Accept"
)

func initRouter(dep dependencies) (router *mux.Router) {
	// v1 := fmt.Sprintf("application/vnd.%s.v1", config.AppName())

	router = mux.NewRouter()
	router.HandleFunc("/ping", pingHandler).Methods(http.MethodGet)
	router.HandleFunc("/users", user.Create(dep.UserService)).Methods(http.MethodPost)
	router.HandleFunc("/users", user.List(dep.UserService)).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", user.FindByID(dep.UserService)).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", user.DeleteByID(dep.UserService)).Methods(http.MethodDelete)
	router.HandleFunc("/users", user.Update(dep.UserService)).Methods(http.MethodPut)

	router.HandleFunc("/books", book.Create(dep.BookService)).Methods(http.MethodPost)
	router.HandleFunc("/books", book.List(dep.BookService)).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", book.FindByID(dep.BookService)).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", book.DeleteByID(dep.BookService)).Methods(http.MethodDelete)
	router.HandleFunc("/books", book.Update(dep.BookService)).Methods(http.MethodPut)

	router.HandleFunc("/transactions", transaction.Create(dep.TransactionService)).Methods(http.MethodPost)
	// router.HandleFunc("/books", transaction.List(dep.BookService)).Methods(http.MethodGet)
	// router.HandleFunc("/books/{id}", transaction.FindByID(dep.BookService)).Methods(http.MethodGet)
	// router.HandleFunc("/books/{id}", transaction.DeleteByID(dep.BookService)).Methods(http.MethodDelete)
	// router.HandleFunc("/books", transaction.Update(dep.BookService)).Methods(http.MethodPut)
	return
}

func pingHandler(rw http.ResponseWriter, req *http.Request) {
	api.Success(rw, http.StatusOK, api.Response{Message: "pong"})
}
