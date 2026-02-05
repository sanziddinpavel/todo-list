package cmd

import (
	"Todo-list/config"
	"Todo-list/infra/db"
	"Todo-list/repo"
	"Todo-list/rest"
	"Todo-list/rest/handler/todo"
	"Todo-list/rest/handler/user"
	"Todo-list/rest/middleware"
	"fmt"
	"os"
)

func Serve() {
	cnf := config.GetConfig()

	dbcon, err := db.NewConnetion()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	middlewares := middleware.NewMiddleware(cnf)

	todoRepo := repo.NewTodoRepo(dbcon)
	todoHandler := todo.NewHandler(middlewares, todoRepo)

	userRepo := repo.NewUserRepo(dbcon)
	userHandler := user.NewHandler(cnf, &userRepo)

	server := rest.NewServer(cnf, todoHandler, userHandler)
	server.Start()
}
