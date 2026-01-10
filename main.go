package main

import (
	"Todo-list/handler"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("GET /hello", http.HandlerFunc(handler.HelloHandler))

	mux.Handle("GET /todos", http.HandlerFunc(handler.GetTodos))
	mux.Handle("POST /todos", http.HandlerFunc(handler.CreateTodos))
	mux.Handle("GET /todos/{id}", http.HandlerFunc(handler.GetTodo))

	mux.Handle("PUT /todos/{id}", http.HandlerFunc(handler.UpdateTodos))
	mux.Handle("DELETE /todos/{id}", http.HandlerFunc(handler.DeleteTodos))

	fmt.Println("server running on :2222")
	err := http.ListenAndServe(":2222", mux)
	if err != nil {
		fmt.Println("Error", err)
	}

}
