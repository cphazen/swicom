package database

import (
	"example/swift-comply/database/entities"
	"example/swift-comply/models"
)

func (d *DB) GetCats() ([]models.Cat, error) {
	rows, err := d.db.Query("SELECT * FROM cats")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return entities.Cats(rows)
}
