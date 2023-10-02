package model

type Product struct {
	ID          string  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       uint64  `json:"price"`
	Stock       uint64  `json:"stock"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type ProductRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       uint64  `json:"price"`
	Stock       uint64  `json:"stock"`
}

type ProductResponse struct {
	ID          string  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       uint64  `json:"price"`
	Stock       uint64  `json:"stock"`
}