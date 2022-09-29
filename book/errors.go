package book

import "errors"

var (
	errEmptyID                 = errors.New("book ID must be present")
	errEmptyName               = errors.New("book name must be present")
	errNoBooks                 = errors.New("no Books present")
	errEmptyAuthor             = errors.New("author name must be present")
	errZeroCopies              = errors.New("copies cannot be zero while creation of book")
	errInvalidPrice            = errors.New("price should be greater than zero")
	errInvalidStatus           = errors.New("invalid Status of book")
	errInvalidAvailableCopies  = errors.New("available copies cannot be greater than total copies")
	errInvalidTotalCopies      = errors.New("total copies must be integer")
	err1InvalidAvailableCopies = errors.New("available copies must be integer")
	errNoBookId                = errors.New("book is not present")
)
