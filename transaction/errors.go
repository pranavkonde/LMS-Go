package transaction

import "errors"

var (
	errEmptyBookID    = errors.New("Book ID must be present")
	errEmptyUserID    = errors.New("User ID must be present")
	errNoTransactions = errors.New("No Transactions present")
)
