package dto

import "time"

type CreateUserRequestDTO struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type CreateUserResponseDTO struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Role      Role      `json:"role"`
	APIKey    string    `json:"api_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
