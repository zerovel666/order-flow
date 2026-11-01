package service

import (
	"context"
	"order-flow/internal/modules/products/infrastructure/dto"
)

type ProductRepository interface {
	Create(ctx context.Context, data dto.Product)
	Update(ctx context.Context, data dto.Product)
	GetAll(ctx context.Context)
	GetMyShop(ctx context.Context, userSellerId int)
}

type ProductService struct {
	repo ProductRepository
}

func NewProductService(repo ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) Create(ctx context.Context, dto dto.CreateProductRequestDTO) {

}

func (s *ProductService) Update(ctx context.Context, dto dto.UpdateProductRequestDTO) {

}

func (s *ProductService) UnActive(ctx context.Context) {

}

func (s *ProductService) GetAll(ctx context.Context) {

}

func (s *ProductService) GetMyShop(ctx context.Context) {

}
