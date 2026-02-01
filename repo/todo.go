package repo

type Todos struct {
	ID          int    `json:"id"`
	Text        string `json:"text"`
	Description string `json:"description"`
	IsDone      bool   `json:"isDone"`
}

type TodoRepo interface {
	Create(t Todos) (*Todos, error)
	Get(todoID int) (*Todos, error)
	List() ([]*Todos, error)
	Delete(todoID int) error
	Update(t Todos) (*Todos, error)
}

type todoRepo struct {
	todoList []*Todos
}

func NewTodoRepo() TodoRepo {
	repo := &todoRepo{}
	generateInitTodos(repo)
	return repo
}

func (r *todoRepo) Create(t Todos) (*Todos, error) {
	t.ID = len(r.todoList) + 1
	r.todoList = append(r.todoList, &t)
	return &t, nil

}

func (r *todoRepo) Get(todoID int) (*Todos, error) {
	for _, todo := range r.todoList {
		if todo.ID == todoID {
			return todo, nil
		}
	}
	return nil, nil

}

func (r *todoRepo) List() ([]*Todos, error) {

	return r.todoList, nil

}

func (r *todoRepo) Delete(todoID int) error {
	var tempList []*Todos
	for _, t := range r.todoList {
		if t.ID != todoID {
			tempList = append(tempList, t)
		}
	}
	r.todoList = tempList
	return nil
}

func (r *todoRepo) Update(todo Todos) (*Todos, error) {
	for idx, t := range r.todoList {
		if t.ID == t.ID {
			r.todoList[idx] = &todo
		}
	}
	return &todo, nil

}

func generateInitTodos(t *todoRepo) {

	todo1 := &Todos{
		ID:          1,
		Text:        "Mango",
		Description: "I need 2 k.g mangoes",
		IsDone:      true,
	}
	todo2 := &Todos{
		ID:          2,
		Text:        "potato",
		Description: " i need 3 k.g potatoes",
		IsDone:      true,
	}
	todo3 := &Todos{
		ID:          3,
		Text:        "Notebook",
		Description: "I need 2 Notebooks",
		IsDone:      true,
	}

	t.todoList = append(t.todoList, todo1, todo2, todo3)
}
