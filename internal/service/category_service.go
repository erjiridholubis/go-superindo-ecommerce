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