package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/erjiridholubis/go-superindo-product/common"
	"github.com/erjiridholubis/go-superindo-product/internal/model"
)

var (
	// Query Create Cart Item
	QueryCreateCartItem = `INSERT INTO cart_items (id, user_id, product_id, quantity) VALUES ($1, $2, $3, $4) RETURNING id, user_id, product_id, quantity`

	// Query Get Cart Item By User ID and Product ID
	QueryGetCartItemByUserIDAndProductID = `SELECT id, user_id, product_id, quantity FROM cart_items WHERE user_id = $1 AND product_id = $2`

	// Query Update Cart Item
	QueryUpdateCartItem = `UPDATE cart_items SET quantity = $1 WHERE id = $2 RETURNING id, user_id, product_id, quantity`

	// Query Get Cart Item By User ID
	QueryGetCartItemByUserID = `SELECT id, user_id, product_id, quantity FROM cart_items WHERE user_id = $1`
)

func (pr *postgreRepository) CreateCartItem(ctx context.Context, req *model.CartItem) (*model.CartItemResponse, error) {
	var cartItem model.CartItemResponse

	err := pr.ConnDB.QueryRowContext(ctx, QueryCreateCartItem, req.ID, req.UserID, req.ProductID, req.Quantity).Scan(
		&cartItem.ID,
		&cartItem.UserID,
		&cartItem.ProductID,
		&cartItem.Quantity,
	)
	if err != nil {
		return nil, err
	}

	cartItem.Kind = common.KindCartItem

	return &cartItem, nil
}

func (pr *postgreRepository) GetCartItemByUserIDAndProductID(ctx context.Context, userID, productID string) (*model.CartItemResponse, error) {
	var cartItem model.CartItemResponse

	err := pr.ConnDB.QueryRowContext(ctx, QueryGetCartItemByUserIDAndProductID, userID, productID).Scan(
		&cartItem.ID,
		&cartItem.UserID,
		&cartItem.ProductID,
		&cartItem.Quantity,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	cartItem.Kind = common.KindCartItem

	return &cartItem, nil
}

func (pr *postgreRepository) UpdateCartItem(ctx context.Context, req *model.CartItem) (*model.CartItemResponse, error) {
	var cartItem model.CartItemResponse

	err := pr.ConnDB.QueryRowContext(ctx, QueryUpdateCartItem, req.Quantity, req.ID).Scan(
		&cartItem.ID,
		&cartItem.UserID,
		&cartItem.ProductID,
		&cartItem.Quantity,
	)
	if err != nil {
		return nil, err
	}

	cartItem.Kind = common.KindCartItem

	return &cartItem, nil
}

func (pr *postgreRepository) GetCartItemByUserID(ctx context.Context, userID string) ([]*model.CartItemResponse, error) {
	var cartItems []*model.CartItemResponse

	rows, err := pr.ConnDB.QueryContext(ctx, QueryGetCartItemByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var cartItem model.CartItemResponse
		err := rows.Scan(
			&cartItem.ID,
			&cartItem.UserID,
			&cartItem.ProductID,
			&cartItem.Quantity,
		)
		if err != nil {
			return nil, err
		}

		cartItem.Kind = common.KindCartItem
		cartItems = append(cartItems, &cartItem)
	}

	return cartItems, nil
}