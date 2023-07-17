package models

type Cat struct {
	Id          string `db:"id"`
	Name        string `db:"name"`
	DateOfBirth string `db:"date_of_birth"`
	BreedId     string `db:"breed_id"`
}

type CatJson struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	DateOfBirth string `json:"date_of_birth"`
}
