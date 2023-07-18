package models

type Organization struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}

type OrganizationJson struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
