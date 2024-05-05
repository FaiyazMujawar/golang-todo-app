package models

type User struct {
	Id        string `json:"_id"`
	Firstname string
	Lastname  string
	Email     string
	Password  string `json:"-"`
}
