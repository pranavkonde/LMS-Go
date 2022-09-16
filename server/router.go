package server

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"github.com/pranavkonde/LMS-Go/api"
	"github.com/pranavkonde/LMS-Go/book"
	"github.com/pranavkonde/LMS-Go/user"
)

const (
// versionHeader = "Accept"
)

const (
	SUPERADMIN = iota
	ADMIN
	USER
)

type JWTClaim struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}
type TokenData struct {
	Id    string
	Email string
	Role  string
}

var RoleMap = map[string]int{"superadmin": SUPERADMIN, "admin": ADMIN, "user": USER}

var jwtKey = []byte("jsd549$^&")

func Authorize(handler http.HandlerFunc, role int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		token := r.Header.Get("Authorization")

		isValid, tokenData, err := ValidateToken(token)
		// isValid = true
		fmt.Println(isValid)
		if err != nil {
			fmt.Println("error")
		}

		fmt.Println("Token Data : ", tokenData)

		if !isValid {
			//Send error response to api
			api.Error(w, http.StatusBadRequest, api.Response{Message: "Token is not valid"})
			return
		}

		tokenRole := tokenData.Role
		if RoleMap[tokenRole] > role {
			api.Error(w, http.StatusBadRequest, api.Response{Message: "You don't have the access"})
			return
		}

		handler.ServeHTTP(w, r)
		return
	}
}

func ValidateToken(tokenString string) (isValid bool, tokenData TokenData, err error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}

	isValid = true

	tokenData = TokenData{
		Id:    claims.Id,
		Email: claims.Email,
		Role:  claims.Role,
	}
	return
}

func initRouter(dep dependencies) (router *mux.Router) {
	// v1 := fmt.Sprintf("application/vnd.%s.v1", config.AppName())

	router = mux.NewRouter()
	router.HandleFunc("/ping", pingHandler).Methods(http.MethodGet)
	router.HandleFunc("/login", user.Login(dep.UserService)).Methods(http.MethodPost)
	router.HandleFunc("/users", Authorize(user.Create(dep.UserService), ADMIN)).Methods(http.MethodPost)
	router.HandleFunc("/users", Authorize(user.List(dep.UserService), ADMIN)).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", Authorize(user.FindByID(dep.UserService), USER)).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", Authorize(user.DeleteByID(dep.UserService), ADMIN)).Methods(http.MethodDelete)
	router.HandleFunc("/users", Authorize(user.Update(dep.UserService), USER)).Methods(http.MethodPut)

	router.HandleFunc("/books", Authorize(book.Create(dep.BookService), ADMIN)).Methods(http.MethodPost)
	router.HandleFunc("/books", Authorize(book.List(dep.BookService), USER)).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", Authorize(book.FindByID(dep.BookService), USER)).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", Authorize(book.DeleteByID(dep.BookService), ADMIN)).Methods(http.MethodDelete)
	router.HandleFunc("/books", Authorize(book.DeleteByID(dep.BookService), ADMIN)).Methods(http.MethodPut)

	router.HandleFunc("/userbook/issue", Authorize(book.DeleteByID(dep.BookService), ADMIN)).Methods(http.MethodPost)
	router.HandleFunc("/userbook", Authorize(book.DeleteByID(dep.BookService), ADMIN)).Methods(http.MethodGet)
	router.HandleFunc("/userbook/return", Authorize(book.DeleteByID(dep.BookService), ADMIN)).Methods(http.MethodPut)
	return

}

func pingHandler(rw http.ResponseWriter, req *http.Request) {
	api.Success(rw, http.StatusOK, api.Response{Message: "pong"})
}
