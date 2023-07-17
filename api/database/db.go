package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type DB struct {
	db *sql.DB
}

type DBEntity[T any] interface {
	getEntity(*sql.Rows) (T, error)
	getEntities(*sql.Rows) ([]T, error)
}

func (d *DB) Open() error {
	pg, err := sql.Open("postgres", pgConnectionString)
	if err != nil {
		return err
	}
	d.db = pg
	return nil
}

func (d *DB) Close() error {
	return d.db.Close()
}
