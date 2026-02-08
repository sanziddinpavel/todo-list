package cmd

import (
	"Todo-list/config"
	"Todo-list/infra/db"
	"Todo-list/repo"
	"Todo-list/rest"
	todoHandler "Todo-list/rest/handler/todo"
	userHandler "Todo-list/rest/handler/user"
	"Todo-list/rest/middleware"
	"Todo-list/user"
	"fmt"
	"os"
)

func Serve() {
	cnf := config.GetConfig()

	dbcon, err := db.NewConnetion(&cnf.DB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = db.MigrateDB(dbcon, "./migrations")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	todoRepo := repo.NewTodoRepo(dbcon)
	userRepo := repo.NewUserRepo(dbcon)

	usrSvc := user.NewService(&userRepo)

	middlewares := middleware.NewMiddleware(cnf)

	todoHandler := todoHandler.NewHandler(middlewares, todoRepo)
	userHandler := userHandler.NewHandler(cnf, usrSvc)

	server := rest.NewServer(cnf, todoHandler, userHandler)
	server.Start()
}
