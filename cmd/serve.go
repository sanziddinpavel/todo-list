package cmd

import (
	"Todo-list/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	manager := middleware.NewManager()
	manager.Use(middleware.Logger, middleware.CorsMiddleware, middleware.Preflight)

	mux := http.NewServeMux()
	InitRoutes(mux, manager)

	fmt.Println("server running on :2222")
	err := http.ListenAndServe(":2222", mux)
	if err != nil {
		fmt.Println("Error", err)
	}
}
