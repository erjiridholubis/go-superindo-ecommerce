package model

type User struct {
	ID	   string  `json:"id"`
	Name   string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRequest struct {
	Name   string `json:"name" validate:"required" `
	Username string `json:"username" validate:"required" `
	Password string `json:"password" validate:"required" `
}

type AuthRequest struct {
	Username string `json:"username" validate:"required" `
	Password string `json:"password" validate:"required" `
}

type AuthResponse struct {
	Kind		string `json:"kind"`
	Token		string `json:"token"`
}