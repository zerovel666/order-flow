package dto

type UpdateProductRequestDTO struct {
	Sku      string   `json:"sku" validate:"nullable"`
	Name     string   `json:"name" validate:"nullable"`
	Price    int      `json:"price" validate:"nullable"`
	Currency Currency `json:"currency" validate:"nullable"`
}
