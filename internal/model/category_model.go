package model

type Category struct {
	ID          string  `json:"id"`
	Name        string `json:"name"`
}

type CategoryRequest struct {
	Name        string `json:"name"`
}