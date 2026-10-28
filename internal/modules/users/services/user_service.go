package services

import (
	"context"
	"order-flow/internal/config/responses"
	"order-flow/internal/modules/users/infrasturcture/structs"
)

type UserRepository interface {
	Create(ctx context.Context, u structs.User) (structs.User, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(ctx context.Context, dto structs.CreateUserRequestDTO) responses.Response {
	u := structs.User{
		Name: dto.Name,
		Role: structs.ClientRole,
	}

	api_key, err := u.GenApiKey()
	if err != nil {
		return responses.ErrorResponse(
			"ERR_GEN_API_KEY",
			[]responses.Message{
				{Lang: "RU", Message: "Ошибка при генерации API-ключа"},
				{Lang: "EN", Message: "Error generating the API-key"},
			},
			nil,
			err.Error(),
		)
	}

	u.APIKeyHash = u.Hash(api_key)
	u.Password = u.Hash(dto.Password)

	u, err = s.repo.Create(ctx, u)
	if err != nil {
		return responses.ErrorResponse(
			"ERR_REGISTRATION_USER",
			[]responses.Message{
				{Lang: "RU", Message: "Ошибка при регистрации пользователя"},
				{Lang: "EN", Message: "Error registration the user"},
			},
			nil,
			err.Error(),
		)
	}
	return responses.SuccessAPIResponse(
		"USER_REGISTRATION",
		[]responses.Message{
			{Lang: "RU", Message: "Пользователь успешно зарегестрирован"},
			{Lang: "EN", Message: "The user has been registrated"},
		},
		structs.CreateUserResponseDTO{
			ID:     u.ID,
			Name:   u.Name,
			APIKey: api_key,
			Role: u.Role,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		},
	)
}
