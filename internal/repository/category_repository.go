package repository

import (
	"context"

	"github.com/erjiridholubis/go-superindo-product/common"
	"github.com/erjiridholubis/go-superindo-product/internal/model"
)

var (
	// Query Get All Category
	QueryGetAllCategory = `SELECT id, name FROM categories`
)

func (pr *postgreRepository) GetAllCategory(ctx context.Context) (resp []*model.CategoryResponse, err error) {
	rows, err := pr.ConnDB.QueryContext(ctx, QueryGetAllCategory)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*model.CategoryResponse
	for rows.Next() {
		var category model.Category
		err = rows.Scan(
			&category.ID,
			&category.Name,
		)
		if err != nil {
			return nil, err
		}

		categories = append(categories, &model.CategoryResponse{
			Kind: common.KindCategory,
			ID: category.ID,
			Category: &category,
		})
	}

	return categories, nil
}