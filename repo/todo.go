package repo

import (
	"Todo-list/domain"
	"database/sql"

	"Todo-list/todo"

	"github.com/jmoiron/sqlx"
)

type TodoRepo interface {
	todo.TodoRepo
}

type todoRepo struct {
	db *sqlx.DB
}

func NewTodoRepo(db *sqlx.DB) TodoRepo {
	return &todoRepo{
		db: db,
	}

}

func (r *todoRepo) Create(t domain.Todos) (*domain.Todos, error) {
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

func (r *todoRepo) Get(id int) (*domain.Todos, error) {
	var td domain.Todos
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

func (r *todoRepo) List(page, limit int64) ([]*domain.Todos, error) {

	offset := ((page - 1) * limit) + 1
	var tdList []*domain.Todos
	query := `
	SELECT 
	id,
	text,
	description,
	is_done
	from todos
	LIMIT $1 OFFSET $2
	`
	err := r.db.Select(&tdList, query, page, offset)
	if err != nil {

		return nil, nil
	}

	return tdList, nil

}
func (r *todoRepo) Count() (int64, error) {

	query := `
	SELECT COUNT(*)
	FROM todos
	
	`
	var count int64
	err := r.db.QueryRow(query).Scan(&count)
	if err != nil {

		return 0, err
	}

	return count, nil

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

func (r *todoRepo) Update(t domain.Todos) (*domain.Todos, error) {
	query := `
UPDATE todos
SET text=$1,
	description=$2,
	is_done=$3
	WHERE id =$4
`
	row := r.db.QueryRow(query, t.Text, t.Description, t.IsDone, t.ID)
	err := row.Err()
	if err != nil {
		return nil, err
	}
	return &t, nil

}
