package db

import (
	"context"
	"database/sql"
)

type Book struct {
	ID              string `db:"id"`
	Name            string `db:"name"`
	Author          string `db:"author"`
	Price           int    `db:"price"`
	TotalCopies     int    `db:"totalcopies"`
	Status          string `db:"status"`
	AvailableCopies int    `db:"availablecopies"`
}

const (
	createBookQuery = `INSERT INTO Books (id,name,author,price,totalcopies,status,availablecopies)
    VALUES(?, ?,?,?,?,?,?)`

	listBooksQuery      = `SELECT * FROM Books`
	findBookByIDQuery   = `SELECT * FROM Books WHERE id = ?`
	deleteBookByIDQuery = `DELETE FROM Books WHERE id = ?`
	updateBookQuery     = `UPDATE Books SET name = ?, author = ?, price=?, totalcopies=?, status=?, availablecopies=?  where id = ?`
)

func (s *store) CreateBook(ctx context.Context, book *Book) (err error) {

	return Transact(ctx, s.db, &sql.TxOptions{}, func(ctx context.Context) error {
		_, err = s.db.Exec(
			createBookQuery,
			book.ID,
			book.Name,
			book.Author,
			book.Price,
			book.TotalCopies,
			book.Status,
			book.AvailableCopies,
		)
		return err
	})
}

func (s *store) ListBooks(ctx context.Context) (books []Book, err error) {
	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		return s.db.SelectContext(ctx, &books, listBooksQuery)
	})
	if err == sql.ErrNoRows {
		return books, ErrBookNotExist
	}
	return
}

func (s *store) FindBookByID(ctx context.Context, id string) (book Book, err error) {
	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		return s.db.GetContext(ctx, &book, findBookByIDQuery, id)
	})
	if err == sql.ErrNoRows {
		return book, ErrBookNotExist
	}
	return
}

func (s *store) DeleteBookByID(ctx context.Context, id string) (err error) {
	return Transact(ctx, s.db, &sql.TxOptions{}, func(ctx context.Context) error {
		res, err := s.db.Exec(deleteBookByIDQuery, id)
		cnt, err := res.RowsAffected()
		if cnt == 0 {
			return ErrBookNotExist
		}
		if err != nil {
			return err
		}
		return err
	})
}

func (s *store) UpdateBook(ctx context.Context, book *Book) (err error) {
	// now := time.Now()

	return Transact(ctx, s.db, &sql.TxOptions{}, func(ctx context.Context) error {
		_, err = s.db.Exec(
			updateBookQuery,
			book.Name,
			book.Author,
			book.Price,
			book.TotalCopies,
			book.Status,
			book.AvailableCopies,
			book.ID,
		)
		return err
	})
}
