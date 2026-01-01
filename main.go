package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello its my first project")
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Please give me GET request", 400)
		return
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(Todolist)

}

func createTodos(w http.ResponseWriter, r *http.Request) {

	var NewTodos Todos
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&NewTodos)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please give me valid json", 400)
		return

	}

	NewTodos.ID = len(Todolist) + 1
	Todolist = append(Todolist, NewTodos)

	incoder := json.NewEncoder(w)
	incoder.Encode(NewTodos)

}

func deleteTodos(w http.ResponseWriter, r *http.Request) {
	todoID := r.PathValue("id")
	tID, err := strconv.Atoi(todoID)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "invalid product id", 400)
		return
	}

	for i, todo := range Todolist {
		if todo.ID == tID {
			Todolist = append(Todolist[:i], Todolist[i+1:]...)
			http.Error(w, "todo deleted", 200)
		}

	}

}
func updateTodos(w http.ResponseWriter, r *http.Request) {
	todoID := r.PathValue("id")
	tID, err := strconv.Atoi(todoID)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "invalid product id", 400)
		return
	}

	for _, todo := range Todolist {
		if todo.ID == tID {

			http.Error(w, "todo", 200)
			return
		}
		http.Error(w, "todo pai ni", 400)

	}

}

func main() {
	mux := http.NewServeMux()
	mux.Handle("GET /hello", http.HandlerFunc(helloHandler))

	mux.Handle("GET /todos", http.HandlerFunc(getTodos))
	mux.Handle("POST /todos", http.HandlerFunc(createTodos))
	mux.Handle("GET /todos/{id}", http.HandlerFunc(getTodos))

	mux.Handle("PUT /todos/{id}", http.HandlerFunc(updateTodos))
	mux.Handle("DELETE /todos/{id}", http.HandlerFunc(deleteTodos))

	fmt.Println("server running on :2222")
	err := http.ListenAndServe(":2222", mux)
	if err != nil {
		fmt.Println("Error", err)
	}

}

type Todos struct {
	ID          int    `json:"id"`
	Text        string `json:"text"`
	Description string `json:"description"`
	IsDone      bool   `json:"isDone"`
}

var Todolist []Todos

func init() {
	todo1 := Todos{
		ID:          1,
		Text:        "Mango",
		Description: "I need 2 k.g mangoes",
		IsDone:      true,
	}
	todo2 := Todos{
		ID:          2,
		Text:        "potato",
		Description: " i need 3 k.g potatoes",
		IsDone:      true,
	}
	todo3 := Todos{
		ID:          3,
		Text:        "Notebook",
		Description: "I need 2 Notebooks",
		IsDone:      true,
	}

	Todolist = append(Todolist, todo1, todo2, todo3)
}
