package user

import "github.com/pranavkonde/LMS-Go/db"

type updateRequest struct {
    ID        string `json:"id"`
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
    Gender    string `json:"gender"`
    Age       int    `json:"age"`
    Address   string `json:"address"`
    Email     string `json:"email"`
    Password  string `json:"password"`
    MobileNum string `json:"mob_no"`
    Role      string `json:"role"`
}

type createRequest struct {
    ID        string `json:"id"`
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
    Gender    string `json:"gender"`
    Age       int    `json:"age"`
    Address   string `json:"address"`
    Email     string `json:"email"`
    Password  string `json:"password"`
    MobileNum string `json:"mob_no"`
    Role      string `json:"role"`
}

type findByIDResponse struct {
    User db.User `json:"user"`
}

type listResponse struct {
    User []db.User `json:"users"`
}

func (cr createRequest) Validate() (err error) {
    if cr.FirstName == "" {
        return errEmptyName
    }
    return
}

func (ur updateRequest) Validate() (err error) {
    if ur.ID == "" {
        return errEmptyID
    }
    if ur.FirstName == "" {
        return errEmptyName
    }
    return
}