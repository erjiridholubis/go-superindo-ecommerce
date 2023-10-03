package deliveries

import (
	"github.com/erjiridholubis/go-superindo-product/common"
	"github.com/erjiridholubis/go-superindo-product/internal/model"
	authSrv "github.com/erjiridholubis/go-superindo-product/internal/service"
	"github.com/gofiber/fiber/v2"
)

func NewAuthHandler(app fiber.Router, authSrv authSrv.AuthService) {
	app.Post("/login", login(authSrv))
}

// login godoc
// @Summary Login
// @Description Login
// @Tags Auth
// @Accept json
// @Produce json
// @Param auth body model.AuthRequest true "Auth" 
// @Success 200 {object} model.AuthResponse
// @Failure 400 {object} common.ApiErrorResponseModel
// @Failure 422 {object} common.ApiErrorResponseModel
// @Failure 404 {object} common.ApiErrorResponseModel
// @Failure 500 {object} common.ApiErrorResponseModel
// @Router /auth/login [post]
func login(authSrv authSrv.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var payload model.AuthRequest
		if err := c.BodyParser(&payload); err != nil {
			return common.ErrorResponseRest(c, fiber.StatusBadRequest, err.Error())
		}

		if err := common.ValidateStruct(payload); err != nil {
			return common.ErrorValidationResponse(c, fiber.StatusUnprocessableEntity, common.ValidationFailedMessage, err)
		}

		resp, err := authSrv.Login(c.Context(), payload)
		if err != nil {
			if err.Error() == common.ErrNotFound {
				return common.ErrorResponseRest(c, fiber.StatusNotFound, err.Error())
			}
			return common.ErrorResponseRest(c, fiber.StatusInternalServerError, err.Error())
		}

		return common.SuccessResponse(c, fiber.StatusOK, resp)
	}
}
