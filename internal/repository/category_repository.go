package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/erjiridholubis/go-superindo-product/common"
	"github.com/erjiridholubis/go-superindo-product/internal/model"
)

var (
	// Query Get All Category
	QueryGetAllCategory = `SELECT id, name FROM categories`

	// Query Get Category By ID
	QueryGetCategoryByID = `SELECT id, name FROM categories WHERE id = $1`

	// Query Create Category
	QueryCreateCategory = `INSERT INTO categories (id, name) VALUES ($1, $2) RETURNING id`
)

func (pr *postgreRepository) GetAllCategory(ctx context.Context) (resp []*model.CategoryResponse, err error) {
	rows, err := pr.ConnDB.QueryContext(ctx, QueryGetAllCategory)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var category model.CategoryResponse
		err := rows.Scan(
			&category.ID,
			&category.Name,
		)
		if err != nil {
			return nil, err
		}

		category.Kind = common.KindCategory
		resp = append(resp, &category)
	}

	return resp, nil
}

func (pr *postgreRepository) GetCategoryByID(ctx context.Context, id string) (*model.CategoryResponse, error) {
	var category model.CategoryResponse

	err := pr.ConnDB.QueryRowContext(ctx, QueryGetCategoryByID, id).Scan(
		&category.ID,
		&category.Name,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	category.Kind = common.KindCategory
	return &category, nil
}

func (pr *postgreRepository) CreateCategory(ctx context.Context, category *model.Category) (productId string, err error) {
	err = pr.ConnDB.QueryRowContext(ctx, QueryCreateCategory, category.ID, category.Name).Scan(&productId)
	if err != nil {
		return "", err
	}

	return productId, nil
}