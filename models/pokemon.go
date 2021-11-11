package models

type Pokemon struct {
	Id   string `json:"id" gorm primary_key`
	Name string `json:"name"`
}
