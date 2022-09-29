package user

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/pranavkonde/LMS-Go/api"
	"github.com/pranavkonde/LMS-Go/db"
)

type Authentication struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var j Authentication
		err := json.NewDecoder(req.Body).Decode(&j)
		if err != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}
		jwtString, err1 := service.GenerateJWT(req.Context(), j.Email, j.Password)
		if err1 != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err1.Error()})
			return
		}

		api.Success(rw, http.StatusCreated, jwtString)

	})
}

func Create(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// api.Success(rw, http.StatusOK, api.Response{Message: "hi"})
		var c CreateRequest
		err := json.NewDecoder(req.Body).Decode(&c)
		if err != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}

		err = c.Validate()
		if err != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}

		err = service.Create(req.Context(), c)
		if isBadRequest(err) {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}

		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			return
		}

		api.Success(rw, http.StatusCreated, api.Response{Message: "Created Successfully"})
	})
}

func List(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// api.Success(rw, http.StatusCreated, api.Response{Message: "Enter List"})
		resp, err := service.List(req.Context())
		if err == errNoUsers {
			// api.Success(rw, http.StatusCreated, api.Response{Message: "Enter List"})
			api.Error(rw, http.StatusNotFound, api.Response{Message: err.Error()})
			return
		}
		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			return
		}

		api.Success(rw, http.StatusOK, resp)
	})
}

func FindByID(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)

		resp, err := service.FindByID(req.Context(), vars["id"])

		if err == errNoUserId {
			api.Error(rw, http.StatusNotFound, api.Response{Message: err.Error()})
			return
		}
		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			return
		}

		api.Success(rw, http.StatusOK, resp)
	})
}

func DeleteByID(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)

		err := service.DeleteByID(req.Context(), vars["id"])
		if err == errNoUserId {
			api.Error(rw, http.StatusNotFound, api.Response{Message: err.Error()})
			return
		}
		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			return
		}

		api.Success(rw, http.StatusOK, api.Response{Message: "Deleted Successfully"})
	})
}

var v db.User
var flag = 0

func Update(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var c UpdateRequest
		err := json.NewDecoder(req.Body).Decode(&c)
		if err != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}

		err = service.Update(req.Context(), c)
		if isBadRequest(err) {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}

		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			return
		}

		api.Success(rw, http.StatusOK, api.Response{Message: "Updated Successfully"})
	})
}

func UpdatePassword(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var c UpdatePasswordStruct
		resp, err := service.List(req.Context())
		if err != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}

		err = json.NewDecoder(req.Body).Decode(&c)

		if err != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}

		for _, v = range resp.User {
			if v.ID == c.ID && v.Password == c.Password {
				flag = 1
				err = service.UpdatePassword(req.Context(), c)
				if isBadRequest(err) {
					api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
					return
				}

				if err != nil {
					api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
					return
				}

				api.Success(rw, http.StatusOK, api.Response{Message: "Updated Successfully"})
			}
		}
		if flag != 1 {
			api.Success(rw, http.StatusOK, api.Response{Message: "Wrong ID or Password"})
		}

	})
}

func isBadRequest(err error) bool {
	return err == errEmptyName || err == errEmptyID
}
