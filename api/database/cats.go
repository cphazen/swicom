package database

import (
	"database/sql"
	"example/swift-comply/models"
	"fmt"
)

func parseCatRow(r *sql.Rows) (models.Cat, error) {
	var cat models.Cat
	var breed models.Breed
	if err := r.Scan(&cat.Id, &cat.Name, &cat.DateOfBirth, &breed.Id); err != nil {
		return cat, err
	}
	cat.Breed = &breed
	return cat, nil
}

func Cat(r *sql.Rows) (models.Cat, error) {
	ok := r.Next()
	if !ok {
		return models.Cat{}, nil
	}
	return parseCatRow(r)
}

func Cats(r *sql.Rows) ([]models.Cat, error) {
	var cats []models.Cat
	for r.Next() {
		cat, err := parseCatRow(r)
		if err != nil {
			return cats, err
		}
		cats = append(cats, cat)
	}
	return cats, nil
}

func (d *DB) GetCats() ([]models.Cat, error) {
	rows, err := d.db.Query("SELECT * FROM cats")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return Cats(rows)
}

func parseCatDetails(r *sql.Rows) (models.Cat, error) {
	var cat models.Cat
	var breed models.Breed
	var breedGroup models.BreedGroup
	if err := r.Scan(&cat.Id, &cat.Name, &cat.DateOfBirth, &breed.Id, &breed.Name, &breedGroup.Id, &breedGroup.Name); err != nil {
		return cat, err
	}
	breed.BreedGroup = &breedGroup
	cat.Breed = &breed
	return cat, nil
}

func CatDetails(r *sql.Rows) (models.Cat, error) {
	ok := r.Next()
	if !ok {
		return models.Cat{}, nil
	}
	return parseCatDetails(r)
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
	return CatDetails(row)
}
