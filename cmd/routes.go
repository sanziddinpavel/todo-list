package cmd

import (
	"Todo-list/handler"
	"Todo-list/middleware"
	"net/http"
)

func InitRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	mux.Handle(
		"GET /hello", manager.With(
			http.HandlerFunc(handler.HelloHandler),
		))

	mux.Handle(
		"GET /todos", manager.With(
			http.HandlerFunc(handler.GetTodos),
		))

	mux.Handle(
		"POST /todos", manager.With(
			http.HandlerFunc(handler.CreateTodos),
		))

	mux.Handle(
		"GET /todos/{id}", manager.With(
			http.HandlerFunc(handler.GetTodo),
		))

	mux.Handle(
		"PUT /todos/{id}", manager.With(
			http.HandlerFunc(handler.UpdateTodos),
		))

	mux.Handle(
		"DELETE /todos/{id}", manager.With(
			http.HandlerFunc(handler.DeleteTodos),
		))

}
