package middleware

import (
	"context"
	"net/http"
	"order-flow/internal/config/request"
	"order-flow/internal/config/response"
	"order-flow/internal/modules/user/infrastructure/dto"
	"order-flow/internal/modules/user/repository"
	"strings"
)

func AuthMiddleware(repo *repository.UserRepository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")
			token, _ := strings.CutPrefix(header, "Bearer ")
			if token == "" {
				writeErr(w)
				return
			}

			u := dto.User{}
			u.APIKeyHash = u.Hash(token)
			exists, err := repo.CheckToken(r.Context(), u.APIKeyHash)
			if err != nil || !exists {
				writeErr(w)
				return
			}

			u, err = repo.GetInfoByToken(r.Context(), u)
			if err != nil {
				writeErr(w)
				return
			}

			ctx := context.WithValue(r.Context(), request.UserKey, &u)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func writeErr(w http.ResponseWriter) {
	response.WriteJSON(
		w,
		401,
		response.ErrorResponse(
			"UNAUTHORIZATION",
			[]response.Message{
				{Lang: "RU", Message: "Неавторизованный пользователь"},
				{Lang: "EN", Message: "Unauthorization user"},
			},
			nil,
			nil,
		),
	)
}
