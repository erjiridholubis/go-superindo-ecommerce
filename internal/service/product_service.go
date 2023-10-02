package service

import (
	"github.com/erjiridholubis/go-superindo-product/internal/repository"
)

type productService struct {
	productRepo repository.PostgreRepository
}

func NewProductService(productRepo repository.PostgreRepository) ProductService {
	return &productService{productRepo}
}

type ProductService interface {

}