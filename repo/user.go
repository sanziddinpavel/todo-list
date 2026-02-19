package repo

import (
	"Todo-list/domain"
	"Todo-list/user"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type UserRepo interface {
	user.UserRepo
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) userRepo {
	return userRepo{
		db: db,
	}

}

func (r *userRepo) Create(user domain.User) (*domain.User, error) {
	query := `
		INSERT INTO users (
			first_name, 
			last_name, 
			email, 
			password, 
			is_done
		)
		VALUES (
		:first_name, 
		:last_name, 
		:email, 
		:password, 
		:is_done)
		RETURNING id;
	`
	var userID int
	rows, err := r.db.NamedQuery(query, user)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if rows.Next() {
		rows.Scan(&userID)
	}
	user.ID = userID
	return &user, nil

}

func (r *userRepo) Find(email, pass string) (*domain.User, error) {
	var user domain.User

	query := `
		SELECT id, 
		first_name, 
		last_name, 
		email, 
		password, 
		is_done
		FROM users
		WHERE email = $1 AND password = $2
		LIMIT 1;
	`

	err := r.db.Get(&user, query, email, pass)

	if err != nil {
		// sql.ErrNoRows means user not found
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
