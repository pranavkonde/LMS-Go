package book

import (
	"context"

	"github.com/google/uuid"
	"github.com/pranavkonde/LMS-Go/db"
	"go.uber.org/zap"
)

type Service interface {
	List(ctx context.Context) (response ListResponse, err error)
	Create(ctx context.Context, req CreateRequest) (err error)
	FindByID(ctx context.Context, id string) (response FindByIDResponse, err error)
	DeleteByID(ctx context.Context, id string) (err error)
	Update(ctx context.Context, req UpdateRequest) (err error)
}

type bookService struct {
	store  db.Storer
	logger *zap.SugaredLogger
}

func (cs *bookService) List(ctx context.Context) (response ListResponse, err error) {
	books, err := cs.store.ListBooks(ctx)
	if err == db.ErrBookNotExist {
		cs.logger.Error("No book present", "err", err.Error())
		return response, errNoBooks
	}
	if err != nil {
		cs.logger.Error("Error listing books", "err", err.Error())
		return
	}

	response.Book = books
	return
}

func (cs *bookService) Create(ctx context.Context, c CreateRequest) (err error) {
	err = c.Validate()
	if err != nil {
		cs.logger.Errorw("Invalid request for book create", "msg", err.Error(), "book", c)
		return
	}
	uuidgen := uuid.New()
	c.ID = uuidgen.String()
	err = cs.store.CreateBook(ctx, &db.Book{

		ID:              c.ID,
		Name:            c.Name,
		Author:          c.Author,
		Price:           c.Price,
		TotalCopies:     c.TotalCopies,
		Status:          c.Status,
		AvailableCopies: c.AvailableCopies,
	})
	if err != nil {
		cs.logger.Error("Error creating book", "err", err.Error())
		return
	}
	return
}

func (cs *bookService) Update(ctx context.Context, c UpdateRequest) (err error) {
	err = c.Validate()
	if err != nil {
		cs.logger.Error("Invalid Request for book update", "err", err.Error(), "book", c)
		return
	}

	err = cs.store.UpdateBook(ctx, &db.Book{
		ID:              c.ID,
		Name:            c.Name,
		Author:          c.Author,
		Price:           c.Price,
		TotalCopies:     c.TotalCopies,
		Status:          c.Status,
		AvailableCopies: c.AvailableCopies,
	})
	if err != nil {
		cs.logger.Error("Error updating book", "err", err.Error(), "book", c)
		return
	}

	return
}

func (cs *bookService) FindByID(ctx context.Context, id string) (response FindByIDResponse, err error) {
	book, err := cs.store.FindBookByID(ctx, id)
	if err == db.ErrBookNotExist {
		cs.logger.Error("No book present", "err", err.Error())
		return response, errNoBookId
	}
	if err != nil {
		cs.logger.Error("Error finding book", "err", err.Error(), "id", id)
		return
	}

	response.Book = book
	return
}

func (cs *bookService) DeleteByID(ctx context.Context, id string) (err error) {
	err = cs.store.DeleteBookByID(ctx, id)
	if err == db.ErrBookNotExist {
		cs.logger.Error("Book Not present", "err", err.Error(), "id", id)
		return errNoBookId
	}
	if err != nil {
		cs.logger.Error("Error deleting book", "err", err.Error(), "id", id)
		return
	}

	return
}

func NewService(s db.Storer, l *zap.SugaredLogger) Service {
	return &bookService{
		store:  s,
		logger: l,
	}
}
