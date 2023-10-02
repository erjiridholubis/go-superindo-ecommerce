package service

import (
	"github.com/erjiridholubis/go-superindo-product/internal/repository"
)

type productService struct {
	productRepo repository.ProductRepository
}

func NewProductService(productRepo repository.ProductRepository) ProductService {
	return &productService{productRepo}
}

type ProductService interface {

}