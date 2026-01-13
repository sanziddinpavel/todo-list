package database

var todoList []Todos

type Todos struct {
	ID          int    `json:"id"`
	Text        string `json:"text"`
	Description string `json:"description"`
	IsDone      bool   `json:"isDone"`
}

func Store(t Todos) Todos {
	t.ID = len(todoList) + 1
	todoList = append(todoList, t)
	return t
}

func List() []Todos {
	return todoList
}

func Get(todoID int) *Todos {
	for _, todo := range todoList {
		if todo.ID == todoID {
			return &todo
		}
	}
	return nil
}

func Update(todo Todos) {
	for idx, t := range todoList {
		if t.ID == todo.ID {
			todoList[idx] = todo
		}
	}
}

func Delete(todoID int) bool {
	tempList := make([]Todos, 0)
	for _, t := range todoList {
		if t.ID != todoID {
			tempList = append(tempList, t)

		}
	}
	if len(tempList) == len(todoList) {
		return false // not found
	}

	todoList = tempList
	return true

}

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

	todoList = append(todoList, todo1, todo2, todo3)
}
