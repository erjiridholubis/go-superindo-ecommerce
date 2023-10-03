package deliveries

import (
	"github.com/erjiridholubis/go-superindo-product/common"
	userSrv "github.com/erjiridholubis/go-superindo-product/internal/service"
	"github.com/gofiber/fiber/v2"
)

func NewUserHandler(app fiber.Router, userSrv userSrv.UserService) {
	app.Get("/profile", getUserProfile(userSrv))
}

// getUserProfile godoc
// @Summary Get User Profile
// @Description Get User Profile by User ID from JWT
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} model.User
// @Failure 404 {object} common.ApiErrorResponseModel
// @Failure 500 {object} common.ApiErrorResponseModel
// @Security Authorization
// @Router /users/profile [get]
func getUserProfile(userSrv userSrv.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("userID").(string)

		resp, err := userSrv.GetUserByID(c.Context(), userID)
		if err != nil {
			if err.Error() == common.ErrNotFound {
				return common.ErrorResponseRest(c, fiber.StatusNotFound, err.Error())
			}
			return common.ErrorResponseRest(c, fiber.StatusInternalServerError, err.Error())
		}

		return common.SuccessResponse(c, fiber.StatusOK, resp)
	}
}

