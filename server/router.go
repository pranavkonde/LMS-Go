package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pranavkonde/LMS-Go/api"
	"github.com/pranavkonde/LMS-Go/book"
	"github.com/pranavkonde/LMS-Go/user"
)

const (
// versionHeader = "Accept"
)

func initRouter(dep dependencies) (router *mux.Router) {
	// v1 := fmt.Sprintf("application/vnd.%s.v1", config.AppName())

	router = mux.NewRouter()
	router.HandleFunc("/ping", pingHandler).Methods(http.MethodGet)
	router.HandleFunc("/login", user.Login()).Methods(http.MethodPost)
	router.HandleFunc("/users", user.Authorize(user.Create(dep.UserService), "RoleAdmin")).Methods(http.MethodPost)
	router.HandleFunc("/users", user.Authorize(user.List(dep.UserService), "RoleAdmin")).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", user.Authorize(user.FindByID(dep.UserService), "UserService")).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", user.Authorize(user.DeleteByID(dep.UserService), "RoleAdmin")).Methods(http.MethodDelete)
	router.HandleFunc("/users", user.Authorize(user.Update(dep.UserService), "RoleUser")).Methods(http.MethodPut)

	router.HandleFunc("/books", user.Authorize(book.Create(dep.BookService), "RoleAdmin")).Methods(http.MethodPost)
	router.HandleFunc("/books", user.Authorize(book.List(dep.BookService), "RoleUser")).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", user.Authorize(book.FindByID(dep.BookService), "RoleUser")).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", user.Authorize(book.DeleteByID(dep.BookService), "Admin")).Methods(http.MethodDelete)
	router.HandleFunc("/books", user.Authorize(book.DeleteByID(dep.BookService), "RoleAdmin")).Methods(http.MethodPut)

	router.HandleFunc("/userbook/issue", user.Authorize(book.DeleteByID(dep.BookService), "RoleAdmin")).Methods(http.MethodPost)
	router.HandleFunc("/userbook", user.Authorize(book.DeleteByID(dep.BookService), "RoleAdmin")).Methods(http.MethodGet)
	router.HandleFunc("/userbook/return", user.Authorize(book.DeleteByID(dep.BookService), "RoleAdmin")).Methods(http.MethodPut)
	return

}

func pingHandler(rw http.ResponseWriter, req *http.Request) {
	api.Success(rw, http.StatusOK, api.Response{Message: "pong"})
}
