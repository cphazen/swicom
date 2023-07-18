package database

import (
	"example/swift-comply/database/mapping"
	"example/swift-comply/models"
	"fmt"
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
	row, err := d.db.Query(`SELECT 
			c.id, 
			c.name, 
			c.date_of_birth, 
			c.breed_id, 
			b.name, 
			b.breed_group_id, 
			bg.name 
		FROM cats c
		LEFT JOIN breeds b ON c.breed_id = b.id
		LEFT JOIN breed_groups bg ON b.breed_group_id = bg.id
		WHERE c.id = $1`, id)
	if err != nil {
		fmt.Println("err", err)
		return models.Cat{}, err
	}
	fmt.Println("row", row)
	defer row.Close()
	return mapping.CatDetails(row)
}
