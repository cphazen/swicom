package models

type Breed struct {
	Id         int         `db:"id"`
	Name       string      `db:"name"`
	BreedGroup *BreedGroup `db:"breed_group_id"`
}

type BreedJson struct {
	Id         int             `json:"id"`
	Name       string          `json:"name"`
	BreedGroup *BreedGroupJson `json:"breed_group,omitempty"`
}
