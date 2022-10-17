package book

import (
	"github.com/pranavkonde/LMS-Go/db"
)

type UpdateRequest struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Author          string `json:"author"`
	Price           int    `json:"price"`
	TotalCopies     int    `json:"totalcopies"`
	Status          string `json:"status"`
	AvailableCopies int    `json:"availablecopies"`
}

type CreateRequest struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Author          string `json:"author"`
	Price           int    `json:"price"`
	TotalCopies     int    `json:"totalcopies"`
	Status          string `json:"status"`
	AvailableCopies int    `json:"availablecopies"`
}
type bookresp struct {
	Name   string `json:"name"`
	Price  int    `json:"price"`
	Status string `json:"status"`
}

type FindByIDResponse struct {
	Book db.Book `json:"book"`
}

type ListResponse struct {
	Book []db.Book `json:"books"`
}

func (cr CreateRequest) Validate() (err error) {
	if cr.Name == "" {
		return errEmptyName
	}
	if cr.Name == "" {
		return errEmptyName
	}
	if cr.Author == "" {
		return errEmptyAuthor
	}

	if cr.TotalCopies == 0 {
		return errZeroCopies
	}
	// if !unicode.IsNumber(rune(cr.TotalCopies)) {
	// 	return errInvalidTotalCopies
	// }
	// if !unicode.IsNumber(rune(cr.AvailableCopies)) {
	// 	return errInvalidAvailableCopies
	// }
	if cr.Price < 1 {
		return errInvalidPrice
	}
	// if !unicode.IsNumber(cr.Price) {
	//  return errInvalidPrice
	// }
	// if !unicode.IsNumber(rune(cr.Price)) {
	// 	return errInvalidPrice
	// }

	if cr.Status != "Available" {
		return errInvalidStatus
	}
	if cr.AvailableCopies > cr.TotalCopies {
		return errInvalidAvailableCopies
	}
	// if !unicode.IsNumber(rune(cr.AvailableCopies)) {
	// 	return err1InvalidAvailableCopies
	// }
	return
}

func (ur UpdateRequest) Validate() (err error) {
	if ur.ID == "" {
		return errEmptyID
	}
	if ur.Name == "" {
		return errEmptyName
	}
	return
}
