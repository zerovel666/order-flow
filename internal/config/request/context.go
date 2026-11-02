package request

import (
	"errors"
	"net/http"
	"order-flow/internal/modules/user/infrastructure/dto"
)

type ctxKey string

const (
	UserKey ctxKey = "user"
)

func User(r *http.Request) (*dto.User, error) {
	user, ok := r.Context().Value(UserKey).(*dto.User)
	if !ok || user == nil {
		return nil, errors.New("Пользователь в контексте не найден")
	}

	return user, nil
}
