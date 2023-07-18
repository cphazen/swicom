package models

type Cat struct {
	Id          int    `db:"id"`
	Name        string `db:"name"`
	DateOfBirth string `db:"date_of_birth"`
	Breed       *Breed `db:"breed_id"`
}

type CatJson struct {
	Id          int        `json:"id"`
	Name        string     `json:"name"`
	DateOfBirth string     `json:"date_of_birth"`
	Breed       *BreedJson `json:"breed,omitempty"`
}
