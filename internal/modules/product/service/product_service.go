package service

import (
	"context"
	"order-flow/internal/config/response"
	"order-flow/internal/modules/product/infrastructure/dto"
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

func (s *ProductService) Create(ctx context.Context, dto dto.CreateProductRequestDTO) response.Response {
	return response.SuccessAPIResponse(
		"test",
		[]response.Message{},
		nil,
	)
}

func (s *ProductService) Update(ctx context.Context, dto dto.UpdateProductRequestDTO) response.Response {
	return response.SuccessAPIResponse(
		"test",
		[]response.Message{},
		nil,
	)
}

func (s *ProductService) UnActive(ctx context.Context) response.Response {
	return response.SuccessAPIResponse(
		"test",
		[]response.Message{},
		nil,
	)
}

func (s *ProductService) GetAll(ctx context.Context) response.Response {
	return response.SuccessAPIResponse(
		"test",
		[]response.Message{},
		nil,
	)
}

func (s *ProductService) GetMyShop(ctx context.Context) response.Response {
	return response.SuccessAPIResponse(
		"test",
		[]response.Message{},
		nil,
	)
}
