// middleware/jwt.go
package middleware

import (
	"strings"
	"time"

	"github.com/erjiridholubis/go-superindo-product/common"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

const (
	secretKey  = "API-SUPERINDO-PRODUCT"
	expiration = time.Hour * 1 // Token expiration time
)

// GenerateToken akan membuat token JWT untuk pengguna yang telah login.
func GenerateToken(userID string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject: userID,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiration)),
	})

	tokenString, err := claims.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func extractBearerToken(authorizationHeader string) (string, bool) {
    if strings.HasPrefix(authorizationHeader, "Bearer ") {
        return strings.TrimPrefix(authorizationHeader, "Bearer "), true
    }
    return "", false
}

// JWTMiddleware adalah middleware untuk mengamankan rute dengan JWT.
func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var token string

		reqHeaders := c.GetReqHeaders()

		 if getToken, ok := extractBearerToken(reqHeaders["Authorization"]); ok {
			token = getToken
		 } else {
			return common.ErrorResponseRest(c, fiber.StatusForbidden, "Invalid auth token")
		}

		claims := jwt.RegisteredClaims{}
		_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

		if err != nil {
			return common.ErrorResponseRest(c, fiber.StatusUnauthorized, err.Error())
		}

		return c.Next()
	}
}
