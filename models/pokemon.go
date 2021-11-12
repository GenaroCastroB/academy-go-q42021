package models

type Pokemon struct {
	Id   int    `json:"id" gorm primary_key`
	Name string `json:"name"`
}
