package user

import (
	"Todo-list/rest/middleware"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	mux.Handle(
		"POST /users", manager.With(
			http.HandlerFunc(h.CreateUser),
		))
	mux.Handle(
		"POST /users/Login", manager.With(
			http.HandlerFunc(h.Login),
		))

}
