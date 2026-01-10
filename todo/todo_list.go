package todo

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
