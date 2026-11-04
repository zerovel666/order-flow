package dto

type Currency string

const (
	KZTCurrency Currency = "KZT"
)

type CreateProductRequestDTO struct {
	Sku               string   `json:"sku" validate:"required"`
	Name              string   `json:"name" validate:"required"`
	Price             int      `json:"price"`
	Currency          Currency `json:"currency" validate:"required"`
	AvailableQuantity int      `json:"available_quantity"`
}

type CreateProductResponseDTO struct {
	Product
	Inventory Inventory `json:"inventory"`
}
