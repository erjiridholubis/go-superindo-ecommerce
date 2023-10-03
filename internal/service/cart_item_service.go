package service

import (
	"context"
	"errors"

	"github.com/erjiridholubis/go-superindo-product/common"
	"github.com/erjiridholubis/go-superindo-product/internal/model"
	"github.com/erjiridholubis/go-superindo-product/internal/repository"
)

type cartItemService struct {
	postgreRepo repository.PostgreRepository
}

func NewCartItemService(postgreRepo repository.PostgreRepository) CartItemService {
	return &cartItemService{postgreRepo}
}

type CartItemService interface {
	CreateCartItem(ctx context.Context, cartItem *model.CartItemRequest, userID string) (resp *model.CartItemResponse, err error)
	GetCartItemByUserID(ctx context.Context, userID string) (resp *model.CartItemList, err error)
}

func (cs *cartItemService) CreateCartItem(ctx context.Context, cartItem *model.CartItemRequest, userID string) (resp *model.CartItemResponse, err error) {

	data := &model.CartItem{
		ID: common.GenerateUUID(),
		UserID: userID,
		ProductID: cartItem.ProductID,
		Quantity: cartItem.Quantity,
	}

	// Check if cart item exist
	cartItemExist, err := cs.postgreRepo.GetCartItemByUserIDAndProductID(ctx, userID, cartItem.ProductID)
	if err != nil {
		return nil, err
	}

	if cartItemExist != nil {
		// Update cart item if exist
		data.ID = cartItemExist.ID
		data.Quantity = cartItemExist.Quantity + cartItem.Quantity

		resp, err = cs.postgreRepo.UpdateCartItem(ctx, data)
		if err != nil {
			return nil, err
		}

		return resp, nil
	}

	// Create cart item if not exist
	resp, err = cs.postgreRepo.CreateCartItem(ctx, data)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (cs *cartItemService) GetCartItemByUserID(ctx context.Context, userID string) (resp *model.CartItemList, err error) {
	cartItems, err := cs.postgreRepo.GetCartItemByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if len(cartItems) == 0 {
		return nil, errors.New(common.ErrNotFound)
	}

	resp = &model.CartItemList{
		Kind: common.KindCollection,
		CartItems: cartItems,
	}

	return resp, nil
}