package server

import (
	"fmt"
	"strconv"

	"github.com/pranavkonde/Library-Management-System-Golang/config"
	"github.com/urfave/negroni"
)

func StartAPIServer() {
	port := config.AppPort()
	server := negroni.Classic()

	dependencies, err := initDependencies()
	if err != nil {
		panic(err)
	}

	router := initRouter(dependencies)
	server.UseHandler(router)
	addr := fmt.Sprintf(":%s", strconv.Itoa(port))
	server.Run(addr)
}
