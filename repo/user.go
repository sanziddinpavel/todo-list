package repo

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID        int    `db:"id" json:"id"`
	FirstName string `db:"first_name" json:"first_name"`
	LastName  string `db:"last_name" json:"last_name"`
	Email     string `db:"email" json:"email"`
	Password  string `db:"password" json:"password"`
	IsDone    bool   `db:"is_done" json:"is_done"`
}

type UserRepo interface {
	Create(user User) (*User, error)
	Find(email, password string) (*User, error)
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) userRepo {
	return userRepo{
		db: db,
	}

}

func (r *userRepo) Create(user User) (*User, error) {
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

func (r *userRepo) Find(email, pass string) (*User, error) {
	var user User

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
