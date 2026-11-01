package dto

type Currency string

const (
	KZTCurrency Currency = "KZT"
)

type CreateProductRequestDTO struct {
	Sku      string   `json:"sku" validate:"required"`
	Name     string   `json:"name" validate:"required"`
	Price    int      `json:"price" validate:"required"`
	Currency Currency `json:"currency" validate:"required"`
}
