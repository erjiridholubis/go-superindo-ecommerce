package model

type Product struct {
	ID          string  `json:"id"`
	CategoryID	string `json:"category_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       uint64  `json:"price"`
	Stock       uint64  `json:"stock"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type ProductRequest struct {
	CategoryID	string `json:"category_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       uint64  `json:"price"`
	Stock       uint64  `json:"stock"`
}

type ProductResponse struct {
	Kind		string `json:"kind"`
	ID          string  `json:"id"`
	CategoryID	string `json:"category_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       uint64  `json:"price"`
	Stock       uint64  `json:"stock"`
}

type ProductList struct {
	Kind		string `json:"kind"`
	CategoryID	string `json:"category_id"`
	Products	[]*ProductResponse `json:"products"`
}