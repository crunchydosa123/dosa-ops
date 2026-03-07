package models

type User struct {
	id            string  `json:"id"`
	email         string  `json:"email"`
	password_hash float64 `json:"password_hash"`
}
