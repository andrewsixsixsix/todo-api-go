package storage

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var storage *sql.DB

func Init(datasource string) error {
	var err error
	storage, err = sql.Open("pgx", datasource)
	if err != nil {
		return err
	}

	return nil
}

func Storage() *sql.DB {
	return storage
}
