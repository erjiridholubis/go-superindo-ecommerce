package service

import (
	"context"
	"errors"

	"github.com/erjiridholubis/go-superindo-product/common"
	"github.com/erjiridholubis/go-superindo-product/internal/model"
	"github.com/erjiridholubis/go-superindo-product/internal/repository"
)

type categoryService struct {
	postgreRepo repository.PostgreRepository
}

func NewCategoryService(postgreRepo repository.PostgreRepository) CategoryService {
	return &categoryService{postgreRepo}
}

type CategoryService interface {
	GetAllCategory(ctx context.Context) (resp *model.CategoryList, err error)
	GetCategoryByID(ctx context.Context, id string) (resp *model.CategoryResponse, err error)
	CreateCategory(ctx context.Context, category *model.CategoryRequest) (resp *model.CategoryResponse, err error)
}

func (cs *categoryService) GetAllCategory(ctx context.Context) (resp *model.CategoryList, err error) {
	categories, err := cs.postgreRepo.GetAllCategory(ctx)
	if err != nil {
		return nil, err
	}

	if len(categories) == 0 {
		return nil, errors.New(common.ErrNotFound)
	}

	resp = &model.CategoryList{
		Kind: common.KindCollection,
		Categories: categories,
	}

	return resp, nil
}

func (cs *categoryService) GetCategoryByID(ctx context.Context, id string) (resp *model.CategoryResponse, err error) {
	resp, err = cs.postgreRepo.GetCategoryByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, errors.New(common.ErrNotFound)
	}

	return resp, nil
}

func (cs *categoryService) CreateCategory(ctx context.Context, category *model.CategoryRequest) (*model.CategoryResponse, error) {
	data := &model.Category{
		ID: common.GenerateUUID(),
		Name: category.Name,
	}

	_, err := cs.postgreRepo.CreateCategory(ctx, data)
	if err != nil {
		return nil, err
	}

	resp := &model.CategoryResponse{
		Kind: common.KindCategory,
		ID: data.ID,
		Name: data.Name,
	}

	return resp, nil
}