package server

import (
	"github.com/pranavkonde/LMS-Go/app"
	"github.com/pranavkonde/LMS-Go/book"
	"github.com/pranavkonde/LMS-Go/db"
	"github.com/pranavkonde/LMS-Go/user"
)

type dependencies struct {
	UserService user.Service
	BookService book.Service
}

func initDependencies() (dependencies, error) {
	appDB := app.GetDB()
	logger := app.GetLogger()
	dbStore := db.NewStorer(appDB)

	userService := user.NewService(dbStore, logger)
	bookService := book.NewService(dbStore, logger)

	return dependencies{
		UserService: userService,
		BookService: bookService,
	}, nil
}
