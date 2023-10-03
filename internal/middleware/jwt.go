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

// GenerateToen is function to generate JWT token for logged in user.
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

// JWTMiddleware is a middleware to secure routes with JWT.
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
		parsedToken, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

		if err != nil {
			return common.ErrorResponseRest(c, fiber.StatusUnauthorized, err.Error())
		}

		if parsedToken.Valid {
			c.Locals("userID", claims.Subject)
		}

		return c.Next()
	}
}
