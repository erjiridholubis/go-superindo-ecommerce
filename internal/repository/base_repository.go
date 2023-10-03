package repository

import (
	"context"
	"database/sql"

	"github.com/erjiridholubis/go-superindo-product/internal/model"
)

type postgreRepository struct {
	ConnDB *sql.DB
}

func NewPostgreRepository(connDB *sql.DB) PostgreRepository {
	return &postgreRepository{connDB}
}

type PostgreRepository interface {
	// Category
	GetAllCategory(ctx context.Context) (resp []*model.CategoryResponse, err error)
	GetCategoryByID(ctx context.Context, id string) (*model.CategoryResponse, error)

	// Product
	GetAllProduct(ctx context.Context) ([]*model.ProductResponse, error)
}