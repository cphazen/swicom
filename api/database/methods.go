package database

import (
	"example/swift-comply/database/mapping"
	"example/swift-comply/models"
)

func (d *DB) GetCats() ([]models.Cat, error) {
	rows, err := d.db.Query("SELECT * FROM cats")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return mapping.Cats(rows)
}

func (d *DB) GetCatById(id int) (models.Cat, error) {
	row, err := d.db.Query("SELECT * FROM cats WHERE id = $1", id)
	if err != nil {
		return models.Cat{}, err
	}
	defer row.Close()
	return mapping.Cat(row)
}
