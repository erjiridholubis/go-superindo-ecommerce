package common

import (
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	KindCollection = "collection"
	KindUser       = "user"
	KindProduct    = "product"
	KindCartItem   = "cart-item"
	KindCategory   = "category"
	KindAuthorization = "authorization"

	MessageCreateSuccess    = "Successfully adding data"
	ValidationFailedMessage = "Validation Failed"
	
	ErrUsernameAlreadyExist = "Username already exist"
	ErrNotFound             = "Data not found"
	ErrInvalidPassword		= "Invalid password"
)

func GenerateUUID() string {
	return uuid.New().String()
}

func HashAndSalt(pwd []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	
	return string(hash), err
}

func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
    byteHash := []byte(hashedPwd)
    err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
    if err != nil {
        fmt.Println(err)
        return false
    }
    
    return true
}