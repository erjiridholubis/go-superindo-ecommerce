package deliveries

import (
	"github.com/erjiridholubis/go-superindo-product/common"
	"github.com/erjiridholubis/go-superindo-product/internal/model"
	categorySrv "github.com/erjiridholubis/go-superindo-product/internal/service"
	"github.com/gofiber/fiber/v2"
)

func NewCategoryHandler(app fiber.Router, categorySrv categorySrv.CategoryService) {
	app.Get("/", getAllCategory(categorySrv))
	app.Get("/:id", getCategoryByID(categorySrv))
	app.Post("/", createCategory(categorySrv))
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
// @Security Authorization
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

// getCategoryByID godoc
// @Summary Get Category By ID
// @Description Get Category By ID
// @Tags Category
// @Accept json
// @Produce json
// @Param id path string true "Category ID"
// @Success 200 {object} model.CategoryResponse
// @Failure 404 {object} common.ApiErrorResponseModel
// @Failure 500 {object} common.ApiErrorResponseModel
// @Security Authorization
// @Router /categories/{id} [get]
func getCategoryByID(categorySrv categorySrv.CategoryService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		categoryId := c.Params("id")
		category, err := categorySrv.GetCategoryByID(c.Context(), categoryId)
		if err != nil {
			if err.Error() == common.ErrNotFound {
				return common.ErrorResponseRest(c, fiber.StatusNotFound, err.Error())
			}
			return common.ErrorResponseRest(c, fiber.StatusInternalServerError, err.Error())
		}

		return common.SuccessResponse(c, fiber.StatusOK, category)
	}
}

// CreateCategory godoc
// @Summary Create Category
// @Description Create Category
// @Tags Category
// @Accept json
// @Produce json
// @Param category body model.CategoryRequest true "Category Body"
// @Success 200 {object} model.CategoryResponse
// @Failure 422 {object} common.ErrorValidationResponseModel
// @Failure 500 {object} common.ApiErrorResponseModel
// @Security Authorization
// @Router /categories [post]
func createCategory(categorySrv categorySrv.CategoryService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var categoryRequest model.CategoryRequest
		err := c.BodyParser(&categoryRequest)
		if err != nil {
			return common.ErrorResponseRest(c, fiber.StatusBadRequest, err.Error())
		}

		if err := common.ValidateStruct(categoryRequest); err != nil {
			return common.ErrorValidationResponse(c, fiber.StatusUnprocessableEntity, common.ValidationFailedMessage, err)
		}

		category, err := categorySrv.CreateCategory(c.Context(), &categoryRequest)
		if err != nil {
			return common.ErrorResponseRest(c, fiber.StatusInternalServerError, err.Error())
		}

		return common.SuccessResponse(c, fiber.StatusOK, category)
	}
}