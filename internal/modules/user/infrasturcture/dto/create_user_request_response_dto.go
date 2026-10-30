package dto

import "time"

type CreateUserRequestDTO struct {
	Name     string `json:"name" validate:"required"`
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required"`
	Role     Role   `json:"role" validate: "required"`
}

type CreateUserResponseDTO struct {
	ID        uint64    `json:"id"`
	Login     string    `json:"login"`
	Name      string    `json:"name"`
	Role      Role      `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
