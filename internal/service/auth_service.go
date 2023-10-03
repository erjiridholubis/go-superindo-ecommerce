package service

import (
	"context"
	"errors"

	"github.com/erjiridholubis/go-superindo-product/common"
	"github.com/erjiridholubis/go-superindo-product/internal/middleware"
	"github.com/erjiridholubis/go-superindo-product/internal/model"
	"github.com/erjiridholubis/go-superindo-product/internal/repository"
)

type authService struct {
	postgreRepo repository.PostgreRepository
}

func NewAuthService(postgreRepo repository.PostgreRepository) AuthService {
	return &authService{postgreRepo}
}

type AuthService interface {
	Login(ctx context.Context, payload model.AuthRequest) (resp *model.AuthResponse, err error)
	Register(ctx context.Context, payload model.UserRequest) (resp *model.UserResponse, err error)
}

func (as *authService) Login(ctx context.Context, payload model.AuthRequest) (resp *model.AuthResponse, err error) {
	user, err := as.postgreRepo.GetUserByUsername(ctx, payload.Username)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New(common.ErrNotFound)
	}

	if !common.ComparePasswords(user.Password, []byte(payload.Password)) {
		return nil, errors.New(common.ErrInvalidPassword)
	}

	token, err := middleware.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	resp = &model.AuthResponse{
		Kind: common.KindAuthorization,
		Token: token,
	}

	return resp, nil
}

func (as *authService) Register(ctx context.Context, payload model.UserRequest) (resp *model.UserResponse, err error) {
	user, err := as.postgreRepo.GetUserByUsername(ctx, payload.Username)
	if err != nil {
		return nil, err
	}

	if user != nil {
		return nil, errors.New(common.ErrUsernameAlreadyExist)
	}

	hashedPassword, err := common.HashAndSalt([]byte(payload.Password))
	if err != nil {
		return nil, err
	}

	user = &model.User{
		ID: common.GenerateUUID(),
		Name: payload.Name,
		Username: payload.Username,
		Password: hashedPassword,
	}

	id, err := as.postgreRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	resp = &model.UserResponse{
		Kind: common.KindUser,
		ID: id,
		Name: user.Name,
		Username: user.Username,
	}

	return resp, nil
}