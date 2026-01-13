package rest

import (
	"Todo-list/config"

	"Todo-list/rest/middleware"
	"fmt"
	"net/http"
	"strconv"
)

func Start(cnf config.Config) {
	manager := middleware.NewManager()
	manager.Use(
		middleware.Logger,
		middleware.CorsMiddleware,
		middleware.Preflight)

	mux := http.NewServeMux()
	Wrapmux := manager.Wrapmux(mux)

	InitRoutes(mux, manager)

	addr := ":" + strconv.Itoa(cnf.HttpPort)
	fmt.Println("server running on", addr)
	err := http.ListenAndServe(addr, Wrapmux)
	if err != nil {
		fmt.Println("Error", err)

	}
}
