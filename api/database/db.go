package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type DB struct {
	db *sql.DB
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
