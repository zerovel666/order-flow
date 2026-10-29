package service

import (
	"context"
	"order-flow/internal/config/response"
	"order-flow/internal/modules/user/infrasturcture/dto"
)

type UserRepository interface {
	Create(ctx context.Context, u dto.User) (dto.User, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(ctx context.Context, data dto.CreateUserRequestDTO) response.Response {
	u := dto.User{
		Name: data.Name,
		Role: dto.ClientRole,
	}

	api_key, err := u.GenApiKey()
	if err != nil {
		return response.ErrorResponse(
			"ERR_GEN_API_KEY",
			[]response.Message{
				{Lang: "RU", Message: "Ошибка при генерации API-ключа"},
				{Lang: "EN", Message: "Error generating the API-key"},
			},
			nil,
			err.Error(),
		)
	}

	u.APIKeyHash = u.Hash(api_key)
	u.Password = u.Hash(data.Password)

	u, err = s.repo.Create(ctx, u)
	if err != nil {
		return response.ErrorResponse(
			"ERR_REGISTRATION_USER",
			[]response.Message{
				{Lang: "RU", Message: "Ошибка при регистрации пользователя"},
				{Lang: "EN", Message: "Error registration the user"},
			},
			nil,
			err.Error(),
		)
	}
	return response.SuccessAPIResponse(
		"USER_REGISTRATION",
		[]response.Message{
			{Lang: "RU", Message: "Пользователь успешно зарегестрирован"},
			{Lang: "EN", Message: "The user has been registrated"},
		},
		dto.CreateUserResponseDTO{
			ID:        u.ID,
			Name:      u.Name,
			APIKey:    api_key,
			Role:      u.Role,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		},
	)
}
