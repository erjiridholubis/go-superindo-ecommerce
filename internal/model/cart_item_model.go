package model

type CartItem struct {
    ID         string `json:"id"`
    UserID     string `json:"user_id"`
    ProductID  string `json:"product_id"`
    Quantity   uint64  `json:"quantity"`
}

type CartItemRequest struct {
    UserID     string `json:"user_id"`
    ProductID  string `json:"product_id"`
    Quantity   uint64  `json:"quantity"`
}
