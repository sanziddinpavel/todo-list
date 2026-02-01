package cmd

import (
	"Todo-list/config"
	"Todo-list/repo"
	"Todo-list/rest"
	"Todo-list/rest/handler/todo"
	"Todo-list/rest/handler/user"
	"Todo-list/rest/middleware"
)

func Serve() {
	cnf := config.GetConfig()

	middlewares := middleware.NewMiddleware(cnf)

	todoRepo := repo.NewTodoRepo()
	todoHandler := todo.NewHandler(middlewares, todoRepo)

	userRepo := repo.NewUserRepo()
	userHandler := user.NewHandler(cnf, &userRepo)

	server := rest.NewServer(cnf, todoHandler, userHandler)
	server.Start()
}
