package rest

import (
	"Todo-list/config"

	"Todo-list/rest/handler/todo"
	"Todo-list/rest/handler/user"
	"Todo-list/rest/middleware"
	"fmt"
	"net/http"
	"strconv"
)

type Server struct {
	cnf         *config.Config
	todoHandler *todo.Handler
	userHandler *user.Handler
}

func NewServer(
	cnf *config.Config,
	todoHandler *todo.Handler,
	userHandler *user.Handler,
) *Server {
	return &Server{
		cnf:         cnf,
		todoHandler: todoHandler,
		userHandler: userHandler,
	}
}

func (s *Server) Start() {
	manager := middleware.NewManager()
	manager.Use(
		middleware.Logger,
		middleware.CorsMiddleware,
		middleware.Preflight)

	mux := http.NewServeMux()
	Wrapmux := manager.Wrapmux(mux)

	s.todoHandler.RegisterRoutes(mux, manager)
	s.userHandler.RegisterRoutes(mux, manager)

	addr := ":" + strconv.Itoa(s.cnf.HttpPort)
	fmt.Println("server running on", addr)
	err := http.ListenAndServe(addr, Wrapmux)
	if err != nil {
		fmt.Println("Error", err)

	}
}
