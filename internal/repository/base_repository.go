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
	CreateCategory(ctx context.Context, category *model.Category) (string, error)

	// Product
	GetAllProduct(ctx context.Context) ([]*model.ProductResponse, error)
	GetProductByID(ctx context.Context, id string) (*model.ProductResponse, error)
	CreateProduct(ctx context.Context, product *model.Product) (string, error)
	GetProductByCategoryID(ctx context.Context, id string) ([]*model.ProductResponse, error)

	// User
	GetUserByUsername(ctx context.Context, username string) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) (string, error)

	// Cart Item
	CreateCartItem(ctx context.Context, req *model.CartItem) (*model.CartItemResponse, error)
	GetCartItemByUserIDAndProductID(ctx context.Context, userID, productID string) (*model.CartItemResponse, error)
	UpdateCartItem(ctx context.Context, req *model.CartItem) (*model.CartItemResponse, error)
	GetCartItemByUserID(ctx context.Context, userID string) ([]*model.CartItemResponse, error)
}