package model

type CartItem struct {
    ID         string `json:"id"`
    UserID     string `json:"user_id"`
    ProductID  string `json:"product_id"`
    Quantity   uint64  `json:"quantity"`
}

type CartItemRequest struct {
    ProductID  string `json:"product_id" validate:"required" `
    Quantity   uint64  `json:"quantity" validate:"required" `
}

type CartItemResponse struct {
    Kind       string `json:"kind"`
    ID         string `json:"id"`
    UserID     string `json:"user_id"`
    ProductID  string `json:"product_id"`
    Quantity   uint64  `json:"quantity"`
}

type CartItemList struct {
    Kind       string `json:"kind"`
    CartItems  []*CartItemResponse `json:"cart_items"`
}