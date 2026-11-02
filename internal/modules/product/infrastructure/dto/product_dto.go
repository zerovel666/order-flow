package dto

import "time"

type Product struct {
	ID        int
	Sku       string
	Name      string
	Price     int
	Currency  string
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
