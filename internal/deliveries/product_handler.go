package deliveries

import (
	"github.com/erjiridholubis/go-superindo-product/common"
	productSrv "github.com/erjiridholubis/go-superindo-product/internal/service"
	"github.com/gofiber/fiber/v2"
)

func NewProductHandler(app fiber.Router, productSrv productSrv.ProductService) {
	app.Get("/", getAllProduct(productSrv))
	app.Get("/:id", getProductByID(productSrv))
}

// GetAllProduct godoc
// @Summary Get All Product
// @Description Get All Product
// @Tags Product
// @Accept json
// @Produce json
// @Success 200 {object} model.ProductList
// @Failure 404 {object} common.ApiErrorResponseModel
// @Failure 500 {object} common.ApiErrorResponseModel
// @Router /products [get]
func getAllProduct(productSrv productSrv.ProductService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		product, err := productSrv.GetAllProduct(c.Context())
		if err != nil {
			if err.Error() == common.ErrNotFound {
				return common.ErrorResponseRest(c, fiber.StatusNotFound, err.Error())
			}
			return common.ErrorResponseRest(c, fiber.StatusInternalServerError, err.Error())
		}

		return common.SuccessResponse(c, fiber.StatusOK, product)
	}
}

// GetProductByID godoc
// @Summary Get Product By ID
// @Description Get Product By ID
// @Tags Product
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} model.ProductResponse
// @Failure 404 {object} common.ApiErrorResponseModel
// @Failure 500 {object} common.ApiErrorResponseModel
// @Router /products/{id} [get]
func getProductByID(productSrv productSrv.ProductService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		productId := c.Params("id")
		product, err := productSrv.GetProductByID(c.Context(), productId)
		if err != nil {
			if err.Error() == common.ErrNotFound {
				return common.ErrorResponseRest(c, fiber.StatusNotFound, err.Error())
			}
			return common.ErrorResponseRest(c, fiber.StatusInternalServerError, err.Error())
		}

		return common.SuccessResponse(c, fiber.StatusOK, product)
	}
}