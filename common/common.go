package common

import "github.com/google/uuid"

var (
	KindCollection = "collection"
	KindUser       = "user"
	KindProduct    = "product"
	KindCartItem   = "cart-item"
	KindCategory   = "category"

	MessageCreateSuccess    = "Successfully adding data"
	ValidationFailedMessage = "Validation Failed"
	ErrNotFound             = "Data not found"
)

func GenerateUUID() string {
	return uuid.New().String()
}