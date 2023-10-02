package common

import "github.com/gofiber/fiber/v2"

type ErrorResponseModel struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Errors  map[string]interface{} `json:"errors,omitempty"`
}

type ApiErrorResponseModel struct {
	Error ErrorResponseModel `json:"error"`
}

func SuccessResponse(c *fiber.Ctx, status int, data interface{}) error {
	return c.Status(status).JSON(data)
}

func ErrorResponseRest(c *fiber.Ctx, code int, errStr string) error {
	return c.Status(code).JSON(ApiErrorResponseModel{
		Error: ErrorResponseModel{
			Code:    code,
			Message: errStr,
		},
	})
}
