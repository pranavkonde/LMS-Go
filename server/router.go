package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pranavkonde/Library-Managemant-System-Golang/api"
	"github.com/pranavkonde/Library-Management-System-Golang/config"
	"github.com/pranavkonde/Library-Management-System-Golang/user"
)

const (
	versionHeader = "Accept"
)

func initRouter(dep dependencies) (router *mux.Router) {
	v1 := fmt.Sprintf("application/vnd.%s.v1", config.AppName())

	router = mux.NewRouter()
	router.HandleFunc("/ping", pingHandler).Methods(http.MethodGet)

	//User

	router.HandleFunc("/users", user.Create(dep.UserService)).Methods(http.MethodPost).Headers(versionHeader, v1)
	router.HandleFunc("/users", user.List(dep.UserService)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/users/{user_id}", user.FindByID(dep.UserService)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/users/{user_id}", user.DeleteByID(dep.UserService)).Methods(http.MethodDelete).Headers(versionHeader, v1)
	router.HandleFunc("/users", user.Update(dep.UserService)).Methods(http.MethodPut).Headers(versionHeader, v1)

	sh := http.StripPrefix("/docs/", http.FileServer(http.Dir("./swaggerui/")))
	router.PathPrefix("/docs/").Handler(sh)
	return
}

func pingHandler(rw http.ResponseWriter, req *http.Request) {
	api.Success(rw, http.StatusOK, api.Response{Message: "pong"})
}
