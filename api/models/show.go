package models

type Show struct {
	Id           int          `db:"id"`
	Name         string       `db:"name"`
	Organization Organization `db:"organization_id"`
}

type ShowJson struct {
	Id           int              `json:"id"`
	Name         string           `json:"name"`
	Organization OrganizationJson `json:"organization"`
}
