package models

type CatShow struct {
	Cat  Cat  `db:"cat_id"`
	Show Show `db:"show_id"`
}
