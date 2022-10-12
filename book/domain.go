package book

import (
	"github.com/pranavkonde/LMS-Go/db"
)

type updateRequest struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Author          string `json:"author"`
	Price           int    `json:"price"`
	TotalCopies     int    `json:"totalcopies"`
	Status          string `json:"status"`
	AvailableCopies int    `json:"availablecopies"`
}

type createRequest struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Author          string `json:"author"`
	Price           int    `json:"price"`
	TotalCopies     int    `json:"totalcopies"`
	Status          string `json:"status"`
	AvailableCopies int    `json:"availablecopies"`
}

type findByIDResponse struct {
	Book db.Book `json:"book"`
}

type listResponse struct {
	Book []db.Book `json:"books"`
}

func (cr createRequest) Validate() (err error) {
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

func (ur updateRequest) Validate() (err error) {
	if ur.ID == "" {
		return errEmptyID
	}
	if ur.Name == "" {
		return errEmptyName
	}
	return
}
