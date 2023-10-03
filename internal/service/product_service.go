package service

import (
	"context"
	"errors"

	"github.com/erjiridholubis/go-superindo-product/common"
	"github.com/erjiridholubis/go-superindo-product/internal/model"
	"github.com/erjiridholubis/go-superindo-product/internal/repository"
)

type productService struct {
	productRepo repository.PostgreRepository
}

func NewProductService(productRepo repository.PostgreRepository) ProductService {
	return &productService{productRepo}
}

type ProductService interface {
	GetAllProduct(ctx context.Context) (resp *model.ProductList, err error)
	GetProductByID(ctx context.Context, id string) (resp *model.ProductResponse, err error)
}

func (ps *productService) GetAllProduct(ctx context.Context) (resp *model.ProductList, err error) {
	products, err := ps.productRepo.GetAllProduct(ctx)
	if err != nil {
		return nil, err
	}

	if len(products) == 0 {
		return nil, errors.New(common.ErrNotFound)
	}

	resp = &model.ProductList{
		Kind:     common.KindCollection,
		Products: products,
	}

	return resp, nil
}

func (ps *productService) GetProductByID(ctx context.Context, id string) (resp *model.ProductResponse, err error) {
	resp, err = ps.productRepo.GetProductByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, errors.New(common.ErrNotFound)
	}

	return resp, nil
}