package models

type User struct {
	ID    string `json:"id"`
	name  string `json:"name"`
	email string `json:"email"`
}
