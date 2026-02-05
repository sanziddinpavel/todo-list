package db

import (
	"Todo-list/config"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString(cnf *config.DBConfig) string {
	connString := fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s",
		cnf.User,
		cnf.Password,
		cnf.Host,
		cnf.Port,
		cnf.Name,
	)
	if !cnf.EnableSSLMode {
		connString += " sslmode=disable"
	}
	return connString
}

func NewConnetion(cnf *config.DBConfig) (*sqlx.DB, error) {

	dbSource := GetConnectionString(cnf)
	dbcon, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return dbcon, nil

}
