package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConectionString() string {

	return "user=postgres password=admin12345 host=localhost port=5432 dbname=todo sslmode=disable"

}

func NewConnetion() (*sqlx.DB, error) {

	dbSource := GetConectionString()
	dbcon, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return dbcon, nil

}
