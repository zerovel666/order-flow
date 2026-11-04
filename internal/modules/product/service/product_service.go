package service

import (
	"context"
	"order-flow/internal/config/response"
	"order-flow/internal/config/type_error"
	"order-flow/internal/modules/product/infrastructure/dto"
)

type ProductRepository interface {
	Create(ctx context.Context, data dto.Product) (dto.Product, error)
	Update(ctx context.Context, data dto.Product) error
	GetAll(ctx context.Context) ([]dto.Product, error)
	GetById(ctx context.Context, id int) (dto.Product, error)
	UnActive(ctx context.Context, id int) error
	CreateInventory(ctx context.Context, data dto.Inventory) (dto.Inventory, error)
}

type ProductService struct {
	repo ProductRepository
}

func NewProductService(repo ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) Create(ctx context.Context, data dto.CreateProductRequestDTO) response.Response {
	product := dto.Product{
		Sku:      data.Sku,
		Name:     data.Name,
		Price:    data.Price,
		Currency: data.Currency,
		IsActive: true,
	}

	product, err := s.repo.Create(ctx, product)
	if err != nil {
		typeErr := type_error.PgErr(err)
		return response.PgErrResponse(typeErr, err)
	}

	inventory := dto.Inventory{
		ID:                product.ID,
		AvailableQuantity: data.AvailableQuantity,
	}
	inventory, err = s.repo.CreateInventory(ctx, inventory)
	if err != nil {
		typeErr := type_error.PgErr(err)
		return response.PgErrResponse(typeErr, err)
	}

	return response.SuccessAPIResponse(
		"PRODUCT_CREATED",
		[]response.Message{
			{Lang: "RU", Message: "Товар успешно создан"},
			{Lang: "EN", Message: "Product successfully created"},
		},
		dto.CreateProductResponseDTO{
			Product:   product,
			Inventory: inventory,
		},
	)
}

func (s *ProductService) UnActive(ctx context.Context, product_id int) response.Response {
	if err := s.repo.UnActive(ctx, product_id); err != nil {
		return response.PgErrResponse(type_error.PgErr(err), err)
	}

	return response.SuccessAPIResponse("PRODUCT_DEACTIVATED", []response.Message{
		{Lang: "RU", Message: "Товар деактивирован"},
		{Lang: "EN", Message: "Product deactivated"},
	}, nil)
}

func (s *ProductService) GetAll(ctx context.Context) response.Response {
	products, err := s.repo.GetAll(ctx)
	if err != nil {
		return response.PgErrResponse(type_error.PgErr(err), err)
	}
	return response.SuccessAPIResponse(
		"PRODUCTS_LIST",
		[]response.Message{},
		products,
	)
}
