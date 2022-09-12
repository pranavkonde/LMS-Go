package server

import (
	"github.com/pranavkonde/Library-Managemant-System-Golang/app"
	"github.com/pranavkonde/Library-Management-System-Golang/db"
	"github.com/pranavkonde/Library-Management-System-Golang/user"
)

type dependencies struct {
	UserService user.Service
}

func initDependencies() (dependencies, error) {
	appDB := app.GetDB()
	logger := app.GetLogger()
	dbStore := db.NewStorer(appDB)

	userService := user.NewService(dbStore, logger)

	return dependencies{
		UserService: userService,
	}, nil
}
