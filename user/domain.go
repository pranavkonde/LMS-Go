package user

import (
	"net/mail"
	"unicode"

	"github.com/pranavkonde/LMS-Go/db"
)

type UpdateRequest struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    string `json:"gender"`
	Age       int    `json:"age"`
	Address   string `json:"address"`
	//  Email     string `json:"email"`
	Password  string `json:"password"`
	MobileNum string `json:"mob_no"`
	//Role      string `json:"role"`
}

type CreateRequest struct {
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
type UpdatePasswordStruct struct {
	ID          string `json:"id"`
	Password    string `json:"password"`
	NewPassword string `json:"new_password"`
}

type FindByIDResponse struct {
	User db.User `json:"user"`
}

type ListResponse struct {
	User []db.User `json:"users"`
}

func (cr CreateRequest) Validate() (err error) {
	if cr.FirstName == "" {
		return errEmptyName
	}
	for _, r := range cr.FirstName {
		if !unicode.IsLetter(r) {
			return errInvalidFirstName
		}
	}
	if cr.LastName == "" {
		return errEmptyLastName
	}
	for _, r := range cr.LastName {
		if !unicode.IsLetter(r) {
			return errInvalidLastName
		}
	}
	if cr.Password == "" {
		return errEmptyPassword
	}
	if cr.Gender == "" || cr.Gender != "Male" && cr.Gender != "male" && cr.Gender != "Female" && cr.Gender != "female" && cr.Gender != "other" && cr.Gender != "Other" {
		return errInvalidGender
	}
	if cr.Address == "" {
		return errEmptyAddress
	}
	if cr.Email == "" {
		return errEmptyEmail
	}
	if cr.MobileNum == "" {
		return errEmptyMobNo
	}
	if cr.Role == "" {
		return errEmptyRole
	}
	if cr.Role != "user" && cr.Role != "admin" {
		return errRoleType
	}
	_, b := mail.ParseAddress(cr.Email)
	if b != nil {
		return errNotValidMail
	}
	checkEmail := cr.Email
	flag := false
	lastapperance := 0
	for i := 0; i < len(checkEmail); i++ {
		if checkEmail[i] == '@' {
			flag = true
			lastapperance = i
		}
	}
	if !flag {
		return errNotValidMail
	}
	flag = false
	for i := lastapperance; i < len(checkEmail); i++ {
		if checkEmail[i] == '.' {
			flag = true
		}
	}
	if !flag {
		return errNotValidMail
	}

	if len(cr.MobileNum) < 10 || len(cr.MobileNum) > 10 {
		return errInvalidMobNo
	}
	return
}

func (ur UpdateRequest) Validate() (err error) {
	if ur.ID == "" {
		return errEmptyID
	}
	if ur.FirstName == "" {
		return errEmptyName
	}
	for _, r := range ur.FirstName {
		if !unicode.IsLetter(r) {
			return errInvalidFirstName
		}
	}
	if ur.LastName == "" {
		return errEmptyLastName
	}
	for _, r := range ur.LastName {
		if !unicode.IsLetter(r) {
			return errInvalidLastName
		}
	}
	if ur.Gender == "" || ur.Gender != "Male" && ur.Gender != "male" && ur.Gender != "Female" && ur.Gender != "female" && ur.Gender != "other" && ur.Gender != "Other" {
		return errInvalidGender
	}
	if ur.Address == "" {
		return errEmptyAddress
	}
	if ur.MobileNum == "" {
		return errEmptyMobNo
	}
	if len(ur.MobileNum) < 10 || len(ur.MobileNum) > 10 {
		return errInvalidMobNo
	}
	return
}
