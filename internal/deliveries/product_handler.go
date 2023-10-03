package deliveries

import (
	"github.com/erjiridholubis/go-superindo-product/common"
	productSrv "github.com/erjiridholubis/go-superindo-product/internal/service"
	"github.com/gofiber/fiber/v2"
)

func NewProductHandler(app fiber.Router, productSrv productSrv.ProductService) {
	app.Get("/", getAllProduct(productSrv))
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
