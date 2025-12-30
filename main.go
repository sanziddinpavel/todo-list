package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello its my first project")
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("GET /hello", http.HandlerFunc(helloHandler))

	fmt.Println("server running on :2222")
	err := http.ListenAndServe(":2222", mux)
	if err != nil {
		fmt.Println("Error", err)
	}

}
