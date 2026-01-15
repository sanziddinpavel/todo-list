package cmd

import (
	"Todo-list/config"
	"Todo-list/rest"
	"Todo-list/rest/handler/todo"
	"Todo-list/rest/handler/user"
	"Todo-list/rest/middleware"
)

func Serve() {
	cnf := config.GetConfig()
	middlewares := middleware.NewMiddleware(cnf)
	todoHandler := todo.NewHandler(middlewares)
	userHandler := user.NewHandler()
	server := rest.NewServer(cnf, todoHandler, userHandler)
	server.Start()
}
