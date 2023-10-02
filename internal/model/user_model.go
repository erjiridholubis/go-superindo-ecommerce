package model

type User struct {
	ID	   string  `json:"id"`
	Name   string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRequest struct {
	Name   string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}