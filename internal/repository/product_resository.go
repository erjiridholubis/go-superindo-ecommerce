package repository

import (
	"context"

	"github.com/erjiridholubis/go-superindo-product/common"
	"github.com/erjiridholubis/go-superindo-product/internal/model"
)

var (
	// Query Get All Product
	QueryGetAllProduct = `SELECT id, name, description, price, stock, category_id FROM products`
)

// GetAllProduct is a function to get all product data from database
func (pr *postgreRepository) GetAllProduct(ctx context.Context) ([]*model.ProductResponse, error) {
	rows, err := pr.ConnDB.QueryContext(ctx, QueryGetAllProduct)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*model.ProductResponse
	for rows.Next() {
		var product model.ProductResponse
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Stock,
			&product.CategoryID,
		)
		if err != nil {
			return nil, err
		}

		product.Kind = common.KindProduct
		products = append(products, &product)
	}

	return products, nil
}