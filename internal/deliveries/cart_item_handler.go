package deliveries

import (
	"github.com/erjiridholubis/go-superindo-product/common"
	"github.com/erjiridholubis/go-superindo-product/internal/model"
	cartItemSrv "github.com/erjiridholubis/go-superindo-product/internal/service"
	"github.com/gofiber/fiber/v2"
)

func NewCartItemHandler(app fiber.Router, cartItemSrv cartItemSrv.CartItemService) {
	app.Post("/", createCartItem(cartItemSrv))
	app.Get("/:user_id", getCartItemByUserID(cartItemSrv))
}

// createCartItem godoc
// @Summary Create Cart Item
// @Description Create Cart Item
// @Tags Cart Item
// @Accept json
// @Produce json
// @Param cartItem body model.CartItemRequest true "Cart Item"
// @Success 200 {object} model.CartItemResponse
// @Failure 400 {object} common.ErrorValidationResponseModel
// @Failure 422 {object} common.ErrorValidationResponseModel
// @Failure 500 {object} common.ApiErrorResponseModel
// @Security Authorization
// @Router /cart-items [post]
func createCartItem(cartItemSrv cartItemSrv.CartItemService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("userID").(string)

		var cartItemRequest model.CartItemRequest
		err := c.BodyParser(&cartItemRequest)
		if err != nil {
			return common.ErrorResponseRest(c, fiber.StatusBadRequest, err.Error())
		}

		if err := common.ValidateStruct(cartItemRequest); err != nil {
			return common.ErrorValidationResponse(c, fiber.StatusUnprocessableEntity, common.ValidationFailedMessage, err)
		}

		resp, err := cartItemSrv.CreateCartItem(c.Context(), &cartItemRequest, userID)
		if err != nil {
			return common.ErrorResponseRest(c, fiber.StatusInternalServerError, err.Error())
		}

		return common.SuccessResponse(c, fiber.StatusOK, resp)
	}
}

// getCartItemByUserID godoc
// @Summary Get Cart Item By User ID
// @Description Get Cart Item By User ID
// @Tags Cart Item
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Success 200 {object} model.CartItemList
// @Failure 404 {object} common.ApiErrorResponseModel
// @Failure 500 {object} common.ApiErrorResponseModel
// @Security Authorization
// @Router /cart-items/{user_id} [get]
func getCartItemByUserID(cartItemSrv cartItemSrv.CartItemService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Params("user_id")

		resp, err := cartItemSrv.GetCartItemByUserID(c.Context(), userID)
		if err != nil {
			if err.Error() == common.ErrNotFound {
				return common.ErrorResponseRest(c, fiber.StatusNotFound, err.Error())
			}
			return common.ErrorResponseRest(c, fiber.StatusInternalServerError, err.Error())
		}

		return common.SuccessResponse(c, fiber.StatusOK, resp)
	}
}