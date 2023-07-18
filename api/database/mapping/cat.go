package mapping

import (
	"database/sql"
	"example/swift-comply/models"
)

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

func CatJson(c models.Cat) models.CatJson {
	result := models.CatJson{
		Id:          c.Id,
		Name:        c.Name,
		DateOfBirth: c.DateOfBirth,
	}
	if c.Breed != nil {
		breed := BreedJson(*c.Breed)
		result.Breed = &breed
	}
	return result
}
