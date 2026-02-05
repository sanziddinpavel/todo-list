package repo

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Todos struct {
	ID          int    `db:"id" json:"id"`
	Text        string `db:"text" json:"text"`
	Description string `db:"description" json:"description"`
	IsDone      bool   `db:"is_done" json:"is_done"`
}

type TodoRepo interface {
	Create(t Todos) (*Todos, error)
	Get(todoID int) (*Todos, error)
	List() ([]*Todos, error)
	Delete(todoID int) error
	Update(t Todos) (*Todos, error)
}

type todoRepo struct {
	db *sqlx.DB
}

func NewTodoRepo(db *sqlx.DB) TodoRepo {
	return &todoRepo{
		db: db,
	}

}

func (r *todoRepo) Create(t Todos) (*Todos, error) {
	query := `

	INSERT INTO todos(
	text,
	description,
	is_done

	)VALUES(
	$1,
	$2,
	$3
	)
	RETURNING id
	`
	row := r.db.QueryRow(query, t.Text, t.Description, t.IsDone)
	err := row.Scan(&t.ID)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *todoRepo) Get(id int) (*Todos, error) {
	var td Todos
	query := `
	SELECT 
	id,
	text,
	description,
	is_done
	from todos
	where id = $1
	`
	err := r.db.Get(&td, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, nil
	}

	return &td, nil
}

func (r *todoRepo) List() ([]*Todos, error) {
	var tdList []*Todos
	query := `
	SELECT 
	id,
	text,
	description,
	is_done
	from todos
	
	`
	err := r.db.Select(&tdList, query)
	if err != nil {

		return nil, nil
	}

	return tdList, nil

}

func (r *todoRepo) Delete(id int) error {
	query := `
		DELETE FROM todos WHERE id=$1
	`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err

	}
	return nil
}

func (r *todoRepo) Update(t Todos) (*Todos, error) {
	query := `
UPDATE todos
SET text=$1,
	description=$2,
	is_done=$3
	WHERE id =$4
`
	row := r.db.QueryRow(query, t.Text, t.Description, t.IsDone)
	err := row.Err()
	if err != nil {
		return nil, err
	}
	return &t, nil

}
