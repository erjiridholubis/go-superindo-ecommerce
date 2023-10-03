package service

import (
	"context"
	"errors"

	"github.com/erjiridholubis/go-superindo-product/common"
	"github.com/erjiridholubis/go-superindo-product/internal/model"
	"github.com/erjiridholubis/go-superindo-product/internal/repository"
)

type productService struct {
	postgreRepo repository.PostgreRepository
}

func NewProductService(postgreRepo repository.PostgreRepository) ProductService {
	return &productService{postgreRepo}
}

type ProductService interface {
	GetAllProduct(ctx context.Context) (resp *model.ProductList, err error)
	GetProductByID(ctx context.Context, id string) (resp *model.ProductResponse, err error)
	CreateProduct(ctx context.Context, product *model.ProductRequest) (resp *model.ProductResponse, err error)
}

func (ps *productService) GetAllProduct(ctx context.Context) (resp *model.ProductList, err error) {
	products, err := ps.postgreRepo.GetAllProduct(ctx)
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
	resp, err = ps.postgreRepo.GetProductByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, errors.New(common.ErrNotFound)
	}

	return resp, nil
}

func (ps *productService) CreateProduct(ctx context.Context, product *model.ProductRequest) (*model.ProductResponse, error) {
	data := &model.Product{
		ID: common.GenerateUUID(),
		Name: product.Name,
		Description: product.Description,
		Price: product.Price,
		Stock: product.Stock,
		CategoryID: product.CategoryID,
	}

	_, err := ps.postgreRepo.CreateProduct(ctx, data)
	if err != nil {
		return nil, err
	}

	resp := &model.ProductResponse{
		Kind: common.KindProduct,
		ID: data.ID,
		Name: data.Name,
		Description: data.Description,
		Price: data.Price,
		Stock: data.Stock,
		CategoryID: data.CategoryID,
	}

	return resp, nil
}