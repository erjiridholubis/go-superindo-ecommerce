package model

type Category struct {
	ID          string  `json:"id"`
	Name        string `json:"name"`
}

type CategoryRequest struct {
	Name        string `json:"name"`
}

type CategoryResponse struct {
	Kind		string `json:"kind"`
	ID		 	string  `json:"id"`
	Name		string `json:"name"`
}

type CategoryList struct {
	Kind		string `json:"kind"`
	Categories	[]*CategoryResponse `json:"categories"`
}