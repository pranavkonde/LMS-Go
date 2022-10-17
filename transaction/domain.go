package transaction

type Transaction struct {
	ID         string `json:"id"`
	IssueDate  string `json:"issue_date"`
	DueDate    string `json:"due_date"`
	ReturnDate string `json:"return_date"`
	BookID     string `json:"book_id"`
	UserID     string `json:"user_id"`
}

type listResponse struct {
	Transactions []Transaction `json:"transactions"`
	// Count        int           `json:"total_count"` //to be updated
}
type RequestStatus struct {
	UserID string `json:"user_id"`
	BookID string `json:"book_id"`
}

func (cr Transaction) Validate() (err error) {
	if cr.BookID == "" {
		return errEmptyBookID
	}
	if cr.UserID == "" {
		return errEmptyUserID
	}
	return
}
