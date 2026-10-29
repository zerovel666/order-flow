package dto

type LoginRequestDTO struct {
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponseDTO struct {
	APIKey string `json:"api_key"`
}
