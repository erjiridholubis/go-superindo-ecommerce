package deliveries

import (
	"github.com/erjiridholubis/go-superindo-product/common"
	categorySrv "github.com/erjiridholubis/go-superindo-product/internal/service"
	"github.com/gofiber/fiber/v2"
)

func NewCategoryHandler(app fiber.Router, categorySrv categorySrv.CategoryService) {
	app.Get("/", getAllCategory(categorySrv))
}

// getAllCategory godoc
// @Summary Get All Category
// @Description Get All Category
// @Tags Category
// @Accept json
// @Produce json
// @Success 200 {object} model.CategoryList
// @Failure 404 {object} common.ApiErrorResponseModel
// @Failure 500 {object} common.ApiErrorResponseModel
// @Router /categories [get]
func getAllCategory(categorySrv categorySrv.CategoryService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		category, err := categorySrv.GetAllCategory(c.Context())
		if err != nil {
			if err.Error() == common.ErrNotFound {
				return common.ErrorResponseRest(c, fiber.StatusNotFound, err.Error())
			}
			return common.ErrorResponseRest(c, fiber.StatusInternalServerError, err.Error())
		}

		return common.SuccessResponse(c, fiber.StatusOK, category)
	}
}
