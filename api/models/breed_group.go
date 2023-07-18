package models

type BreedGroup struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}

type BreedGroupJson struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
