package domain

type Todos struct {
	ID          int    `db:"id" json:"id"`
	Text        string `db:"text" json:"text"`
	Description string `db:"description" json:"description"`
	IsDone      bool   `db:"is_done" json:"is_done"`
}
