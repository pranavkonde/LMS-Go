package user

import "errors"

var (
	errEmptyID          = errors.New("user ID must be present")
	errEmptyName        = errors.New("user name must be present")
	errNoUsers          = errors.New("no users present")
	errNoUserId         = errors.New("user is not present")
	errEmptyPassword    = errors.New("password cannot be empty")
	errEmptyGender      = errors.New("user gender must be present")
	errEmptyAddress     = errors.New("address must be present")
	errEmptyEmail       = errors.New("email must be present")
	errEmptyMobNo       = errors.New("mob no must be present")
	errEmptyRole        = errors.New("role must be present")
	errRoleType         = errors.New("enter a valid Role type")
	errWrongPassword    = errors.New("wrong password")
	errInvalidFirstName = errors.New("invalid first name")
	errInvalidGender    = errors.New("invalid gender")
	errInvalidLastName  = errors.New("invalid last name")
	errNotValidMail     = errors.New("email is not valid")
	errInvalidMobNo     = errors.New("mob Number is not valid")
	errEmptyLastName    = errors.New("last name cannot be Empty")
)
