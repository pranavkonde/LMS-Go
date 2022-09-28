package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

const (
	createTransactionQuery = `INSERT INTO transaction ( 
        id,issue_date,due_date,return_date,book_id,user_id)
        VALUES(?, ?, ?, ?, ?, ?)`
	listTransactionsQuery  = `SELECT * FROM transaction`
	updateTransactionQuery = `UPDATE transaction SET return_date=? WHERE book_id=? AND user_id=? AND return_date=0 `
	issueCopyQuery         = `UPDATE books SET availablecopies=availablecopies-1 WHERE id = ? AND availablecopies>0`
	returnCopyQuery        = `UPDATE books SET availablecopies=availablecopies+1 WHERE id = ?`
	BookStatusQuery        = `SELECT return_date from transaction WHERE book_id = ? AND user_id =?`
	GetTotalCopiesQuery    = `SELECT books.totalcopies FROM books LEFT JOIN transaction ON books.id =transaction.book_id`
	GetCurrentCopiesQuery  = `SELECT books.availablecopies FROM books LEFT JOIN transaction ON books.id =transaction.book_id where books.id=?`
)

type Transaction struct {
	ID         string `db:"id"`
	IssueDate  int    `db:"issue_date"`
	DueDate    int    `db:"due_date"`
	ReturnDate int    `db:"return_date"`
	BookID     string `db:"book_id"`
	UserID     string `db:"user_id"`
}

func (s *store) CreateTransaction(ctx context.Context, transaction *Transaction) (err error) {
	now := time.Now().UTC().Unix()
	transaction.DueDate = int(now) + 864000

	return Transact(ctx, s.db, &sql.TxOptions{}, func(ctx context.Context) error {
		_, err = s.db.Exec(
			createTransactionQuery,
			transaction.ID,
			now,
			transaction.DueDate,
			0,
			transaction.BookID,
			transaction.UserID,
		)
		if err != nil {
			return err
		}

		_, err = s.db.Exec(
			issueCopyQuery,
			transaction.BookID,
		)
		return err
	})
}

func (s *store) ListTransactions(ctx context.Context) (transactions []Transaction, err error) {
	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		return s.db.SelectContext(ctx, &transactions, listTransactionsQuery)
	})
	if err == sql.ErrNoRows {
		return transactions, ErrTransactionNotExist
	}
	return
}

func (s *store) UpdateTransaction(ctx context.Context, transaction *Transaction) (err error) {

	return Transact(ctx, s.db, &sql.TxOptions{}, func(ctx context.Context) error {
		_, err = s.db.Exec(
			updateTransactionQuery,
			transaction.ReturnDate,
			transaction.BookID,
			transaction.UserID,
		)
		if err != nil {
			return err
		}

		_, err = s.db.Exec(
			returnCopyQuery,
			transaction.BookID,
		)
		return err
	})
}
func (s *store) BookStatus(ctx context.Context, BookId string, UserID string) (res string, err error) {
	return_date := -1
	s.db.GetContext(ctx, &return_date, BookStatusQuery, BookId, UserID)
	fmt.Println(return_date)
	if return_date == 0 {
		res = "issued"
		return res, nil
	} else {
		totalcnt := 0
		currentcnt := 0
		s.db.GetContext(ctx, &totalcnt, GetTotalCopiesQuery, BookId)
		s.db.GetContext(ctx, &currentcnt, GetCurrentCopiesQuery, BookId)

		if currentcnt < 1 {
			res = "Unavailable"
			return res, nil
		} else {
			res = "Available"
			return res, nil
		}
	}
}
