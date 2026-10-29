package service

import (
	"context"
	"order-flow/internal/config/response"
	"order-flow/internal/config/type_error"
	"order-flow/internal/modules/user/infrasturcture/dto"
)

type UserRepository interface {
	Create(ctx context.Context, u dto.User) (dto.User, error)
	Login(ctx context.Context, u dto.User) (bool, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(ctx context.Context, data dto.CreateUserRequestDTO) response.Response {
	u := dto.User{
		Name:  data.Name,
		Login: data.Login,
		Role:  dto.ClientRole,
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
		typeErr := type_error.PgErr(err)
		return response.PgErrResponse(typeErr, err)
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
			Login:     u.Login,
			Role:      u.Role,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		},
	)
}

func (s *UserService) Login(ctx context.Context, data dto.LoginRequestDTO) response.Response {
	u := dto.User{
		Login: data.Login,
	}

	u.Password = u.Hash(data.Password)
	api_key, err := u.GenApiKey()
	u.APIKeyHash = u.Hash(api_key)
	login, err := s.repo.Login(ctx, u)

	if err != nil || !login {
		return response.ErrorResponse(
			"UNAUTHORIZATION",
			[]response.Message{
				{Lang: "RU", Message: "Не удалось авторизировать пользователя"},
				{Lang: "EN", Message: "Failed user authorization"},
			},
			nil,
			nil,
		)
	}

	return response.SuccessAPIResponse(
		"AUTHORIZED",
		[]response.Message{
			{Lang: "RU", Message: "Пользователь успешно авторизован"},
			{Lang: "EN", Message: "The user has been successfully sign up"},
		},
		dto.LoginResponseDTO{
			APIKey: api_key,
		},
	)
}
