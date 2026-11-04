package dto

import "time"

type Product struct {
	ID        int       `json:"id"`
	Sku       string    `json:"sku"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	Currency  Currency  `json:"currency"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Inventory struct {
	ID                int       `json:"id"`
	AvailableQuantity int       `json:"available_quantity"`
	UpdatedAt         time.Time `json:"updated_at"`
}
