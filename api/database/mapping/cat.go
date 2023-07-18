package mapping

import (
	"database/sql"
	"example/swift-comply/models"
)

func parseCatRow(r *sql.Rows) (models.Cat, error) {
	var cat models.Cat
	if err := r.Scan(&cat.Id, &cat.Name, &cat.DateOfBirth, &cat.BreedId); err != nil {
		return cat, err
	}
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
	return models.CatJson{
		Id:          c.Id,
		Name:        c.Name,
		DateOfBirth: c.DateOfBirth,
	}
}
