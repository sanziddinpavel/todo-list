package todo

type Todos struct {
	ID          int    `json:"id"`
	Text        string `json:"text"`
	Description string `json:"description"`
	IsDone      bool   `json:"isDone"`
}
