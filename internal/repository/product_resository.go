package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/erjiridholubis/go-superindo-product/common"
	"github.com/erjiridholubis/go-superindo-product/internal/model"
)

var (
	// Query Get All Product
	QueryGetAllProduct = `SELECT id, name, description, price, stock, category_id FROM products`

	// Query Get Product By ID
	QueryGetProductByID = `SELECT id, name, description, price, stock, category_id FROM products WHERE id = $1`
	
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

// GetProductByID is a function to get product data by ID from database
func (pr *postgreRepository) GetProductByID(ctx context.Context, id string) (*model.ProductResponse, error) {
	var product model.ProductResponse

	err := pr.ConnDB.QueryRowContext(ctx, QueryGetProductByID, id).Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.Stock,
		&product.CategoryID,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	product.Kind = common.KindProduct

	return &product, nil
}