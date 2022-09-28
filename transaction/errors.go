package transaction

import "errors"

var (
	errEmptyBookID    = errors.New("book ID must be present")
	errEmptyUserID    = errors.New("user ID must be present")
	errNoTransactions = errors.New("no Transactions present")
	errNoTransaction  = errors.New("you have already taken the book")
)
