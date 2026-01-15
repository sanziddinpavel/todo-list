package todo

import (
	"Todo-list/rest/middleware"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	mux.Handle(
		"GET /hello", manager.With(
			http.HandlerFunc(h.HelloHandler),
		))

	mux.Handle(
		"GET /todos", manager.With(
			http.HandlerFunc(h.GetTodos),
		))

	mux.Handle(
		"POST /todos", manager.With(
			http.HandlerFunc(h.CreateTodos),
			h.middlewares.AuthenticationJWT,
		))

	mux.Handle(
		"GET /todos/{id}", manager.With(
			http.HandlerFunc(h.GetTodo),
			h.middlewares.AuthenticationJWT,
		))

	mux.Handle(
		"PUT /todos/{id}", manager.With(
			http.HandlerFunc(h.UpdateTodos),
			h.middlewares.AuthenticationJWT,
		))

	mux.Handle(
		"DELETE /todos/{id}", manager.With(
			http.HandlerFunc(h.DeleteTodos),
			h.middlewares.AuthenticationJWT,
		))

}
