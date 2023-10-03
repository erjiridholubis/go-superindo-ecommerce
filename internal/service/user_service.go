package service

import (
	"context"
	"errors"
	
	"github.com/erjiridholubis/go-superindo-product/common"
	"github.com/erjiridholubis/go-superindo-product/internal/model"
	"github.com/erjiridholubis/go-superindo-product/internal/repository"
)

type userService struct {
	postgreRepo repository.PostgreRepository
}

func NewUserService(postgreRepo repository.PostgreRepository) UserService {
	return &userService{postgreRepo}
}

type UserService interface {
	GetUserByID(ctx context.Context, id string) (resp *model.UserResponse, err error)
}

func (us *userService) GetUserByID(ctx context.Context, id string) (resp *model.UserResponse, err error) {
	resp, err = us.postgreRepo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, errors.New(common.ErrNotFound)
	}

	return resp, nil
}